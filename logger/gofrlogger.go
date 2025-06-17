package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"time"

	"golang.org/x/term"
	"moneyx.golang.framework/logger/logrepo"

	"gofr.dev/pkg/gofr/logging"
	"gofr.dev/pkg/gofr/version"
)

const fileMode = 0644

type GoFrLogger struct {
	logRepo    *logrepo.MoneyxLogRepo
	level      logging.Level
	normalOut  io.Writer
	errorOut   io.Writer
	isTerminal bool
	lock       chan struct{}
}

type GoFrLogEntry struct {
	Level       logging.Level `json:"level"`
	Time        time.Time     `json:"time"`
	Message     any           `json:"message"`
	TraceID     string        `json:"trace_id,omitempty"`
	GofrVersion string        `json:"gofrVersion"`
}

func frLogLevelToMoneyxLogLevel(level logging.Level) logrepo.MoneyxLogLevel {
	switch level {
	case logging.DEBUG:
		return logrepo.DEBUG
	case logging.INFO:
		return logrepo.INFO
	case logging.NOTICE:
		return logrepo.INFO
	case logging.WARN:
		return logrepo.WARN
	case logging.ERROR:
		return logrepo.ERROR
	case logging.FATAL:
		return logrepo.FATAL
	default:
		return logrepo.INFO
	}
}

func (l *GoFrLogger) logf(level logging.Level, format string, args ...any) {
	if level < l.level {
		return
	}

	out := l.normalOut
	stacktrace := ""
	if level >= logging.ERROR {
		out = l.errorOut
		stacktrace = string(debug.Stack())
	}

	entry := GoFrLogEntry{
		Level:       level,
		Time:        time.Now(),
		GofrVersion: version.Framework,
	}

	traceID, filteredArgs := extractTraceIDAndFilterArgs(args)
	entry.TraceID = traceID

	switch {
	case len(filteredArgs) == 1 && format == "":
		entry.Message = filteredArgs[0]
	case len(filteredArgs) != 1 && format == "":
		entry.Message = filteredArgs
	case format != "":
		entry.Message = fmt.Sprintf(format, filteredArgs...)
	}

	if l.isTerminal {
		l.prettyPrint(&entry, out)
	} else {
		_ = json.NewEncoder(out).Encode(entry)
	}
	l.logRepo.InsertMessage(frLogLevelToMoneyxLogLevel(level), fmt.Sprintf("%v", entry.Message), entry.Time.Format("2006-01-02 15:04:05"), stacktrace)
}

func (l *GoFrLogger) Debug(args ...any) {
	l.logf(logging.DEBUG, "", args...)
}

func (l *GoFrLogger) Debugf(format string, args ...any) {
	l.logf(logging.DEBUG, format, args...)
}

func (l *GoFrLogger) Info(args ...any) {
	l.logf(logging.INFO, "", args...)
}

func (l *GoFrLogger) Infof(format string, args ...any) {
	l.logf(logging.INFO, format, args...)
}

func (l *GoFrLogger) Notice(args ...any) {
	l.logf(logging.NOTICE, "", args...)
}

func (l *GoFrLogger) Noticef(format string, args ...any) {
	l.logf(logging.NOTICE, format, args...)
}

func (l *GoFrLogger) Warn(args ...any) {
	l.logf(logging.WARN, "", args...)
}

func (l *GoFrLogger) Warnf(format string, args ...any) {
	l.logf(logging.WARN, format, args...)
}

func (l *GoFrLogger) Log(args ...any) {
	l.logf(logging.INFO, "", args...)
}

func (l *GoFrLogger) Logf(format string, args ...any) {
	l.logf(logging.INFO, format, args...)
}

func (l *GoFrLogger) Error(args ...any) {
	l.logf(logging.ERROR, "", args...)
}

func (l *GoFrLogger) Errorf(format string, args ...any) {
	l.logf(logging.ERROR, format, args...)
}

func (l *GoFrLogger) Fatal(args ...any) {
	l.logf(logging.FATAL, "", args...)

	//nolint:revive // exit status is 1 as it denotes failure as signified by Fatal log
	os.Exit(1)
}

func (l *GoFrLogger) Fatalf(format string, args ...any) {
	l.logf(logging.FATAL, format, args...)

	//nolint:revive // exit status is 1 as it denotes failure as signified by Fatal log
	os.Exit(1)
}

func (l *GoFrLogger) prettyPrint(e *GoFrLogEntry, out io.Writer) {
	// Note: we need to lock the pretty print as printing to standard output not concurrency safe
	// the logs when printed in go routines were getting misaligned since we are achieving
	// a single line of log, in 2 separate statements which caused the misalignment.
	l.lock <- struct{}{} // Acquire the channel's lock
	defer func() {
		<-l.lock // Release the channel's token
	}()

	// Pretty printing if the message interface defines a method PrettyPrint else print the log message
	// This decouples the logger implementation from its usage
	fmt.Fprintf(out, "\u001B[38;5;%dm%s\u001B[0m [%s]", e.Level.Color(), e.Level.String()[0:4], e.Time.Format(time.TimeOnly))

	if e.TraceID != "" {
		fmt.Fprintf(out, " \u001B[38;5;8m%s\u001B[0m", e.TraceID)
	}

	fmt.Fprint(out, " ")

	// Print the message
	if fn, ok := e.Message.(logging.PrettyPrint); ok {
		fn.PrettyPrint(out)
	} else {
		fmt.Fprintf(out, "%v\n", e.Message)
	}
}

// NewLogger creates a new logger instance with the specified logging level.
func NewGoFrLogger(level logging.Level, logRepo *logrepo.MoneyxLogRepo) logging.Logger {
	l := &GoFrLogger{
		normalOut: os.Stdout,
		errorOut:  os.Stderr,
		lock:      make(chan struct{}, 1),
		logRepo:   logRepo,
	}

	l.level = level

	l.isTerminal = checkIfTerminal(l.normalOut)

	return l
}

// NewFileLogger creates a new logger instance with logging to a file.
func NewFileLogger(path string) logging.Logger {
	l := &GoFrLogger{
		normalOut: io.Discard,
		errorOut:  io.Discard,
	}

	if path == "" {
		return l
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, fileMode)
	if err != nil {
		return l
	}

	l.normalOut = f
	l.errorOut = f

	return l
}

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return term.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}

func (l *GoFrLogger) ChangeLevel(level logging.Level) {
	l.level = level
}

// GetLogLevelForError returns the log level for the given error.
// If the error implements [logLevelResponder], its log level is returned.
// Otherwise, the default log level "error" is returned.
func GetLogLevelForError(err error) logging.Level {
	level := logging.ERROR

	if e, ok := err.(logging.LogLevelResponder); ok {
		level = e.LogLevel()
	}

	return level
}

// extractTraceIDAndFilterArgs scans log arguments for a trace ID map and
// returns the extracted trace ID (if found) and a filtered list of log arguments
// excluding the trace metadata.
func extractTraceIDAndFilterArgs(args []any) (traceID string, filtered []any) {
	filtered = make([]any, 0, len(args))

	for _, arg := range args {
		if m, ok := arg.(map[string]any); ok {
			if tid, exists := m["__trace_id__"].(string); exists && traceID == "" {
				traceID = tid

				continue
			}
		}

		filtered = append(filtered, arg)
	}

	return traceID, filtered
}

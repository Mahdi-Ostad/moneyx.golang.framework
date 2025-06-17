package logger

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"gofr.dev/pkg/gofr/datasource/sql"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// Config logger config
type GormConfig struct {
	SlowThreshold             time.Duration
	Colorful                  bool
	IgnoreRecordNotFoundError bool
	ParameterizedQueries      bool
	LogLevel                  logger.LogLevel
}

type DefaultMetrics struct {
}

func (d *DefaultMetrics) RecordHistogram(ctx context.Context, name string, value float64, labels ...string) {

}
func (d *DefaultMetrics) SetGauge(name string, value float64, labels ...string) {

}

var (
	// Discard logger will print any log to io.Discard
	Discard = NewGormLogger(NewMoneyxLog(), GormConfig{}, &DefaultMetrics{})
	// Default Default logger
	Default = NewGormLogger(NewMoneyxLog(), GormConfig{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	}, &DefaultMetrics{})
	// Recorder logger records running SQL into a recorder instance
	Recorder = GormTraceRecorder{Interface: Default, BeginAt: time.Now()}

	// RecorderParamsFilter defaults to no-op, allows to be run-over by a different implementation
	RecorderParamsFilter = func(ctx context.Context, sql string, params ...interface{}) (string, []interface{}) {
		return sql, params
	}
)

// New initialize logger
func NewGormLogger(writer *MoneyxLog, config GormConfig, metrics sql.Metrics) logger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		infoStr = logger.Green + "%s\n" + logger.Reset + logger.Green + "[info] " + logger.Reset
		warnStr = logger.BlueBold + "%s\n" + logger.Reset + logger.Magenta + "[warn] " + logger.Reset
		errStr = logger.Magenta + "%s\n" + logger.Reset + logger.Red + "[error] " + logger.Reset
		traceStr = logger.Green + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
		traceWarnStr = logger.Green + "%s " + logger.Yellow + "%s\n" + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s" + logger.Reset
		traceErrStr = logger.RedBold + "%s " + logger.MagentaBold + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
	}

	return &GormLogger{
		MoneyxLog:    writer,
		GormConfig:   config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
		metrics:      metrics,
	}
}

type GormLogger struct {
	*MoneyxLog
	GormConfig
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
	metrics                             sql.Metrics
}

// LogMode log mode
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.Infof(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.Warnf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.Errorf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
//
//nolint:cyclop
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql := ""
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Errorf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Errorf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Warnf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Warnf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Infof(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
	if sql != "" {
		l.metrics.RecordHistogram(context.Background(), "app_sql_stats", float64(elapsed), "type", getOperationType(sql))
	}
}

func getOperationType(query string) string {
	query = strings.TrimSpace(query)
	words := strings.Split(query, " ")

	return words[0]
}

// ParamsFilter filter params
func (l *GormLogger) ParamsFilter(ctx context.Context, sql string, params ...interface{}) (string, []interface{}) {
	if l.GormConfig.ParameterizedQueries {
		return sql, nil
	}
	return sql, params
}

type GormTraceRecorder struct {
	logger.Interface
	BeginAt      time.Time
	SQL          string
	RowsAffected int64
	Err          error
}

// New trace recorder
func (l *GormTraceRecorder) New() *GormTraceRecorder {
	return &GormTraceRecorder{Interface: l.Interface, BeginAt: time.Now()}
}

// Trace implement logger interface
func (l *GormTraceRecorder) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	l.BeginAt = begin
	l.SQL, l.RowsAffected = fc()
	l.Err = err
}

func (l *GormTraceRecorder) ParamsFilter(ctx context.Context, sql string, params ...interface{}) (string, []interface{}) {
	if RecorderParamsFilter == nil {
		return sql, params
	}
	return RecorderParamsFilter(ctx, sql, params...)
}

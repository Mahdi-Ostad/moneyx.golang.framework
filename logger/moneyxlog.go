package logger

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"moneyx.golang.framework/config"
	"moneyx.golang.framework/logger/logrepo"
)

type MoneyxLog struct {
	LogRepo *logrepo.MoneyxLogRepo
	mod     string
	color   bool
	min     logrepo.MoneyxLogLevel
}

var colors = map[logrepo.MoneyxLogLevel]string{
	logrepo.INFO:  "\033[36m",
	logrepo.WARN:  "\033[33m",
	logrepo.ERROR: "\033[31m",
	logrepo.FATAL: "\033[31m",
}

func (s *MoneyxLog) outputf(level logrepo.MoneyxLogLevel, msg string, args ...interface{}) {
	var colorStart, colorReset string
	if s.color {
		colorStart = colors[level]
		colorReset = "\033[0m"
	}
	fmt.Printf("%s%s [%s %s] %s%s\n", time.Now().Format("15:04:05.000"), colorStart, s.mod, level, fmt.Sprintf(msg, args...), colorReset)
	dt := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(fmt.Sprintf("[%s] ", s.mod)+msg, args...)
	stacktrace := string(debug.Stack())
	if level < s.min {
		return
	}
	if level < 3 {
		stacktrace = ""
	}
	//message := fmt.Sprintf("%s%s [%s %s] %s%s\n", time.Now().Format("15:04:05.000"), colorStart, s.mod, level, fmt.Sprintf(msg, args...), colorReset)
	s.LogRepo.InsertMessage(level, message, dt, stacktrace)
}

func (s *MoneyxLog) Errorf(msg string, args ...interface{}) { s.outputf(logrepo.ERROR, msg, args...) }
func (s *MoneyxLog) Warnf(msg string, args ...interface{})  { s.outputf(logrepo.WARN, msg, args...) }
func (s *MoneyxLog) Infof(msg string, args ...interface{})  { s.outputf(logrepo.INFO, msg, args...) }
func (s *MoneyxLog) Debugf(msg string, args ...interface{}) { s.outputf(logrepo.DEBUG, msg, args...) }
func (s *MoneyxLog) Fatalf(msg string, args ...interface{}) { s.outputf(logrepo.FATAL, msg, args...) }

var NewMoneyxLog = sync.OnceValue(getNewMoneyxLog)

func getNewMoneyxLog() *MoneyxLog {
	logRepo := logrepo.NewMoneyxLogRepo()
	logLevel := config.GetConfig("LOG_MIN_LEVEL")
	return &MoneyxLog{LogRepo: logRepo, mod: "Client", color: true, min: logrepo.ParseMoneyxLogLevel(logLevel)}
}

package logger

import (
	_ "errors"
	"fmt"
	_ "log/syslog"
	"os"
	"strings"
	"time"
)

type (
	LogLevel  int
	LogRecord struct {
		Time     time.Time              `json:"time"`
		AppName  string                 `json:"app_name"`
		Level    string                 `json:"level"`
		Caller   string                 `json:"caller"`
		Message  string                 `json:"message"`
		Error    string                 `json:"error,omitempty"`
		TracerID string                 `json:"tracer_id"`
		Fields   map[string]interface{} `json:"fields,omitempty"`
	}

	Logger struct {
		level    LogLevel
		handler  Handler
		tracerID string
		appName  string
	}
)

const (
	LevelNone LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	ScopeLevel = "SCOPE_LEVEL_LOGGER"
	AppName    = "APP_NAME_LOGGER"
)

var logLevels = [...]string{
	LevelNone:  "NONE",
	LevelDebug: "DEBUG",
	LevelInfo:  "INFO",
	LevelWarn:  "WARN",
	LevelError: "ERROR",
	LevelFatal: "FATAL",
}

func (l LogLevel) String() string {
	if l >= LevelNone && l < LevelFatal {
		return logLevels[l]
	}
	return "UNKNOWN"
}

func (l *Logger) log(msg string, level LogLevel) {
	if l.handler == nil {
		return
	}
	record := LogRecord{
		Time:     time.Now().Local(),
		Level:    level.String(),
		Caller:   getCallerInfo(),
		Message:  msg,
		TracerID: l.tracerID,
		AppName:  l.appName,
	}

	if l.level <= level {
		_ = l.handler.Handle(&record)
	}
}

func (l *Logger) logError(msg string, err error, level LogLevel) {
	if l.handler == nil {
		return
	}
	record := LogRecord{
		Time:     time.Now().Local(),
		Level:    level.String(),
		Caller:   getCallerInfo(),
		Message:  msg,
		TracerID: l.tracerID,
		AppName:  l.appName,
		Error:    err.Error(),
	}

	if l.level <= level {
		_ = l.handler.Handle(&record)
	}
}

func NewLogger(level LogLevel) *Logger {
	logg := &Logger{level: level}
	return logger(logg)
}

func NewLoggerWithTrace(tracerID string) *Logger {
	level := getLevel()
	logg := &Logger{level: level, tracerID: tracerID, appName: os.Getenv(AppName)}
	return logger(logg)
}

func getLevel() LogLevel {
	scopeLevelLogger := os.Getenv(ScopeLevel)
	var level = LevelInfo
	if strings.EqualFold(scopeLevelLogger, "debug") {
		level = LevelDebug
	}
	return level
}

func logger(logg *Logger) *Logger {
	handler := newJSONHandler(os.Stdout)
	logg.SetHandler(handler)
	return logg
}
func (l *Logger) Debug(msg string) {
	l.log(msg, LevelDebug)
}

func (l *Logger) Info(msg string) {
	l.log(msg, LevelInfo)
}

func (l *Logger) Warn(msg string) {
	l.log(msg, LevelWarn)
}

func (l *Logger) Error(msg string, err error) {
	l.logError(msg, err, LevelError)
}

func (l *Logger) Fatal(msg string) {
	l.log(msg, LevelFatal)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.log(fmt.Sprintf(format, args...), LevelInfo)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.log(fmt.Sprintf(format, args...), LevelWarn)
}

func (l *Logger) Errorf(format string, err error, args ...interface{}) {
	l.logError(fmt.Sprintf(format, args...), err, LevelError)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.log(fmt.Sprintf(format, args...), LevelFatal)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log(fmt.Sprintf(format, args...), LevelDebug)
}

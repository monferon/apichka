package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

type Logger struct {
	logger *slog.Logger
}

// var _ Interface = (*Logger)(nil)

func NewLogger(level string) *Logger {
	lvl := new(slog.LevelVar)

	var log slog.Level

	switch strings.ToLower(level) {
	case "debug":
		log = slog.LevelDebug
	case "info":
		log = slog.LevelInfo
	case "warn":
		log = slog.LevelWarn
	case "error":
		log = slog.LevelError
	default:
		log = slog.LevelInfo
	}

	lvl.Set(log)

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: lvl,
	}))

	return &Logger{logger: logger}
}

func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.sendMessage("debug", message, args...)
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.log(message, args...)
}

func (l *Logger) Warn(message string, args ...interface{}) {
	l.log(message, args...)
}

func (l *Logger) Error(message interface{}, args ...interface{}) {
	l.sendMessage("error", message, args...)
}

func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.sendMessage("fatal", message, args...)

	os.Exit(1)
}

func (l *Logger) log(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info(message)
	} else {
		l.logger.Info(message, args...)
	}
}

func (l *Logger) sendMessage(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(msg.Error(), args...)
	case string:
		l.log(msg, args...)
	default:
		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}

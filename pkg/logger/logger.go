package logger

import "log/slog"

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

func (l *Logger) Debug(message interface{}, args ...interface{}) {
	//TODO implement me
	panic("implement m	e")
}

func (l *Logger) Info(message string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *Logger) Warn(message string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *Logger) Error(message interface{}, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

package logger

import (
	"fmt"
	"log"
	"os"
)

type SimpleLogger struct {
	newLogger *log.Logger
}

type Level int8

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func NewSimpleLogger() *SimpleLogger  {
	logger := log.New(os.Stdout, "", 0)

	return &SimpleLogger{newLogger: logger}
}

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	case LevelPanic:
		return "PANIC"
	}
	return ""
}

func (l SimpleLogger) output(level Level, message string) {
	content := fmt.Sprintf("%-5s : %s\n", level.String(), message)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

func (s SimpleLogger) Info(v ...interface{}) {
	s.output(LevelInfo, fmt.Sprint(v...))
}

func (s SimpleLogger) Infof(format string, v ...interface{}) {
	s.output(LevelInfo, fmt.Sprintf(format, v...))
}

func (s SimpleLogger) Error(v ...interface{}) {
	s.output(LevelError, fmt.Sprint(v...))
}

func (s SimpleLogger) Errorf(format string, v ...interface{}) {
	s.output(LevelError, fmt.Sprintf(format, v...))
}

func (s SimpleLogger) Debug(v ...interface{}) {
	s.output(LevelDebug, fmt.Sprint(v...))
}

func (s SimpleLogger) Debugf(format string, v ...interface{}) {
	s.output(LevelDebug, fmt.Sprintf(format, v...))
}

func (s SimpleLogger) Panic(v ...interface{}) {
	s.output(LevelPanic, fmt.Sprint(v...))
}

func (s SimpleLogger) Panicf(format string, v ...interface{}) {
	s.output(LevelPanic, fmt.Sprintf(format, v...))
}

func (s SimpleLogger) Warn(v ...interface{}) {
	s.output(LevelWarn, fmt.Sprint(v...))
}

func (s SimpleLogger) Warnf(format string, v ...interface{}) {
	s.output(LevelWarn, fmt.Sprintf(format, v...))
}

func (s SimpleLogger) Fatal(v ...interface{}) {
	s.output(LevelFatal, fmt.Sprint(v...))
}

func (s SimpleLogger) Fatalf(format string, v ...interface{}) {
	s.output(LevelFatal, fmt.Sprintf(format, v...))
}

func (s SimpleLogger) Println(message string) {
	s.newLogger.Println(message)
}

func (s SimpleLogger) Print(message string) {
	s.newLogger.Print(message)
}

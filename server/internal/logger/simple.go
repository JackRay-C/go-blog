package logger

import (
	"blog/internal/config"
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

type Fields map[string]interface{}

type SimpleLogger struct {
	//newLogger *log.Logger
	consoleLogger *log.Logger
	fileLogger *log.Logger
	caller    []string
	ctx       context.Context
	fields    Fields
	config    *config.App
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

func NewSimpleLogger(a *config.App) (*SimpleLogger, error) {
	s := a.Logs.Simple

	if s.Directory == "" {
		s.Directory = "logs"
	}
	if s.FileName == "" {
		s.FileName = "latest.log"
	}

	writer := &lumberjack.Logger{
		Filename:  path.Join(s.Directory, s.FileName),
		MaxSize:   s.LogMaxSize,
		MaxAge:    s.LogMaxAge,
		LocalTime: true,
		Compress:  true,
	}

	fileLogger := log.New(writer, "", log.Ldate | log.Ltime)
	consoleLogger := log.New(os.Stdout, "", log.Ldate | log.Ltime  )

	return &SimpleLogger{consoleLogger: consoleLogger ,fileLogger: fileLogger, config: a}, nil
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

func (l Level) Color() string  {
	switch l {
	case LevelDebug:
		return "\033[1;34mDEBUG\033[0m"
	case LevelInfo:
		return "\033[1;32mINFO\033[0m"
	case LevelWarn:
		return "\033[1;33mWARN\033[0m"
	case LevelError:
		return "\033[1;31mERROR\033[0m"
	case LevelFatal:
		return "\033[1;31mFATAL\033[0m"
	case LevelPanic:
		return "\033[1;31mPANIC\033[0m"
	}
	return ""
}


func (l *SimpleLogger) clone() *SimpleLogger {
	nl := *l
	return &nl
}

func (l *SimpleLogger) TextFormat(level Level, message string) string {
	pc, _, _, ok := runtime.Caller(3)
	if ok {
		f := runtime.FuncForPC(pc)
		if l.config.Logs.Simple.LogInConsole {
			l.caller = []string{fmt.Sprintf("\033[1;36m%-50s \033[0m", f.Name())}
		} else {
			l.caller = []string{fmt.Sprintf("%-50s", f.Name())}
		}
	}
	if l.config.Logs.Simple.LogInConsole {
		return fmt.Sprintf("%s %-16s : %s", l.caller, level.Color(), message)
	} else {
		return fmt.Sprintf("%s %-5s : %s", l.caller, level.String(), message)
	}
}


func (l *SimpleLogger) JsonFormat(level Level, message string) string {
	pc, _, _, ok := runtime.Caller(3)
	if ok {
		f := runtime.FuncForPC(pc)
		if l.config.Logs.Simple.LogInConsole {
			l.caller = []string{fmt.Sprintf("\033[1;36m%-50s \033[0m", f.Name())}
		} else {
			l.caller = []string{fmt.Sprintf("%-50s", f.Name())}
		}
	}
	data := make(Fields, 4)
	data["level"] = level.String()
	data["message"] = message
	data["caller"] = l.caller
	data["time"] = time.Now().Local().UnixNano()
	marshal, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (l *SimpleLogger) output(level Level, message string) {
	//content := fmt.Sprintf("%-5s : %s\n", level.String(), message)
	content := ""
	if l.config.Logs.Simple.Format == "console" {
		content = l.TextFormat(level, message)
	} else if l.config.Logs.Simple.Format == "json" {
		content = l.JsonFormat(level, message)
	}

	switch level {
	case LevelDebug:
		l.fileLogger.Println(content)
		if l.config.AppMode == config.DevelopmentMode.String() {
			l.consoleLogger.Println(content)
		}
	case LevelInfo:
		l.fileLogger.Println(content)
		l.consoleLogger.Println(content)
	case LevelWarn:
		l.fileLogger.Println(content)
		if l.config.AppMode == config.DevelopmentMode.String() {
			l.consoleLogger.Println(content)
		}
	case LevelError:
		l.fileLogger.Println(content)
		l.consoleLogger.Println(content)
	case LevelFatal:
		l.fileLogger.Println(content)
		l.consoleLogger.Println(content)
		os.Exit(1)
	case LevelPanic:
		l.fileLogger.Println(content)
		l.consoleLogger.Println(content)
		os.Exit(1)
	}
}

func (s *SimpleLogger) Info(v ...interface{}) {
	s.output(LevelInfo, fmt.Sprint(v...))
}

func (s *SimpleLogger) Infof(format string, v ...interface{}) {
	s.output(LevelInfo, fmt.Sprintf(format, v...))
}

func (s *SimpleLogger) Error(v ...interface{}) {
	s.output(LevelError, fmt.Sprint(v...))
}

func (s *SimpleLogger) Errorf(format string, v ...interface{}) {
	s.output(LevelError, fmt.Sprintf(format, v...))
}

func (s *SimpleLogger) Debug(v ...interface{}) {
	s.output(LevelDebug, fmt.Sprint(v...))
}

func (s *SimpleLogger) Debugf(format string, v ...interface{}) {
	s.output(LevelDebug, fmt.Sprintf(format, v...))
}

func (s *SimpleLogger) Panic(v ...interface{}) {
	s.output(LevelPanic, fmt.Sprint(v...))
}

func (s *SimpleLogger) Panicf(format string, v ...interface{}) {
	s.output(LevelPanic, fmt.Sprintf(format, v...))
}

func (s *SimpleLogger) Warn(v ...interface{}) {
	s.output(LevelWarn, fmt.Sprint(v...))
}

func (s *SimpleLogger) Warnf(format string, v ...interface{}) {
	s.output(LevelWarn, fmt.Sprintf(format, v...))
}

func (s *SimpleLogger) Fatal(v ...interface{}) {
	s.output(LevelFatal, fmt.Sprint(v...))
}

func (s *SimpleLogger) Fatalf(format string, v ...interface{}) {
	s.output(LevelFatal, fmt.Sprintf(format, v...))
}

func (s *SimpleLogger) Println(message string) {
	s.fileLogger.Println(message)
	s.consoleLogger.Println(message)
}

func (s *SimpleLogger) Print(message string) {
	s.fileLogger.Println(message)
	s.consoleLogger.Println(message)
}

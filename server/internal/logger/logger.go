package logger

import (
	"blog/internal/config"
	"errors"
	"strings"
)

type Logger interface {
	Info(v... interface{})
	Infof(format string, v... interface{})
	Error(v... interface{})
	Errorf(format string, v... interface{})
	Debug(v... interface{})
	Debugf(format string, v... interface{})
	Panic(v... interface{})
	Panicf(format string, v... interface{})
	Warn(v... interface{})
	Warnf(format string, v... interface{})
	Fatal(v... interface{})
	Fatalf(format string, v... interface{})
	Println(message string)
	Print(message string)
}

const (
	SimpleLog = "simple"
	ZapLog = "zap"

)

func New(setting *config.App) (Logger, error)  {
	if strings.ToLower(setting.AppLogType) == SimpleLog {
		return NewSimpleLogger(setting)
	} else if strings.ToLower(setting.AppLogType) == ZapLog {
		return NewZapLogger(setting)
	} else {
		// 默认日志
		return nil, errors.New("un support log type")
	}
}
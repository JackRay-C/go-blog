package logger

import (
	"blog/core/setting"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"time"
)

type ZapLogger struct {
	newLogger *zap.Logger
}

func NewZapLogger(zapSetting *setting.Zap) (logger *ZapLogger) {
	// 判断是否存在文件夹，不存在则创建
	if ok, _ := existPath(zapSetting.Director); !ok {
		log.Printf("创建日志目录%s\n", zapSetting.Director)
		_ = os.Mkdir(zapSetting.Director, os.ModePerm)
	}
	level := getZapLevel(zapSetting)

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = &ZapLogger{newLogger: zap.New(getZapEncoderCore(zapSetting), zap.AddStacktrace(level))}
	} else {
		logger = &ZapLogger{newLogger: zap.New(getZapEncoderCore(zapSetting))}
	}

	defer func(newLogger *zap.Logger) {
		err := newLogger.Sync()
		if err != nil {
			return
		}
	}(logger.newLogger)

	if zapSetting.ShowLine {
		logger.newLogger.WithOptions(zap.AddCaller())
	}
	return logger
}

func getZapLevel(zapSetting *setting.Zap) zapcore.Level {
	switch zapSetting.Level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func getZapEncoderConfig(zapSetting *setting.Zap) (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logs",
		CallerKey:      "caller",
		StacktraceKey:  zapSetting.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case zapSetting.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case zapSetting.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case zapSetting.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case zapSetting.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

func getZapEncoder(zapSetting *setting.Zap) zapcore.Encoder {
	if zapSetting.Format == "json" {
		return zapcore.NewJSONEncoder(getZapEncoderConfig(zapSetting))
	} else {
		return zapcore.NewConsoleEncoder(getZapEncoderConfig(zapSetting))
	}
}

func getZapSyncer(zapSetting *setting.Zap) (zapcore.WriteSyncer, error) {
	l := &lumberjack.Logger{
		Filename:  zapSetting.Director + "/" + zapSetting.LinkName,
		MaxSize:   zapSetting.LogMaxSize,
		MaxAge:    zapSetting.LogMaxAge,
		LocalTime: true,
		Compress:  true,
	}
	if zapSetting.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(l)), nil
	}

	return zapcore.AddSync(l), nil
}

func getZapEncoderCore(zapSetting *setting.Zap) (core zapcore.Core) {
	writer, err := getZapSyncer(zapSetting)
	if err != nil {
		return nil
	}
	return zapcore.NewCore(getZapEncoder(zapSetting), writer, getZapLevel(zapSetting))

}

func existPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func zapTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 15:04:05"))
}

func (z ZapLogger) Println(message string) {
	z.newLogger.Info(message)
}

func (z ZapLogger) Print(message string) {
	z.newLogger.Info(message)
}

func (z ZapLogger) Info(v ...interface{}) {
	z.newLogger.Info(fmt.Sprint(v...))
}

func (z ZapLogger) Infof(format string, v ...interface{}) {
	z.newLogger.Info(fmt.Sprintf(format, v...))
}

func (z ZapLogger) Error(v ...interface{}) {
	z.newLogger.Error(fmt.Sprint(v...))
}

func (z ZapLogger) Errorf(format string, v ...interface{}) {
	z.newLogger.Info(fmt.Sprintf(format, v...))
}

func (z ZapLogger) Debug(v ...interface{}) {
	z.newLogger.Debug(fmt.Sprint(v...))
}

func (z ZapLogger) Debugf(format string, v ...interface{}) {
	z.newLogger.Debug(fmt.Sprintf(format, v...))
}

func (z ZapLogger) Panic(v ...interface{}) {
	z.newLogger.Panic(fmt.Sprint(v...))
}

func (z ZapLogger) Panicf(format string, v ...interface{}) {
	z.newLogger.Panic(fmt.Sprintf(format, v...))
}

func (z ZapLogger) Warn(v ...interface{}) {
	z.newLogger.Warn(fmt.Sprint(v...))
}

func (z ZapLogger) Warnf(format string, v ...interface{}) {
	z.newLogger.Warn(fmt.Sprintf(format, v...))
}

func (z ZapLogger) Fatal(v ...interface{}) {
	z.newLogger.Fatal(fmt.Sprint(v...))
}

func (z ZapLogger) Fatalf(format string, v ...interface{}) {
	z.newLogger.Fatal(fmt.Sprintf(format, v...))
}

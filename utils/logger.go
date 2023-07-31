package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	conf "guying-go/configs"
	"os"
	"time"
)

/**
 * @program: guying-go
 * @description: 日志
 * @author: sickle
 * @create: 2023-07-27 17:29
 **/
var Logger *zap.Logger

func init() {

	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(level.CapitalString())
	}

	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}

	zapLoggerEncoderConfig := zapcore.EncoderConfig{
		TimeKey:          "time",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		MessageKey:       "message",
		StacktraceKey:    "stacktrace",
		EncodeCaller:     customCallerEncoder,
		EncodeTime:       customTimeEncoder,
		EncodeLevel:      customLevelEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		LineEnding:       "\n",
		ConsoleSeparator: " ",
	}

	var allCore []zapcore.Core

	level := zap.InfoLevel

	switch conf.Conf.Zap.Level {
	case "debug":
		level = zap.DebugLevel
		// 输出到控制台
		stdoutWriter := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
		consoleCore := zapcore.NewCore(zapcore.NewJSONEncoder(zapLoggerEncoderConfig), stdoutWriter, level)
		allCore = append(allCore, consoleCore)
	case "info":
	default:

	}
	// 输出到文件
	fileWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:  conf.Conf.Zap.File,    // ⽇志⽂件路径
			MaxSize:   conf.Conf.Zap.MaxSize, // 单位为MB,默认为512MB
			MaxAge:    conf.Conf.Zap.MaxAge,  // 文件最多保存多少天
			LocalTime: true,                  // 采用本地时间
			Compress:  false,                 // 是否压缩日志
		}),
		Size: 4096,
	}

	fileCore := zapcore.NewCore(zapcore.NewConsoleEncoder(zapLoggerEncoderConfig), fileWriter, zap.DebugLevel)
	allCore = append(allCore, fileCore)
	core := zapcore.NewTee(allCore...)

	// 所有的core
	Logger = zap.New(core, zap.AddCaller())
	return
}

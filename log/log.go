package log

import (
	"apiproject/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var Logger *zap.Logger

func Init() {
	Logger = GetFileRotatelogsLogger()
}

/**
获取基于lestrrat-go/file-rotatelogs归档的zap logger
*/
func GetFileRotatelogsLogger() *zap.Logger {
	//按实际需求灵活定义日志级别
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level <= zapcore.InfoLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.WarnLevel
	})

	infoWriter := getWriter(config.GlobalConfig.LogDir + "info.log")
	warnWriter := getWriter(config.GlobalConfig.LogDir + "warn.log")

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(
			// 编码器配置
			zapcore.NewJSONEncoder(NewEncoderConfig()),
			// 打印到控制台和文件
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
				zapcore.AddSync(infoWriter)),
			// 日志级别
			infoLevel,
		),
		zapcore.NewCore(
			// 编码器配置
			zapcore.NewJSONEncoder(NewEncoderConfig()),
			// 打印到控制台和文件
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
				zapcore.AddSync(warnWriter)),
			// 日志级别
			warnLevel,
		),
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", "miguo"))
	return zap.New(core, caller, development, filed)
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y-%m-%d_%H:%M:%S", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		//最多保留多久
		rotatelogs.WithMaxAge(time.Hour*time.Duration(config.GlobalConfig.LogMaxDayCount*24)),
		//多久做一次归档
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

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

/**
获取logger的全局字段配置
*/
func GetLoggerGlobalOption() zap.Option {
	return zap.Fields(zap.String("serviceName", config.GlobalConfig.ServiceName))
}

/**
获取zap logger: 用于写普通日志(除了kafka 自定义hook的日志)
*/
func GetCommonLogger() *zap.Logger {
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		getCoreList()...,
	)

	return zap.New(core, caller, development, GetLoggerGlobalOption())
}

/**
获取zap自定义hook的logger: 用于写自定义hook的日志
*/
func GetHookLogger() *zap.Logger {
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		getHookCoreList()...,
	)

	return zap.New(core, caller, development, GetLoggerGlobalOption())
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

/*
获取rotatelogs Hook
*/
func getRotatelogsHook(filename string) io.Writer {
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

/**
获取zap core列表
*/
func getCoreList() (coreList []zapcore.Core) {
	//按实际需求灵活定义日志级别
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level <= zapcore.InfoLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.WarnLevel
	})

	rotatelogsInfoHook := getRotatelogsHook(config.GlobalConfig.LogDir + "info.log")
	rotatelogsWarnHook := getRotatelogsHook(config.GlobalConfig.LogDir + "warn.log")
	kafkaLogHook := &KafkaLogHook{}
	mongoLogHook := &MongoLogHook{}

	//构建hook的WriteSyncer列表
	var infoWriteSyncerList, warnWriteSyncerList []zapcore.WriteSyncer
	infoWriteSyncerList = append(infoWriteSyncerList, zapcore.AddSync(os.Stdout), zapcore.AddSync(rotatelogsInfoHook))
	warnWriteSyncerList = append(warnWriteSyncerList, zapcore.AddSync(os.Stdout), zapcore.AddSync(rotatelogsWarnHook))
	if config.GlobalConfig.LogKafkaHookSwitch {
		infoWriteSyncerList = append(infoWriteSyncerList, zapcore.AddSync(kafkaLogHook))
		warnWriteSyncerList = append(warnWriteSyncerList, zapcore.AddSync(kafkaLogHook))
	}
	if config.GlobalConfig.LogMongoHookSwitch {
		infoWriteSyncerList = append(infoWriteSyncerList, zapcore.AddSync(mongoLogHook))
		warnWriteSyncerList = append(warnWriteSyncerList, zapcore.AddSync(mongoLogHook))
	}

	coreList = append(coreList,
		zapcore.NewCore(
			// 编码器配置
			zapcore.NewJSONEncoder(NewEncoderConfig()),
			// 打印到控制台和文件
			zapcore.NewMultiWriteSyncer(infoWriteSyncerList...),
			// 日志级别
			infoLevel,
		),
		zapcore.NewCore(
			// 编码器配置
			zapcore.NewJSONEncoder(NewEncoderConfig()),
			// 打印到控制台和文件
			zapcore.NewMultiWriteSyncer(warnWriteSyncerList...),
			// 日志级别
			warnLevel,
		))

	return
}

/**
获取为自定义hook提供的core列表: 用于记录自定义hook的日志
*/
func getHookCoreList() (coreList []zapcore.Core) {
	//按实际需求灵活定义日志级别
	debugLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})

	HookLogHook := getRotatelogsHook(config.GlobalConfig.LogDir + "hook.log")

	//构建hook的WriteSyncer列表
	var writeSyncerList []zapcore.WriteSyncer
	writeSyncerList = append(writeSyncerList, zapcore.AddSync(os.Stdout))
	if config.GlobalConfig.LogKafkaHookSwitch ||
		config.GlobalConfig.LogMongoHookSwitch {
		writeSyncerList = append(writeSyncerList, zapcore.AddSync(HookLogHook))
	}

	coreList = append(coreList,
		zapcore.NewCore(
			// 编码器配置
			zapcore.NewJSONEncoder(NewEncoderConfig()),
			// 打印到控制台和文件
			zapcore.NewMultiWriteSyncer(writeSyncerList...),
			// 日志级别
			debugLevel,
		))

	return
}

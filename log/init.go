package log

import (
	"apiproject/config"
	"go.uber.org/zap"
)

var Logger *zap.Logger
var HookLogger *zap.Logger

/**
初始化logger
*/
func Init() {
	Logger = GetCommonLogger()
	if config.GlobalConfig.LogKafkaHookSwitch ||
		config.GlobalConfig.LogMongoHookSwitch {
		HookLogger = GetHookLogger()
	}

	Logger.Info("初始化logger, 完成")
}

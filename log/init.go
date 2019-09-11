package log

import (
	"apiproject/config"
	"go.uber.org/zap"
)

var Logger *zap.Logger
var KafkaHookLogger *zap.Logger

func Init() {
	Logger = GetCommonLogger()
	if config.GlobalConfig.LogKafkaHookSwitch {
		KafkaHookLogger = GetKafkaHookLogger()
	}

	Logger.Info("创建kafka Producer完成", zap.Any("broker", config.GlobalConfig.KafkaBroker))
}

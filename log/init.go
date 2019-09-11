package log

import (
	"apiproject/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

var KafkaLogProducer *kafka.Producer
var Logger *zap.Logger
var KafkaHookLogger *zap.Logger

func Init() {
	var err error

	Logger = GetCommonLogger()
	if config.GlobalConfig.LogKafkaHookSwitch {
		KafkaHookLogger = GetKafkaHookLogger()
		KafkaLogProducer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.GlobalConfig.KafkaBroker})
		if err != nil {
			panic(err)
		}

		Logger.Info("创建kafka Producer, kafka log hook开启")
	}

	if err != nil {
		Logger.Error("创建kafka Producer异常", zap.Error(err))
	}

	Logger.Info("创建kafka Producer完成", zap.Any("broker", config.GlobalConfig.KafkaBroker))
}

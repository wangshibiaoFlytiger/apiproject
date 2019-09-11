package kafka

import (
	"apiproject/config"
	"apiproject/log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

var KafkaProducer *kafka.Producer

func Init() {
	var err error
	KafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.GlobalConfig.KafkaBroker})

	if err != nil {
		log.Logger.Error("创建kafka Producer异常", zap.Error(err))
	}

	log.Logger.Info("创建kafka Producer完成", zap.Any("broker", config.GlobalConfig.KafkaBroker))
}

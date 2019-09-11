package kafka

import (
	"apiproject/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

var KafkaProducer *kafka.Producer

/**
初始化kafka
*/
func Init() {
	if !config.GlobalConfig.KafkaSwitch {
		return
	}

	var err error
	KafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.GlobalConfig.KafkaBroker})

	if err != nil {
		log.Panicln("初始化kafka, 异常", err)
	}

	log.Println("初始化kafka, 完成", config.GlobalConfig.KafkaBroker)
}

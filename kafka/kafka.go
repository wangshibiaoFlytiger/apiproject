package kafka

import (
	"apiproject/config"
	"github.com/Shopify/sarama"
	"log"
	"strings"
)

var KafkaProducer sarama.SyncProducer

/**
初始化kafka
*/
func Init() {
	if !config.GlobalConfig.KafkaSwitch {
		return
	}

	var err error
	clientConfig := sarama.NewConfig()
	clientConfig.Producer.Return.Successes = true
	client, err := sarama.NewClient(strings.Split(config.GlobalConfig.KafkaBroker, ","), clientConfig)
	if err != nil {
		log.Panicf("初始化kafka, 创建client, 异常: %q", err)
	}

	KafkaProducer, err = sarama.NewSyncProducerFromClient(client)
	if err != nil {
		log.Panicf("初始化kafka, 创建生产者, 异常: %q", err)
	}

	log.Println("初始化kafka, 完成", config.GlobalConfig.KafkaBroker)
}

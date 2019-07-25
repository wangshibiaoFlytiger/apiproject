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

/**
发送kafka消息
*/
func SendKafkaMessage(topic string, message string) bool {
	success := false
	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	err := KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		log.Logger.Error("发送kafka消息异常", zap.Any("topic", topic), zap.Any("message", message), zap.Error(err))
	} else {
		log.Logger.Info("发送kafka消息完成", zap.Any("topic", *m.TopicPartition.Topic), zap.Any("partition", m.TopicPartition.Partition), zap.Any("offset", m.TopicPartition.Offset))
		success = true
	}

	close(deliveryChan)

	return success
}

package s_kafka

import (
	"apiproject/kafka"
	"apiproject/log"
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

var KafkaService = &kafkaService{}

type kafkaService struct {
}

/**
发送kafka消息
*/
func (this *kafkaService) SendKafkaMessage(topic string, message string) bool {
	partition, offset, err := kafka.KafkaProducer.SendMessage(&sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(message)})
	if err != nil {
		log.Logger.Error("发送kafka消息, 异常", zap.Error(err))
		return false
	}

	log.Logger.Info("发送kafka消息, 完成", zap.Any("topic", topic), zap.Any("message", message), zap.Any("partition", partition), zap.Any("offset", offset))
	return true
}

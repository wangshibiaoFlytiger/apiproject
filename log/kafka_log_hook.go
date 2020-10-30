package log

import (
	"apiproject/config"
	"apiproject/kafka"
	"errors"
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

type KafkaLogHook struct {
}

/**
实现KafkaLogHook的io.writer接口方法, 使当前对象可以作为zap的hook使用
*/
func (this *KafkaLogHook) Write(p []byte) (n int, err error) {
	if !this.SendKafkaMessage(config.GlobalConfig.LogKafkaTopic, string(p)) {
		HookLogger.Error("写kafka日志异常")
		return 0, errors.New("写kafka日志异常")
	}

	return
}

/**
发送kafka日志
*/
func (this *KafkaLogHook) SendKafkaMessage(topic string, message string) bool {
	partition, offset, err := kafka.KafkaProducer.SendMessage(&sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(message)})
	if err != nil {
		HookLogger.Error("发送kafka日志, 异常", zap.Any("topic", topic), zap.Any("message", message), zap.Error(err))
		return false
	}

	HookLogger.Info("发送kafka日志, 完成", zap.Any("topic", topic), zap.Any("message", message), zap.Any("partition", partition), zap.Any("offset", offset))
	return true
}

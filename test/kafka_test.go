package test

import (
	"github.com/Shopify/sarama"
	"log"
	"testing"
)

/**
测试sarama的kafka生产者
*/
func TestSaramaKafkaProducer(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	client, err := sarama.NewClient([]string{"172.18.100.222:6667", "172.18.100.223:6667", "172.18.100.224:6667"}, config)
	if err != nil {
		log.Fatalf("unable to create kafka client: %q", err)
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		log.Fatalf("unable to create kafka producer: %q", err)
	}
	defer producer.Close()

	topic := "test"
	text := "消息"
	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(text)})
	if err != nil {
		log.Fatalf("unable to produce message: %q", err)
	}

	println(partition, offset)
}

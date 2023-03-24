package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"strings"
)

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(brokers []string) *Producer {
	conf := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(brokers, ","),
		"acks":              "all",
		"compression.type":  "snappy",
	}
	producer, err := kafka.NewProducer(conf)
	if err != nil {
		logrus.Errorf("Failed to create producer: %s", err)
	}

	return &Producer{
		producer,
	}
}

func (p *Producer) SendMessage(topic string, message string) (bool, error) {
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}
	err := p.producer.Produce(msg, nil)
	if err != nil {
		logrus.Errorf("Failed to produce message: %v\n", err)
		return false, err
	}

	return true, nil
}

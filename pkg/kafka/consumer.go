package kafka

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/common"
	db "github.com/nguyenkhoa0721/go-project-layout/pkg/db/postgres/sqlc"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

func NewConsumer(ctx context.Context, brokers []string, topic string, callback func([]byte) error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        strings.Join(brokers, ","),
		"group.id":                 "myGroup_1",
		"session.timeout.ms":       6000,
		"auto.offset.reset":        "earliest",
		"enable.auto.offset.store": false,
	})

	if err != nil {
		logrus.Errorf("Failed to create consumer: %s", err)
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		logrus.Errorf("Failed to subscribe to topic: %s", err)
		os.Exit(1)
	}

	run := true

	for run {
		select {
		case <-ctx.Done():
			run = false
			break
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}
			switch e := ev.(type) {
			case *kafka.Message:
				isSuccess := false
				for i := 0; i < RETRY; i++ {
					err := callback(e.Value)
					if err != nil {
						logrus.Errorf("Retry %d/%d. Failed to process message: %s", i+1, RETRY, err)
						time.Sleep(RETRY_DELAY * time.Millisecond)
						continue
					}

					isSuccess = true
					break
				}

				if !isSuccess {
					common.GetCommon().Store.CreateKafkaDeadLetter(context.Background(), db.CreateKafkaDeadLetterParams{
						ID:    fmt.Sprint(common.GetCommon().Uuid.GenerateUuid(1)),
						Topic: topic,
						Value: string(e.Value),
					})
				}

				_, err = c.StoreMessage(e)
				if err != nil {
					logrus.Errorf("Failed to store offset: %s", err)
					run = false
					continue
				}
			case kafka.Error:
				logrus.Errorf("%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
					continue
				}
			}
		}
	}

	defer func() {
		logrus.Infof("Consumer: Closing %s consumer\n", topic)
		c.Close()
	}()
}

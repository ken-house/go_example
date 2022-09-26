package service

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go_example/internal/meta"
)

type KafkaService interface {
	Process(ctx context.Context) error
}

type kafkaService struct {
	kafkaConsumerClient meta.KafkaConsumerClient
}

func NewKafkaService(
	kafkaConsumerClient meta.KafkaConsumerClient,
) KafkaService {
	return &kafkaService{
		kafkaConsumerClient: kafkaConsumerClient,
	}
}

func (svc *kafkaService) Process(ctx context.Context) error {
	var ConsumerFunc = func(msg *sarama.ConsumerMessage) {
		fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
	return svc.kafkaConsumerClient.ConsumeTopic("first", sarama.OffsetOldest, ConsumerFunc)
}

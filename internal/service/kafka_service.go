package service

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go_example/internal/meta"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

type KafkaService interface {
	Process(ctx context.Context) error
	ProcessBatch(ctx context.Context)
}

type kafkaService struct {
	kafkaConsumerClient      meta.KafkaConsumerClient
	kafkaConsumerGroupClient meta.KafkaConsumerGroupClient
}

func NewKafkaService(
	kafkaConsumerClient meta.KafkaConsumerClient,
	kafkaConsumerGroupClient meta.KafkaConsumerGroupClient,
) KafkaService {
	return &kafkaService{
		kafkaConsumerClient:      kafkaConsumerClient,
		kafkaConsumerGroupClient: kafkaConsumerGroupClient,
	}
}

func (svc *kafkaService) Process(ctx context.Context) error {
	var consumerFunc = func(msg *sarama.ConsumerMessage) {
		fmt.Printf("from Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
	// 单个消费者
	//return svc.kafkaConsumerClient.ConsumeTopic("second", sarama.OffsetOldest, consumerFunc)
	// 消费者组
	return svc.kafkaConsumerGroupClient.ConsumeTopic(ctx, []string{"second"}, consumerFunc)
}

func (svc *kafkaService) ProcessBatch(ctx context.Context) {
	workChannel := make(chan string, 100)
	var consumerFunc = func(msg *sarama.ConsumerMessage) {
		// 将数据写入到channel中
		message := fmt.Sprintf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		workChannel <- message
	}

	// 启动3个消费者将kafka数据写入到workChannel
	pool, _ := ants.NewPool(3)
	defer pool.Release()
	pool.Submit(func() {
		svc.kafkaConsumerGroupClient.ConsumeTopic(ctx, []string{"second"}, consumerFunc)
	})

	// 启动一个goroutine读取channel，每50个消息处理一次
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		messageList := make([]string, 0, 100)
		tickTimer := time.NewTicker(10 * time.Second)
		defer wg.Done()
		for {
			select {
			case <-ctx.Done(): // 程序终止时
				handlerFunc(messageList)
				return
			case msg, ok := <-workChannel: // 从channel中读取数据
				if !ok { // channel为空时
					handlerFunc(messageList)
					return
				}
				messageList = append(messageList, msg)
				if len(messageList) >= 50 { // 每50个处理一次
					handlerFunc(messageList)
					// 清空messageList
					messageList = messageList[0:0]
				}
			case <-tickTimer.C: // 间隔时间到时
				handlerFunc(messageList)
				// 清空messageList
				messageList = messageList[0:0]
			}
		}
	}()
}

func handlerFunc(messageList []string) {
	if len(messageList) > 0 {
		fmt.Printf("messageList:%v\n\n", messageList)
	}
}

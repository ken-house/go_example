package service

import (
	"context"
	"fmt"
	"github.com/go_example/internal/meta"
	"log"
)

type RabbitmqService interface {
	DeclareAndBind(ctx context.Context, exchangeName string, exchangeType string, queueName string, bindingKey string) error
	Process(ctx context.Context, queueName string, consumerName string)
}

type rabbitmqService struct {
	rabbitmqClient meta.RabbitmqClient
}

func NewRabbitmqService(rabbitmqClient meta.RabbitmqClient) RabbitmqService {
	return &rabbitmqService{
		rabbitmqClient: rabbitmqClient,
	}
}

// DeclareAndBind 声明队列或交换机并进行绑定
func (svc *rabbitmqService) DeclareAndBind(ctx context.Context, exchangeName string, exchangeType string, queueName string, bindingKey string) error {
	// 1、声明一个交换机
	err := svc.rabbitmqClient.ExchangeDeclare(ctx, exchangeName, exchangeType, false, false, false, false, nil)
	if err != nil {
		return err
	}

	// 2、声明一个队列
	_, err = svc.rabbitmqClient.QueueDeclare(ctx, queueName, false, false, false, false, nil)
	if err != nil {
		return err
	}

	// 3、绑定队列到交换机
	return svc.rabbitmqClient.QueueBind(ctx, queueName, bindingKey, exchangeName, false, nil)
}

// Process 启动消费
func (svc *rabbitmqService) Process(ctx context.Context, queueName string, consumerName string) {
	deliveries, err := svc.rabbitmqClient.Consume(ctx, queueName, consumerName, true, false, false, false, nil)
	if err != nil {
		log.Fatalln("Consume err:", err.Error())
	}

	// 启动协程从channel中读取消息
	go func() {
		for d := range deliveries {
			fmt.Println(d.DeliveryTag, string(d.Body))
		}
	}()
}

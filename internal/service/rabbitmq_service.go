package service

import (
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type RabbitmqService interface {
	RabbitmqMessage(ctx *gin.Context, message string, args map[string]interface{}) amqp.Publishing
	Producer(ctx *gin.Context, channel *amqp.Channel, exchangeName string, bindingKey string, mandatory bool, immediate bool, message amqp.Publishing) error
}

type rabbitmqService struct {
}

func NewRabbitmqService() RabbitmqService {
	return &rabbitmqService{}
}

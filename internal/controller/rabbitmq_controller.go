package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/meta"
	"github.com/go_example/internal/utils/negotiate"
	"github.com/ken-house/go-contrib/prototype/errorAssets"
	"github.com/streadway/amqp"
	"net/http"
)

type RabbitmqController interface {
	Producer(ctx *gin.Context) (int, gin.Negotiate)
}

type rabbitmqController struct {
	rabbitmqClient meta.RabbitmqClient
}

func NewRabbitmqController(rabbitmqClient meta.RabbitmqClient) RabbitmqController {
	return &rabbitmqController{
		rabbitmqClient: rabbitmqClient,
	}
}

func (ctr *rabbitmqController) Producer(ctx *gin.Context) (int, gin.Negotiate) {
	channel := ctr.rabbitmqClient.GetChannel(ctx)

	message := "你好啊"
	publishMsg := amqp.Publishing{
		Body: []byte(message),
	}
	err := channel.Publish("hello_exchange", "hello", false, false, publishMsg)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_SYSTEM.ToastError())
	}
	return negotiate.JSON(http.StatusOK, gin.H{
		"data": "发送成功",
	})
}

package controller

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/lib/errorAssets"
	"github.com/go_example/internal/meta"
	"github.com/go_example/internal/utils/negotiate"
	"go.uber.org/zap"
	"net/http"
)

type KafkaController interface {
	ProducerSync(ctx *gin.Context) (int, gin.Negotiate)
	ProducerAsync(ctx *gin.Context) (int, gin.Negotiate)
}

type kafkaController struct {
	kafkaProducerSyncClient  meta.KafkaProducerSyncClient
	kafkaProducerAsyncClient meta.KafkaProducerAsyncClient
	kafkaConsumerClient      meta.KafkaConsumerClient
}

func NewKafkaController(
	kafkaProducerSyncClient meta.KafkaProducerSyncClient,
	kafkaProducerAsyncClient meta.KafkaProducerAsyncClient,
) KafkaController {
	return &kafkaController{
		kafkaProducerSyncClient:  kafkaProducerSyncClient,
		kafkaProducerAsyncClient: kafkaProducerAsyncClient,
	}
}

// ProducerSync 同步发送
func (ctr kafkaController) ProducerSync(ctx *gin.Context) (int, gin.Negotiate) {
	message := "hello one message"
	msg := &sarama.ProducerMessage{
		Topic: "second",
		Value: sarama.StringEncoder(message),
	}
	// 单条消息发送
	partition, offset, err := ctr.kafkaProducerSyncClient.SendMessage(msg)
	fmt.Println(partition, offset, err)
	// 单条消息发送
	partition, offset, err = ctr.kafkaProducerSyncClient.SendOne("second", "", message, 0)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_SYSTEM.ToastError())
	}
	fmt.Println(partition, offset, err)
	messageList := []string{
		"hello world sync",
		"hi",
	}
	// 多条消息发送
	err = ctr.kafkaProducerSyncClient.SendMany("second", "", messageList, 0)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_SYSTEM.ToastError())
	}

	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"message": "OK",
		},
	})
}

// ProducerAsync 异步发送
func (ctr kafkaController) ProducerAsync(ctx *gin.Context) (int, gin.Negotiate) {
	for i := 0; i <= 20; i++ {
		message := "hello world async cc"
		err := ctr.kafkaProducerAsyncClient.SendOne("second", "", message, 0)
		if err != nil {
			zap.L().Error("kafkaController.ProducerAsync err", zap.Error(err))
			return negotiate.JSON(http.StatusOK, errorAssets.ERR_SYSTEM.ToastError())
		}
		//time.Sleep(time.Second)
	}

	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"message": "OK",
		},
	})
}

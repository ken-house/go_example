package service

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HelloService interface {
	SayHello(c *gin.Context) map[string]string
}

type helloService struct {
}

func NewHelloService() HelloService {
	return &helloService{}
}

func (svc *helloService) SayHello(c *gin.Context) map[string]string {
	return map[string]string{
		"hello": "world，golang",
		"env":   viper.GetString("server.mode"),
	}
}

package service

import "github.com/gin-gonic/gin"

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
	}
}

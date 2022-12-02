package controller

import (
	"github.com/ken-house/go-contrib/prototype/errorAssets"
	"net/http"

	"github.com/go_example/internal/utils/negotiate"

	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/service"
)

type HelloController interface {
	Say(ctx *gin.Context) (int, gin.Negotiate)
	Email(ctx *gin.Context) (int, gin.Negotiate)
}

type helloController struct {
	helloSvc service.HelloService
	emailSvc service.EmailService
}

func NewHelloController(
	helloSvc service.HelloService,
	emailSvc service.EmailService,
) HelloController {
	return &helloController{
		helloSvc: helloSvc,
		emailSvc: emailSvc,
	}
}

func (ctr *helloController) Say(ctx *gin.Context) (int, gin.Negotiate) {
	//time.Sleep(5 * time.Second)
	data := ctr.helloSvc.SayHello(ctx)
	return negotiate.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (ctr *helloController) Email(ctx *gin.Context) (int, gin.Negotiate) {
	err := ctr.emailSvc.Send(ctx)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_SYSTEM.ToastError())
	}
	return negotiate.JSON(http.StatusOK, gin.H{
		"data": "发送成功",
	})
}

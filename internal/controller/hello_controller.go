package controller

import (
	"github.com/go_example/internal/meta"
	"github.com/spf13/cast"
	"net/http"

	"github.com/go_example/internal/utils/negotiate"

	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/service"
)

type HelloController interface {
	Say(ctx *gin.Context) (int, gin.Negotiate)
}

type helloController struct {
	helloSvc service.HelloService
}

func NewHelloController(
	helloSvc service.HelloService,
) HelloController {
	return &helloController{
		helloSvc: helloSvc,
	}
}

func (ctr *helloController) Say(ctx *gin.Context) (int, gin.Negotiate) {
	newCtx, span := meta.HttpTracer.Start(ctx.Request.Context(), "helloController_Say")
	defer span.End()

	uid := cast.ToInt(ctx.DefaultQuery("id", "0"))

	//time.Sleep(5 * time.Second)
	data := ctr.helloSvc.SayHello(newCtx, uid)
	return negotiate.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

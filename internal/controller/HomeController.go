package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/utils/negotiate"
)

type HomeController interface {
	Index(*gin.Context) (int, gin.Negotiate)
}

type homeController struct {
}

func NewHomeController() HomeController {
	return &homeController{}
}

func (ctr *homeController) Index(ctx *gin.Context) (int, gin.Negotiate) {
	userInfo, _ := ctx.Get("userInfo")
	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"text":     "welcome index page",
			"userInfo": userInfo,
		},
	})
}

package controller

import (
	"net/http"

	"github.com/go_example/internal/utils/negotiate"

	"github.com/go_example/internal/model"

	"github.com/go_example/internal/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(*gin.Context) (int, gin.Negotiate)
}

type loginController struct {
	loginSvc service.LoginService
}

func NewLoginController(loginSvc service.LoginService) LoginController {
	return &loginController{
		loginSvc: loginSvc,
	}
}

func (ctr *loginController) Login(c *gin.Context) (int, gin.Negotiate) {
	var paramData model.LoginForm
	if err := c.ShouldBind(&paramData); err != nil {
		return negotiate.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"token":   "",
				"message": err.Error(),
			},
		})
	}

	// 登录验证
	token, err := ctr.loginSvc.Login(paramData)
	if err != nil {
		return negotiate.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"token":   "",
				"message": err.Error(),
			},
		})
	}

	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"token":   token,
			"message": "OK",
		},
	})
}

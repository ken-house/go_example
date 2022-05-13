package controller

import (
	"net/http"
	"strings"

	"github.com/go_example/common/auth"

	"github.com/go_example/internal/utils/negotiate"

	"github.com/go_example/internal/model"

	"github.com/go_example/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(*gin.Context) (int, gin.Negotiate)
	RefreshToken(*gin.Context) (int, gin.Negotiate)
}

type authController struct {
	authSvc service.AuthService
}

func NewAuthController(authSvc service.AuthService) AuthController {
	return &authController{
		authSvc: authSvc,
	}
}

func (ctr *authController) Login(c *gin.Context) (int, gin.Negotiate) {
	var paramData model.LoginForm
	if err := c.ShouldBind(&paramData); err != nil {
		return negotiate.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"accessToken":  "",
				"refreshToken": "",
				"message":      err.Error(),
			},
		})
	}

	// 登录验证
	accessToken, refreshToken, err := ctr.authSvc.Login(paramData)
	if err != nil {
		return negotiate.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"accessToken":  "",
				"refreshToken": "",
				"message":      err.Error(),
			},
		})
	}

	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
			"message":      "OK",
		},
	})
}

func (ctr *authController) RefreshToken(c *gin.Context) (int, gin.Negotiate) {
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		return negotiate.JSON(http.StatusOK, gin.H{"message": "令牌为空"})
	}

	parts := strings.SplitN(authorization, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return negotiate.JSON(http.StatusOK, gin.H{"message": "令牌格式错误"})
	}

	claims, err := auth.ParseToken(parts[1], "refresh_token")
	if err != nil {
		return negotiate.JSON(http.StatusOK, gin.H{"message": "刷新令牌失败，请重新登录"})
	}

	// 重新生成令牌
	userInfo, err := ctr.authSvc.GetUserInfo(claims.UserInfo.Id)
	if err != nil {
		return negotiate.JSON(http.StatusOK, gin.H{"message": "用户信息错误，请重新登录"})
	}
	accessToken, refreshToken, err := auth.GenToken(userInfo)
	if err != nil {
		return negotiate.JSON(http.StatusOK, gin.H{"message": "生成令牌失败，请重新登录"})
	}

	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
			"message":      "OK",
		},
	})
}

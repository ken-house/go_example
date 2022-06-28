package controller

import (
	"net/http"
	"strings"

	"github.com/go_example/common/errorAssets"

	"github.com/go_example/internal/lib/auth"

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
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_PARAM.ToastError())
	}

	userInfo, err := ctr.authSvc.FindIdentity(paramData)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_LOGIN.ToastError())
	}

	// 登录验证
	accessToken, refreshToken, err := auth.GenToken(ctr.authSvc, userInfo)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_SYSTEM.ToastError())
	}

	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}

func (ctr *authController) RefreshToken(c *gin.Context) (int, gin.Negotiate) {
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_PARAM.ToastError())
	}

	parts := strings.SplitN(authorization, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_PARAM.ToastError())
	}

	claims, err := auth.ParseToken(ctr.authSvc, parts[1], "refresh_token")
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_REFRESH_TOKEN.ToastError())
	}

	// 重新生成令牌
	userInfo, err := ctr.authSvc.GetUserInfo(claims.UserInfo.Id)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_REFRESH_TOKEN.ToastError())
	}
	accessToken, refreshToken, err := auth.GenToken(ctr.authSvc, userInfo)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_REFRESH_TOKEN.ToastError())
	}

	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}

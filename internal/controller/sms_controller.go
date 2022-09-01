package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/lib/errorAssets"
	"github.com/go_example/internal/model"
	"github.com/go_example/internal/service"
	"github.com/go_example/internal/utils/negotiate"
	"net/http"
)

type SmsController interface {
	SendCode(ctx *gin.Context) (int, gin.Negotiate)
}

type smsController struct {
	smsSvc service.SmsService
}

func NewSmsController(smsSvc service.SmsService) SmsController {
	return &smsController{
		smsSvc: smsSvc,
	}
}

// SendCode 发送短信验证码
func (ctr *smsController) SendCode(ctx *gin.Context) (int, gin.Negotiate) {
	var params model.SendPhoneCodeForm
	if err := ctx.ShouldBind(&params); err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_PARAM.ToastError())
	}
	// 发送短信验证码
	code, err := ctr.smsSvc.SendCode(ctx, params.Phone)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_SMS_SEND_FAIL.ToastError())
	}

	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"code": code,
		},
	})
}

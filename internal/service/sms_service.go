package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/meta"
	"github.com/go_example/internal/utils/tools"
	"github.com/ken-house/go-contrib/prototype/alibabaSmsClient"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type SmsService interface {
	SendCode(ctx *gin.Context, phone string) (code string, err error)
}

type smsService struct {
	smsClient meta.AlibabaSmsClient
}

func NewSmsService(
	smsClient meta.AlibabaSmsClient,
) SmsService {
	return &smsService{
		smsClient: smsClient,
	}
}

// SendCode 发送短信验证码
func (svc *smsService) SendCode(ctx *gin.Context, phone string) (code string, err error) {
	var params alibabaSmsClient.SendSmsParams
	if err := viper.Sub("alibaba_sms_code").Unmarshal(&params); err != nil {
		return "", err
	}
	params.Phone = phone
	code = tools.GetRandomString(6, 1)
	templateParam := struct {
		Code string `json:"code"`
	}{
		Code: code,
	}
	templateParamByte, err := json.Marshal(templateParam)
	if err != nil {
		return "", err
	}
	params.TemplateParam = string(templateParamByte)
	err = svc.smsClient.SendCode(params)
	if err != nil {
		zap.L().Error("smsService.SendCode err", zap.Error(err), zap.String("phone", params.Phone))
		return "", err
	}
	return code, nil
}

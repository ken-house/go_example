package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/meta"
	"github.com/go_example/internal/utils/tools"
	"github.com/ken-house/go-contrib/prototype/alibabaSmsClient"
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
	params := alibabaSmsClient.SendSmsParams{
		Phone:        phone,
		SignName:     meta.GlobalConfig.AlibabaSmsCode.SignName,
		TemplateCode: meta.GlobalConfig.AlibabaSmsCode.TemplateCode,
	}
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

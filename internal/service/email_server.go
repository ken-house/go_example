package service

import (
	"context"
	"github.com/go_example/internal/meta"
	"github.com/ken-house/go-contrib/prototype/emailClient"
)

type EmailService interface {
	Send(ctx context.Context) error
}

type emailService struct {
	emailClient meta.EmailClient
}

func NewEmailService(emailClient meta.EmailClient) (EmailService, func(), error) {
	return &emailService{
		emailClient: emailClient,
	}, nil, nil
}

func (svc *emailService) Send(ctx context.Context) error {
	userList := []emailClient.EmailUser{
		{
			Name:         "Ken",
			EmailAddress: "xudengtang@zonst.cn",
		},
	}

	message := emailClient.EmailMessage{
		Subject:     "测试邮件",
		ContentType: "text/html",
		Body:        "<h1 style='color:#F00'>hello world</h1>",
		AttachFilePathList: []string{
			"/Users/zonst/Downloads/WechatIMG51.jpeg",
			"/Users/zonst/Downloads/download.zip",
		},
	}
	err := svc.emailClient.Send(userList, message)
	if err != nil {
		meta.SentryClient.CaptureException(err)
	}
	return err
}

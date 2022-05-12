package service

import (
	"github.com/go_example/common/auth"
	"github.com/go_example/internal/model"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	"github.com/pkg/errors"
)

type LoginService interface {
	Login(model.LoginForm) (string, error)
}

type loginService struct {
	UserRepo MysqlRepo.UserRepository
}

func NewLoginService(
	userRepo MysqlRepo.UserRepository,
) LoginService {
	return &loginService{
		UserRepo: userRepo,
	}
}

func (svc *loginService) Login(formData model.LoginForm) (token string, err error) {
	// 验证用户登录
	userInfo, err := svc.UserRepo.FindIdentity(formData)
	if err != nil {
		return "", errors.New("用户名密码不正确")
	}

	// 登录成功则生成Token
	return auth.GenToken(userInfo)
}

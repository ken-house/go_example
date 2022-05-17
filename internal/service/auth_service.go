package service

import (
	"github.com/go_example/internal/lib/auth"
	"github.com/go_example/internal/model"
	MysqlModel "github.com/go_example/internal/model/mysql"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	"github.com/pkg/errors"
)

type AuthService interface {
	Login(model.LoginForm) (string, string, error)
	GetUserInfo(int) (MysqlModel.User, error)
}

type authService struct {
	UserRepo MysqlRepo.UserRepository
}

func NewAuthService(
	userRepo MysqlRepo.UserRepository,
) AuthService {
	return &authService{
		UserRepo: userRepo,
	}
}

func (svc *authService) Login(formData model.LoginForm) (accessToken string, refreshToken string, err error) {
	// 验证用户登录
	userInfo, err := svc.UserRepo.FindIdentity(formData)
	if err != nil {
		return "", "", errors.New("用户名密码不正确")
	}

	// 登录成功则生成Token
	return auth.GenToken(userInfo)
}

func (svc *authService) GetUserInfo(userId int) (MysqlModel.User, error) {
	return svc.UserRepo.GetUserInfoById(userId)
}

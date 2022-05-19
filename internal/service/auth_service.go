package service

import (
	"github.com/go_example/internal/model"
	MysqlModel "github.com/go_example/internal/model/mysql"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	RedisRepo "github.com/go_example/internal/repository/redis"
	"github.com/pkg/errors"
)

type AuthService interface {
	FindIdentity(model.LoginForm) (MysqlModel.User, error)
	GetUserInfo(int) (MysqlModel.User, error)
	CheckAuthTokenRedis(int, string, string) error
	SaveAuthTokenRedis(int, string, string) error
}

type authService struct {
	UserMysqlRepo MysqlRepo.UserRepository
	UserRedisRepo RedisRepo.UserRepository
}

func NewAuthService(
	userMysqlRepo MysqlRepo.UserRepository,
	userRedisRepo RedisRepo.UserRepository,
) AuthService {
	return &authService{
		UserMysqlRepo: userMysqlRepo,
		UserRedisRepo: userRedisRepo,
	}
}

func (svc *authService) FindIdentity(formData model.LoginForm) (MysqlModel.User, error) {
	return svc.UserMysqlRepo.GetUserInfoByFormData(formData)
}

func (svc *authService) GetUserInfo(userId int) (MysqlModel.User, error) {
	return svc.UserMysqlRepo.GetUserInfoById(userId)
}

// CheckAuthTokenRedis 单点登录，检查用户认证redis中token是否为有效
func (svc *authService) CheckAuthTokenRedis(userId int, token string, grantType string) (err error) {
	value, err := svc.UserRedisRepo.GetUserAuthToken(userId, grantType)
	if err != nil {
		return err
	}
	if token != value {
		return errors.New("账号已在其他设备登录")
	}
	return nil
}

// SaveAuthTokenRedis 单点登录，保存token到用户的认证redis
func (svc *authService) SaveAuthTokenRedis(userId int, accessTokenSign string, refreshTokenSign string) error {
	tokenMap := make(map[string]string, 2)
	tokenMap["access_token"] = accessTokenSign
	tokenMap["refresh_token"] = refreshTokenSign
	return svc.UserRedisRepo.SetUserAuthToken(userId, tokenMap)
}

package service

import (
	"github.com/gin-gonic/gin"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	RedisRepo "github.com/go_example/internal/repository/redis"
	"github.com/spf13/viper"
)

type HelloService interface {
	SayHello(c *gin.Context) map[string]string
}

type helloService struct {
	UserRepo      MysqlRepo.UserRepository
	UserRedisRepo RedisRepo.UserRepository
}

func NewHelloService(
	userRepo MysqlRepo.UserRepository,
	userRedisRepo RedisRepo.UserRepository,
) HelloService {
	return &helloService{
		UserRepo:      userRepo,
		UserRedisRepo: userRedisRepo,
	}
}

func (svc *helloService) SayHello(c *gin.Context) map[string]string {
	user, _ := svc.UserRepo.GetUserInfo(1)
	value := svc.UserRedisRepo.GetValue("aa")
	return map[string]string{
		"hello": "world，golang",
		"env":   viper.GetString("server.mode"),
		"user":  user.Username,
		"value": value,
	}
}

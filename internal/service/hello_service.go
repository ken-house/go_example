package service

import (
	"github.com/gin-gonic/gin"
	CacheRepo "github.com/go_example/internal/repository/cache"
	MongoRepo "github.com/go_example/internal/repository/mongodb"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	RedisRepo "github.com/go_example/internal/repository/redis"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type HelloService interface {
	SayHello(c *gin.Context) map[string]string
}

type helloService struct {
	userRepo      MysqlRepo.UserRepository
	userRedisRepo RedisRepo.UserRepository
	userCacheRepo CacheRepo.UserRepository
	userMongoRepo MongoRepo.UserRepository
}

func NewHelloService(
	userRepo MysqlRepo.UserRepository,
	userRedisRepo RedisRepo.UserRepository,
	userCacheRepo CacheRepo.UserRepository,
	userMongoRepo MongoRepo.UserRepository,
) HelloService {
	return &helloService{
		userRepo:      userRepo,
		userRedisRepo: userRedisRepo,
		userCacheRepo: userCacheRepo,
		userMongoRepo: userMongoRepo,
	}
}

func (svc *helloService) SayHello(c *gin.Context) map[string]string {
	uid := cast.ToInt(c.DefaultQuery("id", "0"))
	// 先从gocache中读取数据
	user, err := svc.userCacheRepo.GetUserInfo(uid)
	if err != nil {
		user, err = svc.userRepo.GetUserInfoById(uid)
		if err == nil {
			// 写入数据到gocache中
			_ = svc.userCacheRepo.SetUserInfo(uid, user)
		}
	}

	// redis操作数据
	value := svc.userRedisRepo.GetValue("aa")

	// mongodb操作数据
	if user.Username != "" {
		_ = svc.userMongoRepo.SetUserInfo(uid, user.Username, user.Password)
	}
	userInfo, _ := svc.userMongoRepo.GetUserInfo(uid)

	// 测试
	zap.L().Debug("this is test")

	return map[string]string{
		"hello":          "world，golang",
		"env":            viper.GetString("server.mode"),
		"user":           user.Username,
		"value":          value,
		"mongo_username": userInfo.Username,
	}
}

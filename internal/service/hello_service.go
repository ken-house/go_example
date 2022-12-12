package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/meta"
	CacheRepo "github.com/go_example/internal/repository/cache"
	MongoRepo "github.com/go_example/internal/repository/mongodb"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	RedisRepo "github.com/go_example/internal/repository/redis"
	"github.com/ken-house/go-contrib/prototype/requester"
	"github.com/ken-house/go-contrib/utils/encrypt"
	"github.com/ken-house/go-contrib/utils/tools"
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
	//zap.L().Debug("this is test")
	fmt.Println(meta.GlobalConfig.Redis.Single.PoolSize)

	// http请求
	responseData := struct {
		Hello string `json:"hello"`
	}{}
	httpClient := requester.NewRequestClient("http://127.0.0.1:8081", nil, nil)
	response, err := httpClient.Get(c, "/hello", &responseData, nil)
	if err != nil {
		zap.L().Error("请求失败", zap.Error(err))
		meta.SentryClient.CaptureExceptionForGin(c, err)
	}
	fmt.Println(response)
	fmt.Printf("responseData：%+v\n", responseData)

	hello := "world，golang"
	iv := tools.GenerateRandStr(16, 3)
	encryptStr, _ := encrypt.AesEncrypt(hello, meta.GlobalConfig.Common.AesKey, iv)
	decryptStr, _ := encrypt.AesDecrypt(encryptStr, meta.GlobalConfig.Common.AesKey, iv)

	return map[string]string{
		"hello":          hello,
		"encryptStr":     encryptStr,
		"decryptStr":     decryptStr,
		"env":            viper.GetString("server.mode"),
		"user":           user.Username,
		"value":          value,
		"mongo_username": userInfo.Username,
		"version":        "2.0",
	}
}

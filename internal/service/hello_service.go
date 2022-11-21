package service

import (
	"context"
	"fmt"
	"github.com/go_example/internal/meta"
	CacheRepo "github.com/go_example/internal/repository/cache"
	MongoRepo "github.com/go_example/internal/repository/mongodb"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	RedisRepo "github.com/go_example/internal/repository/redis"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/attribute"
)

type HelloService interface {
	SayHello(ctx context.Context, uid int) map[string]string
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

func (svc *helloService) SayHello(ctx context.Context, uid int) map[string]string {
	newCtx, span := meta.HttpTracer.Start(ctx, "helloService_SayHello")
	defer span.End()
	span.SetAttributes(attribute.Int("uid", uid))

	// 先从gocache中读取数据
	user, err := svc.userCacheRepo.GetUserInfo(newCtx, uid)
	if err != nil {
		user, err = svc.userRepo.GetUserInfoById(newCtx, uid)
		if err == nil {
			// 写入数据到gocache中
			_ = svc.userCacheRepo.SetUserInfo(newCtx, uid, user)
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
	//responseData := struct {
	//	Hello string `json:"hello"`
	//}{}
	//httpClient := requester.NewRequestClient("http://127.0.0.1:8081", nil, nil)
	//response, err := httpClient.Get(ctx, "/hello", &responseData, nil)
	//if err != nil {
	//	zap.L().Error("请求失败", zap.Error(err))
	//}
	//fmt.Println(response)
	//fmt.Printf("responseData：%+v\n", responseData)

	return map[string]string{
		"hello":          "world，golang",
		"env":            viper.GetString("server.mode"),
		"user":           user.Username,
		"value":          value,
		"mongo_username": userInfo.Username,
		"version":        "2.0",
	}
}

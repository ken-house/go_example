package service

import (
	"context"
	"fmt"
	"github.com/go_example/internal/meta"
	CacheRepo "github.com/go_example/internal/repository/cache"
	MongoRepo "github.com/go_example/internal/repository/mongodb"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	RedisRepo "github.com/go_example/internal/repository/redis"
	"github.com/ken-house/go-contrib/prototype/requester"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"net/http"
)

type HelloService interface {
	SayHello(ctx context.Context, uid int) map[string]string
	Request(ctx context.Context) string
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

	return map[string]string{
		"hello":          "world，golang",
		"env":            viper.GetString("server.mode"),
		"user":           user.Username,
		"value":          value,
		"mongo_username": userInfo.Username,
		"version":        "2.0",
	}
}

// Request 模拟HTTP请求
func (svc *helloService) Request(ctx context.Context) string {
	newCtx, span := meta.HttpTracer.Start(ctx, "helloService_Request")
	defer span.End()

	// http请求
	responseData := struct {
		Data struct {
			Hello string `json:"hello"`
		} `json:"data"`
	}{}
	func(ctx context.Context) {
		newCtx, span = meta.HttpTracer.Start(ctx, "requester.httpClient.Get", trace.WithAttributes(semconv.PeerServiceKey.String("go_example")))
		defer span.End()

		client := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
		httpClient := requester.NewRequestClient("http://127.0.0.1:8080", nil, client)
		_, err := httpClient.Get(newCtx, "/hello", &responseData, nil)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, "request failed")
			zap.L().Error("请求失败", zap.Error(err))
		}
	}(newCtx)
	return responseData.Data.Hello
}

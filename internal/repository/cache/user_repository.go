package cache

import (
	"context"
	"encoding/json"
	"errors"
	"go.opentelemetry.io/otel/attribute"
	"time"

	"github.com/go_example/internal/meta"
	MysqlModel "github.com/go_example/internal/model/mysql"
)

type UserRepository interface {
	SetUserInfo(ctx context.Context, uid int, userInfo MysqlModel.User) error
	GetUserInfo(ctx context.Context, uid int) (MysqlModel.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (repo *userRepository) SetUserInfo(ctx context.Context, uid int, userInfo MysqlModel.User) error {
	_, span := meta.HttpTracer.Start(ctx, "cache_userRepository_SetUserInfo")
	defer span.End()

	cacheKey := GetCacheKey(UserInfoKey, uid)
	userInfoByte, _ := json.Marshal(userInfo)
	span.SetAttributes(attribute.String("cacheKey", cacheKey), attribute.String("userInfo", string(userInfoByte)))
	meta.CacheDriver.Set(cacheKey, userInfo, time.Hour)
	return nil
}

func (repo *userRepository) GetUserInfo(ctx context.Context, uid int) (MysqlModel.User, error) {
	_, span := meta.HttpTracer.Start(ctx, "cache_userRepository_GetUserInfo")
	defer span.End()

	cacheKey := GetCacheKey(UserInfoKey, uid)
	data, isExist := meta.CacheDriver.Get(cacheKey)
	span.SetAttributes(attribute.String("cacheKey", cacheKey), attribute.Bool("existCacheKey", isExist))
	if !isExist {
		return MysqlModel.User{}, errors.New("key不存在")
	}
	return data.(MysqlModel.User), nil
}

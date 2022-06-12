package redis

import (
	"context"
	"time"

	"github.com/go_example/internal/meta"
)

type UserRepository interface {
	GetValue(key string) string
	GetUserAuthToken(int, string) (string, error)
	SetUserAuthToken(int, map[string]string) error
}

type userRepository struct {
	redisClient meta.RedisSingleClient
}

func NewUserRepository(redisClient meta.RedisSingleClient) UserRepository {
	return &userRepository{
		redisClient: redisClient,
	}
}

func (repo *userRepository) GetValue(key string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	value, _ := repo.redisClient.Get(ctx, key).Result()
	return value
}

// GetUserAuthToken 获取用户令牌
func (repo *userRepository) GetUserAuthToken(userId int, grantType string) (token string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	key := GetRedisKey(AuthTokenKey, userId)
	return repo.redisClient.HGet(ctx, key, grantType).Result()
}

// SetUserAuthToken 设置用户令牌
func (repo *userRepository) SetUserAuthToken(userId int, tokenMap map[string]string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	key := GetRedisKey(AuthTokenKey, userId)
	_, err = repo.redisClient.HMSet(ctx, key, tokenMap).Result()
	return
}

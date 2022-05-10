package redis

import (
	"context"
	"time"

	"github.com/go_example/internal/meta"
)

type UserRepository interface {
	GetValue(key string) string
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

package redisClient

import (
	"context"

	"github.com/pkg/errors"

	"github.com/go-redis/redis/v8"
)

type GroupClient interface {
	redis.UniversalClient
}

type groupClient struct {
	redis.UniversalClient
}

func NewGroupClient(cfg GroupConfig) (GroupClient, func(), error) {
	client, err := NewRedisGroupClient(cfg)
	if err != nil {
		return nil, nil, err
	}
	cli := &groupClient{UniversalClient: client}
	return cli, func() {
		client.Close()
	}, nil
}

type GroupConfig struct {
	Addrs    []string `json:"addrs" mapstructure:"addrs"`
	Password string
}

func NewRedisGroupClient(cfg GroupConfig) (redis.UniversalClient, error) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Addrs,
		Password: cfg.Password,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return client, err
}

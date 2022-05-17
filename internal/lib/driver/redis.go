package driver

import (
	"github.com/go_example/common/redisClient"
	"github.com/go_example/internal/meta"
	"github.com/spf13/viper"
)

// GetRedisSingleClient 连接Redis单机
func GetRedisSingleClient() (meta.RedisSingleClient, func(), error) {
	var cfg redisClient.SingleConfig
	if err := viper.Sub("redis." + meta.RedisSingleDriverKey).Unmarshal(&cfg); err != nil {
		return nil, nil, err
	}
	return redisClient.NewSingleClient(cfg)
}

// GetRedisGroupClient 连接RedisCluster集群
func GetRedisGroupClient() (meta.RedisGroupClient, func(), error) {
	var cfg redisClient.GroupConfig
	if err := viper.Sub("redis." + meta.RedisGroupDriverKey).Unmarshal(&cfg); err != nil {
		return nil, nil, err
	}
	return redisClient.NewGroupClient(cfg)
}

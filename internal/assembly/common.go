package assembly

import (
	"github.com/go_example/internal/meta"
	"github.com/ken-house/go-contrib/prototype/consulClient"
	"github.com/ken-house/go-contrib/prototype/jenkinsClient"
	"github.com/ken-house/go-contrib/prototype/mongoClient"
	"github.com/ken-house/go-contrib/prototype/mysqlClient"
	"github.com/ken-house/go-contrib/prototype/redisClient"
	"github.com/ken-house/go-contrib/utils/env"
	"github.com/spf13/viper"
)

// NewMysqlSingleClient 单机数据库连接
func NewMysqlSingleClient() (meta.MysqlSingleClient, func(), error) {
	var cfg mysqlClient.SingleConfig
	if err := viper.Sub("mysql." + meta.MysqlSingleDriverKey).Unmarshal(&cfg); err != nil {
		return nil, nil, err
	}
	return mysqlClient.NewSingleClient(cfg)
}

// NewMysqlGroupClient 主从数据库连接
func NewMysqlGroupClient() (meta.MysqlGroupClient, func(), error) {
	var cfg mysqlClient.GroupConfig
	if err := viper.Sub("mysql." + meta.MysqlGroupDriverKey).Unmarshal(&cfg); err != nil {
		return nil, nil, err
	}
	cfg.IsDebug = !env.IsReleasing()
	return mysqlClient.NewGroupClient(cfg)
}

// NewRedisSingleClient 连接Redis单机
func NewRedisSingleClient() (meta.RedisSingleClient, func(), error) {
	var cfg redisClient.RedisConfig
	if err := viper.Sub("redis." + meta.RedisSingleDriverKey).Unmarshal(&cfg); err != nil {
		return nil, nil, err
	}
	return redisClient.NewClient(cfg)
}

// NewRedisGroupClient 连接RedisCluster集群
func NewRedisGroupClient() (meta.RedisGroupClient, func(), error) {
	var cfg redisClient.GroupConfig
	if err := viper.Sub("redis." + meta.RedisGroupDriverKey).Unmarshal(&cfg); err != nil {
		return nil, nil, err
	}
	return redisClient.NewGroupClient(cfg)
}

// NewMongoClient 连接mongodb单机
func NewMongoClient() (meta.MongoClient, func(), error) {
	var cfg mongoClient.MongoConfig
	if err := viper.Sub("mongodb").Unmarshal(&cfg); err != nil {
		return nil, nil, err
	}
	return mongoClient.NewMongoClient(cfg)
}

// NewConsulClient 连接consul单机
func NewConsulClient() (meta.ConsulClient, error) {
	addr := viper.GetString("consul.addr")
	return consulClient.NewClient(addr)
}

func NewJenkinsClient() (meta.JenkinsClient, error) {
	var cfg jenkinsClient.JenkinsConfig
	if err := viper.Sub("jenkins").Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return jenkinsClient.NewJenkinsClient(cfg)
}

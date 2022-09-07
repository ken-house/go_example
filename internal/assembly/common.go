package assembly

import (
	"github.com/go_example/internal/meta"
	"github.com/ken-house/go-contrib/prototype/alibabaSmsClient"
	"github.com/ken-house/go-contrib/prototype/consulClient"
	"github.com/ken-house/go-contrib/prototype/jenkinsClient"
	"github.com/ken-house/go-contrib/prototype/mongoClient"
	"github.com/ken-house/go-contrib/prototype/mysqlClient"
	"github.com/ken-house/go-contrib/prototype/redisClient"
	"github.com/ken-house/go-contrib/utils/env"
)

// NewMysqlSingleClient 单机数据库连接
func NewMysqlSingleClient() (meta.MysqlSingleClient, func(), error) {
	cfg := mysqlClient.SingleConfig{
		MaxIdle:     meta.GlobalConfig.Mysql.Single.MaxIdle,
		MaxOpen:     meta.GlobalConfig.Mysql.Single.MaxOpen,
		MaxLifetime: meta.GlobalConfig.Mysql.Single.MaxLifetime,
		Dsn:         meta.GlobalConfig.Mysql.Single.Dsn,
		IsDebug:     !env.IsReleasing(),
	}
	return mysqlClient.NewSingleClient(cfg)
}

// NewMysqlGroupClient 主从数据库连接
func NewMysqlGroupClient() (meta.MysqlGroupClient, func(), error) {
	cfg := mysqlClient.GroupConfig{
		MaxIdle:     meta.GlobalConfig.Mysql.Group.MaxIdle,
		MaxOpen:     meta.GlobalConfig.Mysql.Group.MaxOpen,
		MaxLifetime: meta.GlobalConfig.Mysql.Group.MaxLifetime,
		IsDebug:     !env.IsReleasing(),
		Slaves:      nil,
	}
	cfg.Master.Dsn = meta.GlobalConfig.Mysql.Group.Master.Dsn
	slaveList := make([]struct {
		Dsn string `json:"dsn" mapstructure:"dsn"`
	}, 0, 10)
	var dsn struct {
		Dsn string `json:"dsn" mapstructure:"dsn"`
	}
	for _, v := range meta.GlobalConfig.Mysql.Group.Slaves {
		dsn.Dsn = v.Dsn
		slaveList = append(slaveList, dsn)
	}
	cfg.Slaves = slaveList
	return mysqlClient.NewGroupClient(cfg)
}

// NewRedisSingleClient 连接Redis单机
func NewRedisSingleClient() (meta.RedisSingleClient, func(), error) {
	cfg := redisClient.RedisConfig{
		Addr:     meta.GlobalConfig.Redis.Single.Addr,
		Password: meta.GlobalConfig.Redis.Single.Password,
		DB:       meta.GlobalConfig.Redis.Single.DB,
		PoolSize: meta.GlobalConfig.Redis.Single.PoolSize,
	}
	return redisClient.NewClient(cfg)
}

// NewRedisGroupClient 连接RedisCluster集群
func NewRedisGroupClient() (meta.RedisGroupClient, func(), error) {
	cfg := redisClient.GroupConfig{
		Addrs:    meta.GlobalConfig.Redis.Group.Addrs,
		Password: meta.GlobalConfig.Redis.Group.Password,
		PoolSize: meta.GlobalConfig.Redis.Group.PoolSize,
	}
	return redisClient.NewGroupClient(cfg)
}

// NewMongoClient 连接mongodb单机
func NewMongoClient() (meta.MongoClient, func(), error) {
	cfg := mongoClient.MongoConfig{
		Addr:     meta.GlobalConfig.Mongodb.Addr,
		Username: meta.GlobalConfig.Mongodb.Username,
		Password: meta.GlobalConfig.Mongodb.Password,
		MaxOpen:  meta.GlobalConfig.Mongodb.MaxOpen,
	}
	return mongoClient.NewMongoClient(cfg)
}

// NewConsulClient 连接consul单机
func NewConsulClient() (meta.ConsulClient, error) {
	addr := meta.GlobalConfig.Consul.Addr
	return consulClient.NewClient(addr)
}

func NewJenkinsClient() (meta.JenkinsClient, error) {
	cfg := jenkinsClient.JenkinsConfig{
		Host:     meta.GlobalConfig.Jenkins.Host,
		Username: meta.GlobalConfig.Jenkins.Username,
		Password: meta.GlobalConfig.Jenkins.Password,
	}
	return jenkinsClient.NewJenkinsClient(cfg)
}

// NewAlibabaSmsClient alibaba短信连接
func NewAlibabaSmsClient() (meta.AlibabaSmsClient, error) {
	cfg := alibabaSmsClient.ClientConfig{
		EndPoint:        meta.GlobalConfig.AlibabaSms.EndPoint,
		AccessKeyId:     meta.GlobalConfig.AlibabaSms.AccessKeyId,
		AccessKeySecret: meta.GlobalConfig.AlibabaSms.AccessKeySecret,
	}
	return alibabaSmsClient.CreateClient(cfg)
}

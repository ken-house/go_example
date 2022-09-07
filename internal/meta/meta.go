package meta

import (
	"github.com/go_example/internal/model"
	"github.com/ken-house/go-contrib/prototype/alibabaSmsClient"
	"github.com/ken-house/go-contrib/prototype/jenkinsClient"
	"time"

	"github.com/ken-house/go-contrib/prototype/mongoClient"

	"github.com/ken-house/go-contrib/prototype/consulClient"
	"github.com/ken-house/go-contrib/prototype/mysqlClient"
	"github.com/ken-house/go-contrib/prototype/redisClient"
	"github.com/patrickmn/go-cache"
)

// EnvMode 定义运行环境
var EnvMode string

// CfgFile 配置文件所在路径
const CfgFile = "./configs"

// DebugUseConfigCenter 本地调试是否使用配置中心
const DebugUseConfigCenter = false

// GlobalConfig 全局配置
var GlobalConfig model.GlobalConfig

const (
	MysqlSingleDriverKey = "single" // 单机数据库
	MysqlGroupDriverKey  = "group"  // MysqlGroupDriverKey Mysql配置驱动Key
	RedisSingleDriverKey = "single" // redis普通连接
	RedisGroupDriverKey  = "group"  // redisCluster连接
)

const HEALTHCHECK_SERVICE = "grpc.health.v1.Health"

// SocketWhiteIpList IP白名单
var SocketWhiteIpList = []string{"127.0.0.1", "192.168.163.*"}

// CacheDriver go-cache缓存对象
var CacheDriver = cache.New(5*time.Minute, 10*time.Minute)

// MysqlGroupClient 主从数据库连接
type MysqlGroupClient mysqlClient.GroupClient

// MysqlSingleClient 单机数据库连接
type MysqlSingleClient mysqlClient.SingleClient

// RedisSingleClient redis普通连接
type RedisSingleClient redisClient.RedisClient

// RedisGroupClient redisCluster连接
type RedisGroupClient redisClient.GroupClient

// MongoClient mongodb单机连接
type MongoClient mongoClient.MongoClient

// ConsulClient consul连接
type ConsulClient consulClient.ConsulClient

// JenkinsClient jenkins连接
type JenkinsClient jenkinsClient.JenkinsClient

// AlibabaSmsClient alibaba短信连接
type AlibabaSmsClient alibabaSmsClient.AlibabaSmsClient

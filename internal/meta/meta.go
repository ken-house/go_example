package meta

import (
	"time"

	"github.com/go_example/common/consulClient"
	"github.com/go_example/common/mysqlClient"
	"github.com/go_example/common/redisClient"
	"github.com/patrickmn/go-cache"
)

// EnvMode 定义运行环境
var EnvMode string

// CfgFile 配置文件所在路径
const CfgFile = "./configs"

const (
	MysqlSingleDriverKey = "single"     // 单机数据库
	MysqlGroupDriverKey  = "go_example" // MysqlGroupDriverKey Mysql配置驱动Key
	RedisSingleDriverKey = "single"     // redis普通连接
	RedisGroupDriverKey  = "group"      // redisCluster连接
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
type RedisSingleClient redisClient.SingleClient

// RedisGroupClient redisCluster连接
type RedisGroupClient redisClient.GroupClient

// ConsulClient consul连接
type ConsulClient consulClient.ConsulClient

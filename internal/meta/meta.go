package meta

import (
	"github.com/go_example/internal/model"
	"github.com/ken-house/go-contrib/prototype/alibabaSmsClient"
	"github.com/ken-house/go-contrib/prototype/jenkinsClient"
	"github.com/ken-house/go-contrib/prototype/kafkaClient"
	"github.com/ken-house/go-contrib/prototype/nacosClient"
	"github.com/ken-house/go-contrib/prototype/openTelemetry"
	"go.opentelemetry.io/otel/trace"
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

// NacosConfig Nacos配置
var NacosConfig nacosClient.Config

// GlobalConfig 全局配置
var GlobalConfig model.GlobalConfig

const (
	MysqlSingleDriverKey = "single" // 单机数据库
	MysqlGroupDriverKey  = "group"  // MysqlGroupDriverKey Mysql配置驱动Key
	RedisSingleDriverKey = "single" // redis普通连接
	RedisGroupDriverKey  = "group"  // redisCluster连接
)

const HEALTHCHECK_SERVICE = "grpc.health.v1.Health"

// CacheDriver go-cache缓存对象
var CacheDriver = cache.New(5*time.Minute, 10*time.Minute)

// HttpTracer http服务分布式追踪对象
var HttpTracer trace.Tracer

// GrpcTracer grpc服务分布式追踪对象
var GrpcTracer trace.Tracer

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

// NacosServiceClient nacos服务注册与发现
type NacosServiceClient nacosClient.ServiceClient

// KafkaProducerSyncClient kafka同步生产者
type KafkaProducerSyncClient kafkaClient.ProducerSyncClient

// KafkaProducerAsyncClient kafka同步生产者
type KafkaProducerAsyncClient kafkaClient.ProducerAsyncClient

// KafkaConsumerClient kafka消费者
type KafkaConsumerClient kafkaClient.ConsumerClient

// KafkaConsumerGroupClient kafka消费者组
type KafkaConsumerGroupClient kafkaClient.ConsumerGroupClient

// TracerProvider 分布式追踪提供者
type TracerProvider openTelemetry.TracerProvider

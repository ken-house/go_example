package assembly

import (
	"github.com/go_example/internal/meta"
	"github.com/ken-house/go-contrib/prototype/alibabaSmsClient"
	"github.com/ken-house/go-contrib/prototype/consulClient"
	"github.com/ken-house/go-contrib/prototype/jenkinsClient"
	"github.com/ken-house/go-contrib/prototype/kafkaClient"
	"github.com/ken-house/go-contrib/prototype/mongoClient"
	"github.com/ken-house/go-contrib/prototype/mysqlClient"
	"github.com/ken-house/go-contrib/prototype/nacosClient"
	"github.com/ken-house/go-contrib/prototype/openTelemetry"
	"github.com/ken-house/go-contrib/prototype/redisClient"
	"github.com/ken-house/go-contrib/utils/env"
)

// NewMysqlSingleClient 单机数据库连接
func NewMysqlSingleClient() (meta.MysqlSingleClient, func(), error) {
	cfg := meta.GlobalConfig.Mysql.Single
	cfg.IsDebug = !env.IsReleasing()
	return mysqlClient.NewSingleClient(meta.GlobalConfig.Mysql.Single)
}

// NewMysqlGroupClient 主从数据库连接
func NewMysqlGroupClient() (meta.MysqlGroupClient, func(), error) {
	cfg := meta.GlobalConfig.Mysql.Group
	cfg.IsDebug = !env.IsReleasing()
	return mysqlClient.NewGroupClient(cfg)
}

// NewRedisSingleClient 连接Redis单机
func NewRedisSingleClient() (meta.RedisSingleClient, func(), error) {
	return redisClient.NewClient(meta.GlobalConfig.Redis.Single)
}

// NewRedisGroupClient 连接RedisCluster集群
func NewRedisGroupClient() (meta.RedisGroupClient, func(), error) {
	return redisClient.NewGroupClient(meta.GlobalConfig.Redis.Group)
}

// NewMongoClient 连接mongodb单机
func NewMongoClient() (meta.MongoClient, func(), error) {
	return mongoClient.NewMongoClient(meta.GlobalConfig.Mongodb)
}

// NewConsulClient 连接consul单机
func NewConsulClient() (meta.ConsulClient, error) {
	return consulClient.NewClient(meta.GlobalConfig.Consul)
}

// NewJenkinsClient 连接jenkins
func NewJenkinsClient() (meta.JenkinsClient, error) {
	return jenkinsClient.NewJenkinsClient(meta.GlobalConfig.Jenkins)
}

// NewAlibabaSmsClient alibaba短信连接
func NewAlibabaSmsClient() (meta.AlibabaSmsClient, error) {
	return alibabaSmsClient.CreateClient(meta.GlobalConfig.AlibabaSms)
}

// NewNacosServiceClient nacos服务注册与服务发现
func NewNacosServiceClient() (meta.NacosServiceClient, func(), error) {
	return nacosClient.NewServiceClient(meta.NacosConfig)
}

// NewProducerSyncClient kafka同步生产者
func NewProducerSyncClient() (meta.KafkaProducerSyncClient, func(), error) {
	return kafkaClient.NewProducerSyncClient(meta.GlobalConfig.Kafka)
}

// NewProducerAsyncClient kafka异步生产者
func NewProducerAsyncClient() (meta.KafkaProducerAsyncClient, func(), error) {
	return kafkaClient.NewProducerAsyncClient(meta.GlobalConfig.Kafka)
}

// NewConsumerClient kafka消费者
func NewConsumerClient() (meta.KafkaConsumerClient, func(), error) {
	return kafkaClient.NewConsumerClient(meta.GlobalConfig.Kafka)
}

// NewConsumerGroupClient kafka消费者组
func NewConsumerGroupClient() (meta.KafkaConsumerGroupClient, func(), error) {
	return kafkaClient.NewConsumerGroupClient(meta.GlobalConfig.Kafka)
}

// NewTracerProvider 分布式追踪提供者
func NewTracerProvider() (meta.TracerProvider, func(), error) {
	return openTelemetry.NewTracerProvider(meta.GlobalConfig.TracerConfig)
}

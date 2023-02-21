package model

import (
	"github.com/ken-house/go-contrib/prototype/alibabaSmsClient"
	"github.com/ken-house/go-contrib/prototype/consulClient"
	"github.com/ken-house/go-contrib/prototype/emailClient"
	"github.com/ken-house/go-contrib/prototype/jenkinsClient"
	"github.com/ken-house/go-contrib/prototype/kafkaClient"
	"github.com/ken-house/go-contrib/prototype/mongoClient"
	"github.com/ken-house/go-contrib/prototype/mysqlClient"
	"github.com/ken-house/go-contrib/prototype/rabbitmqClient"
	"github.com/ken-house/go-contrib/prototype/redisClient"
	"github.com/ken-house/go-contrib/prototype/sentryClient"
)

// GlobalConfig 项目全局配置
type GlobalConfig struct {
	Server struct {
		Http      ServerInfo `json:"http"`
		HttpPprof ServerInfo `json:"http_pprof"`
		Socket    ServerInfo `json:"socket"`
		Grpc      ServerInfo `json:"grpc"`
	} `json:"server" mapstructure:"server"`

	Certs CertInfo `json:"certs" mapstructure:"certs"`

	Common Common `json:"common" mapstructure:"common"`

	Mysql struct {
		Group  mysqlClient.GroupConfig  `json:"group" mapstructure:"group"`
		Single mysqlClient.SingleConfig `json:"single" mapstructure:"single"`
	} `json:"mysql" mapstructure:"mysql"`

	Redis struct {
		Group  redisClient.GroupConfig `json:"group" mapstructure:"group"`
		Single redisClient.RedisConfig `json:"single" mapstructure:"single"`
	} `json:"redis" mapstructure:"redis"`

	Mongodb mongoClient.MongoConfig `json:"mongodb" mapstructure:"mongodb"`

	Consul consulClient.ConsulConfig `json:"consul" mapstructure:"consul"`

	Jenkins jenkinsClient.JenkinsConfig `json:"jenkins" mapstructure:"jenkins"`

	AlibabaSms alibabaSmsClient.ClientConfig `json:"alibaba_sms" mapstructure:"alibaba_sms"`

	AlibabaSmsCode AlibabaSmsCode `json:"alibaba_sms_code" mapstructure:"alibaba_sms_code"`

	Kafka kafkaClient.Config `json:"kafka" mapstructure:"kafka"`

	Email emailClient.EmailConf `json:"email" mapstructure:"email"`

	Sentry sentryClient.SentryConfig `json:"sentry" mapstructure:"sentry"`

	Rabbitmq rabbitmqClient.Config `json:"rabbitmq" mapstructure:"rabbitmq"`
}

// ServerInfo 服务信息
type ServerInfo struct {
	Name string `json:"name" mapstructure:"name"`
	Addr string `json:"addr" mapstructure:"addr"`
	Port string `json:"port" mapstructure:"port"`
}

// CertInfo 证书信息
type CertInfo struct {
	CurKey string   `json:"cur_key" mapstructure:"cur_key"`
	Keys   []string `json:"keys" mapstructure:"keys"`
}

// Common 项目配置信息
type Common struct {
	SocketWhiteIpList []string `json:"socket_white_ip_list" mapstructure:"socket_white_ip_list"`
	AesKey            string   `json:"aes_key" mapstructure:"aes_key"`
}

// AlibabaSmsCode 阿里巴巴短信格式
type AlibabaSmsCode struct {
	SignName     string `json:"sign_name" mapstructure:"sign_name"`
	TemplateCode string `json:"template_code" mapstructure:"template_code"`
}

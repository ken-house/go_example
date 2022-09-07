package model

// GlobalConfig 项目全局配置
type GlobalConfig struct {
	Server struct {
		Http      ServerInfo `json:"http"`
		HttpPprof ServerInfo `json:"http_pprof"`
		Socket    ServerInfo `json:"socket"`
		Grpc      ServerInfo `json:"grpc"`
	} `json:"server" mapstructure:"server"`

	Certs CertInfo `json:"certs" mapstructure:"certs"`

	Mysql struct {
		Group  MysqlGroup  `json:"group" mapstructure:"group"`
		Single MysqlSingle `json:"single" mapstructure:"single"`
	} `json:"mysql" mapstructure:"mysql"`

	Redis struct {
		Group  RedisGroup  `json:"group" mapstructure:"group"`
		Single RedisSingle `json:"single" mapstructure:"single"`
	} `json:"redis" mapstructure:"redis"`

	Mongodb Mongo `json:"mongodb" mapstructure:"mongodb"`

	Consul Consul `json:"consul" mapstructure:"consul"`

	Jenkins Jenkins `json:"jenkins" mapstructure:"jenkins"`

	AlibabaSms AlibabaSms `json:"alibaba_sms" mapstructure:"alibaba_sms"`

	AlibabaSmsCode AlibabaSmsCode `json:"alibaba_sms_code" mapstructure:"alibaba_sms_code"`
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

// MysqlGroup Mysql主从配置
type MysqlGroup struct {
	MaxIdle     int `json:"max_idle" mapstructure:"max_idle"`
	MaxOpen     int `json:"max_open" mapstructure:"max_open"`
	MaxLifetime int `json:"max_lifetime" mapstructure:"max_lifetime"`
	Master      struct {
		Dsn string `json:"dsn" mapstructure:"dsn"`
	} `json:"master" mapstructure:"master"`
	Slaves []struct {
		Dsn string `json:"dsn" mapstructure:"dsn"`
	} `json:"slaves" mapstructure:"slaves"`
}

// MysqlSingle Mysql单机
type MysqlSingle struct {
	MaxIdle     int    `json:"max_idle" mapstructure:"max_idle"`
	MaxOpen     int    `json:"max_open" mapstructure:"max_open"`
	MaxLifetime int    `json:"max_lifetime" mapstructure:"max_lifetime"`
	Dsn         string `json:"dsn" mapstructure:"dsn"`
	IsDebug     bool   `json:"is_debug" mapstructure:"is_debug"`
}

// RedisSingle Redis单机
type RedisSingle struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	Password string `json:"password" mapstructure:"password"`
	DB       int    `json:"db" mapstructure:"db"`
	PoolSize int    `json:"pool_size" mapstructure:"pool_size"`
}

// RedisGroup Redis集群
type RedisGroup struct {
	Addrs    []string `json:"addrs" mapstructure:"addrs"`
	Password string   `json:"password" mapstructure:"password"`
	PoolSize int      `json:"pool_size" mapstructure:"pool_size"`
}

// Mongo mongo配置
type Mongo struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	MaxOpen  uint64 `json:"max_open" mapstructure:"max_open"`
}

// Consul consul服务
type Consul struct {
	Addr string `json:"addr" mapstructure:"addr"`
}

// Jenkins jenkins服务
type Jenkins struct {
	Host     string `json:"host" mapstructure:"host"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

// AlibabaSms 阿里巴巴短信服务
type AlibabaSms struct {
	EndPoint        string `json:"end_point" mapstructure:"end_point"`
	AccessKeyId     string `json:"access_key_id" mapstructure:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" mapstructure:"access_key_secret"`
}

// AlibabaSmsCode 阿里巴巴短信格式
type AlibabaSmsCode struct {
	SignName     string `json:"sign_name" mapstructure:"sign_name"`
	TemplateCode string `json:"template_code" mapstructure:"template_code"`
}

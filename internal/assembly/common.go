package assembly

import (
	"github.com/go_example/common/mysqlClient"
	"github.com/go_example/internal/meta"
	"github.com/go_example/internal/utils/env"
	"github.com/spf13/viper"
)

// NewMysqlSingleClient 单机数据库连接
func NewMysqlSingleClient() (meta.MysqlSingleClient, func(), error) {
	var cfg mysqlClient.SingleConfig
	if err := viper.Sub("mysql." + meta.MysqlSingleDriverKey).Unmarshal(&cfg); err != nil {
		return nil, nil, err
	}
	return mysqlClient.NewSimpleClient(cfg)
}

// NewMysqlGroupClient 主从数据库连接
func NewMysqlGroupClient() (meta.MysqlGroupClient, func(), error) {
	var cfg mysqlClient.GroupConfig
	if err := viper.Sub("mysql." + meta.MysqlDriverKey).Unmarshal(&cfg); err != nil {
		return nil, nil, err
	}
	cfg.IsDebug = !env.IsReleasing()
	return mysqlClient.NewGroupClient(cfg)
}

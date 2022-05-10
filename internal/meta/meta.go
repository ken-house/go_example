package meta

import "github.com/go_example/common/mysqlClient"

// EnvMode 定义运行环境
var EnvMode string

// CfgFile 配置文件所在路径
const CfgFile = "./configs"

const (
	MysqlSingleDriverKey = "single"     // 单机数据库
	MysqlDriverKey       = "go_example" // MysqlDriverKey Mysql配置驱动Key
)

// MysqlGroupClient 主从数据库连接
type MysqlGroupClient mysqlClient.GroupClient

// MysqlSingleClient 单机数据库连接
type MysqlSingleClient mysqlClient.SingleClient

# go_example

## 简介

本项目为基础的 web 开发架构设计，主要采用 gin 框架，使用 cobra 生成应用和命令文件的脚手架，使用 wire 解决依赖注入问题，
最终实现一个高性能、可扩展、多应用的 web 框架。 除 HTTP 服务外，还包含 Socket 服务、GRPC 服务、Consul/Nacos 服务注册与服务发现、配置中心、Kafka 等功能。还支持 zap 日志、pprof 性能分析、sentry 异常监控。
支持 Docker 容器化部署、K8s 部署；

## 功能

- 连接 MySQL 单机及主从数据库；
- 连接 Redis 单机及 Cluster 集群；
- 使用 JWT 登录验证及单点登录；
- Excel 文件（.xlsx）导入导出；
- 提供 WebSocket 服务；
- 提供 gRPC 服务；
- 支持 Consul 服务注册与服务发现、配置中心；
- 增加 go-cache 内存缓存使用；
- 连接 MongoDB 单机及集群；
- 增加 PProf 性能分析；
- 增加 gRPC 熔断降级处理；
- 增加 zap 高性能日志库；
- 增加 gin 参数验证；
- 增加 jenkins 服务；
- 增加阿里云短信服务；
- 支持自动生成 CRUD 代码；
- 支持 nacos 服务注册与发现、配置中心；
- 支持 kafka 服务；
- 支持 Docker 容器化部署；
- 支持 K8s 容器编排部署；
- 支持单元测试；
- 支持邮件服务；
- 提供 ip 查询地区服务；
- 增加 jsoniter 优化 json 格式化；
- 支持 sentry 异常监控；

## 主要贡献

- https://github.com/gin-gonic/gin
- https://github.com/spf13/cobra
- https://github.com/google/wire
- https://github.com/spf13/viper
- https://xorm.io/xorm
- https://github.com/go-redis/redis
- https://github.com/golang-jwt/jwt
- https://github.com/qax-os/excelize
- https://github.com/gorilla/websocket
- https://google.golang.org/grpc
- https://github.com/hashicorp/consul
- https://github.com/patrickmn/go-cache
- https://go.mongodb.org/mongo-driver
- https://github.com/felixge/fgprof
- https://github.com/afex/hystrix-go
- https://github.com/uber-go/zap
- https://github.com/go-playground/validator
- https://github.com/dlclark/regexp2
- https://github.com/bndr/gojenkins
- https://github.com/alibabacloud-go
- https://github.com/nacos-group/nacos-sdk-go
- https://github.com/Shopify/sarama
- https://github.com/panjf2000/ants
- https://github.com/go-gomail/gomail/tree/v2
- https://github.com/lionsoul2014/ip2region
- https://github.com/json-iterator/go
- https://github.com/getsentry/sentry-go

## 版本

- 版本 v1.0.0 实现了 cobra+gin 框架的结合；
- 版本 v1.1.0 增加了 wire 解决依赖注入及项目文件目录整体架构；
- 版本 v1.2.0 增加了 xorm 连接 Mysql 数据库；
- 版本 v1.3.1 实现了 redis 连接；
- 版本 v1.4.0 接入 JWT 实现登录验证；
- 版本 v1.4.1 增加 JWT Refresh Token；
- 版本 v1.4.2 升级 JWT 密钥；
- 版本 v1.4.3 单点登录；
- 版本 v1.4.4 单点登录，对代码依赖优化；
- 版本 v1.5.0 增加 xlsx 文件导入导出；
- 版本 v1.6.0 实现 Gin 优雅关机；
- 版本 v1.7.0 实现 WebSocket 服务；
- 版本 v1.7.1 实现 WebSocket 心跳检测及客户端来源检查；
- 版本 v1.8.0 实现 gRPC 服务；
- 版本 v1.8.1 增加 gRPC 使用服务端认证及使用客户端/服务端各自证书认证；
- 版本 v1.8.2 实现 gRPC 基于 Token 认证；
- 版本 v1.9.0 使用 consul 做服务注册与服务发现；
- 版本 v1.10.0 增加 go-cache 内存缓存；
- 版本 v1.11.0 实现 mongodb 连接；
- 版本 v1.11.1 修改为 mongodb 可连接单机或集群；
- 版本 v1.12.0 增加了 pprof 性能分析器
- 版本 v1.13.0 增加 gRPC 熔断降级处理；
- 版本 v1.14.0 统一错误码；
- 版本 v2.0.0 调整代码结构引用自有包；
- 版本 v2.1.0 增加 zap 高性能日志；
- 版本 v2.2.0 增加 gin 参数验证；
- 版本 v2.3.0 增加 jenkins 服务；
- 版本 v2.4.0 增加阿里云短信服务；
- 版本 v2.5.0 支持自动生成 CRUD 代码；
- 版本 v2.6.0 支持 Nacos 配置中心；
- 版本 v2.6.1 支持 Nacos 配置中心监听，客户端自动感知；
- 版本 v2.6.2 增加 consul 配置中心，nacos 服务注册与服务发现；
- 版本 v2.7.0 增加 kafka 生产消费服务；
- 版本 v3.0.0 支持 Docker 容器化部署；
- 版本 v3.1.0 支持 K8s 容器编排部署；
- 版本 v3.1.1 实现 K8s 容器编排部署全面升级；
- 版本 v3.2.0 增加单元测试示例；
- 版本 v3.3.0 对自动生成 CRUD 代码升级，实现完全自动化；
- 版本 v3.4.0 增加邮件服务、ip 查询地区、优化 json 格式化；
- 版本 v3.4.1 支持 sentry 异常监控；
- 版本 v3.4.2 支持 AES 对称加密；

## 环境安装

可以使用 Linux 安装，也可以通过 Docker 安装相关服务，以下使用 Docker 安装服务：

1. MySQL 服务；
   这里推荐使用 MySQL8.0.18 版本，高版本数据库在 xorm 生成 model 时会报错

```shell
# 拉取镜像后，启动容器
docker pull mysql:8.0.18
docker run -itd -p 3306:3306 -v /data/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root --name mysql mysql:8.0.18

# 进入容器，设置可远程连接
docker exec -it mysql /bin/bash
# 连接mysql
root@af4b3222a631: mysql -u root -p
# 设置本地连接和远程连接
mysql> use mysql;
# 可以看到root有一个本地连接和一个远程连接，plugin都是caching_sha2_password
mysql> select user,host,plugin from user;
+------------------+-----------+-----------------------+
| user             | host      | plugin                |
+------------------+-----------+-----------------------+
| root             | %         | caching_sha2_password |
| mysql.infoschema | localhost | caching_sha2_password |
| mysql.session    | localhost | caching_sha2_password |
| mysql.sys        | localhost | caching_sha2_password |
| root             | localhost | caching_sha2_password |
+------------------+-----------+-----------------------+

# 修改本地连接
mysql> alter user 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root';
# 修改远程连接
mysql> alter user 'root'@'%'IDENTIFIED WITH mysql_native_password BY 'root';

# 现在可以使用navicat工具连接Mysql，创建项目所需的数据库go_example及user表
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `gender` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '性别 1男 2女 0未知',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';
```

2. Redis 服务；

```shell
# 拉去镜像，运行容器
docker pull redis:latest
docker run -itd -p 6379:6379 -v /data/redis:/data --name redis redis:latest
```

3. Mongodb 服务；

```shell
docker pull mongo:latest
docker run -itd -p 27017:27017 -v /data/mongo:/data/db --name mongo mongo:latest
```

4. Consul 服务；

```shell
docker pull consul:latest
docker run -itd -p 8500:8500 -v /data/consul/data:/consul/data -v /data/consul/conf:/consul/conf --name consul consul:latest
```

5. Jenkins 服务；

```shell
# 拉取jenkins/jenkins:latest镜像
docker pull jenkins/jenkins
docker run -itd -p 8088:8080 -p 50000:50000 \
           -u root -e TZ="Asia/shanghai" \
           -v /data/jenkins:/var/jenkins_home \
           -v /var/run/docker.sock:/var/run/docker.sock \
           --name jenkins jenkins/jenkins
```

6. Nacos 服务；

```shell
docker image pull nacos/nacos-server
docker run -itd
      -p 8848:8848 \ # 用于nacos管理界面；
      -p 9848:9848 \ # Nacos2.0版本增加gRPC通信方式，用于客户端gRPC向服务端发起连接和请求；
      -p 9849:9849 \ # Nacos2.0用于服务间通信同步数据，若集群模式必须要开启；
      -e MODE=standalone \
      -v /data/nacos/logs:/home/nacos/logs \
      -v /data/nacos/data:/home/nacos/data \
      --name nacos \
      nacos/nacos-server
```

访问http://127.0.0.1:8848/nacos/index.html进入管理界面，需做如下操作：

1. 命名空间创建 debug、test、release；
2. 配置管理 - 配置列表在相应的命名空间创建 Data-id 配置，以生产环境为例，其中 Group 为 go_example，Data Id 为 release-go_example-common.yaml；

如下图所示：
![image](assets/images/nacos.png) 7. Kafka 服务

```shell
# 创建自定义网络
docker network create --driver bridge --subnet 192.168.0.0/16 --gateway 192.168.0.1 mynet
# 运行zookeeper
docker pull wurstmeister/zookeeper

# 运行容器
docker run -d -p 2181:2181 --net mynet -v /etc/localtime:/etc/localtime -v ~/dockerVolumes/zookeeperVolume/data:/opt/zookeeper-3.4.13/data --name zookeeper wurstmeister/zookeeper

# 配置kafka集群，这里创建三个broker节点
docker run -d -p 9092:9092 --net mynet -e KAFKA_BROKER_ID=0 -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181/kafka -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka-0:9092 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 -v ~/dockerVolumes/kafkaVolume/kafka/0:/kafka -v /etc/localtime:/etc/localtime --name kafka-0 wurstmeister/kafka
docker run -d -p 9093:9093 --net mynet -e KAFKA_BROKER_ID=1 -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181/kafka -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka-1:9093 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9093 -v ~/dockerVolumes/kafkaVolume/kafka/1:/kafka -v /etc/localtime:/etc/localtime --name kafka-1 wurstmeister/kafka
docker run -d -p 9094:9094 --net mynet -e KAFKA_BROKER_ID=2 -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181/kafka -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka-2:9094 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9094 -v ~/dockerVolumes/kafkaVolume/kafka/2:/kafka -v /etc/localtime:/etc/localtime --name kafka-2 wurstmeister/kafka
```

除此之外，在使用还需要修改 hosts 文件

```shell
127.0.0.1 kafka-0 kafka-1 kafka-2
```

8. Kubernets 部署
   安装 Docker Desktop，通过 Docker Desktop 安装 kubernets 单点集群，操作步骤如下：
   ![image](assets/images/k8s.png)

安装完成后，通过如下命令查看集群情况

```shell
# 查看集群情况
kubectl cluster-info
# 查看集群节点列表
kubectl get nodes
```

## 使用

要求 golang 版本必须支持 Go Modules，本系统使用 1.18.3 版本。

克隆到本地目录

```shell
git clone git@github.com:ken-house/go_example.git
```

加载相应的依赖包

```shell
go mod tidy
```

启动服务

```go
go run main.go http
```

访问：http://127.0.0.1:8080/hello

## 目录结构

```
go_example/
├── cmd
│   ├── http.go // http服务入口
│   └── root.go // 根命令
├── configs // 配置文件
│   ├── debug
│   │   └── common.yaml
│   ├── dev
│   ├── prod
│   └── test
├── go.mod
├── go.sum
├── internal // 项目主体内容
│   ├── assembly // wire 定义依赖
│   │   ├── controller.go
│   │   ├── server.go
│   │   ├── service.go
│   │   └── wire_gen.go
│   ├── controller // 控制器文件
│   │   └── hello_controller.go
│   ├── meta // 定义常量、全局变量等
│   │   └── meta.go
│   ├── model // 模型类文件
│   ├── repository // 数据库等仓库文件
│   ├── server // 开启服务
│   │   └── http.go
│   ├── service // 程序服务类
│   │   └── hello_service.go
│   └── utils // 工具类
│       └── negotiate
│           └── negotiate.go
├── LICENSE
├── main.go
└── README.md
```

## 项目框架实现过程

### 新建项目

创建一个项目目录，例如：go_example，执行命令初始化 go modules。

```shell
go mod init github.com/go_example
```

### 创建脚手架

安装 cobra

```shell
cobra init ../go_example
```

创建 web 服务入口

```shell
cobra add http
```

### 引入 Gin 框架做 web 服务

在 cmd/http.go 文件中，创建 Gin 引擎，注册路由，运行服务。

```go
// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long:  `http server`,
	Run: func(cmd *cobra.Command, args []string) {
		// 实例化依赖注入服务
		httpSrv, clean, err := assembly.NewHttpServer()
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		defer clean()

		// 设置gin的运行环境
		gin.SetMode(meta.EnvMode)

		// 初始化engine
		app := gin.Default()

		// 注册路由
		httpSrv.Register(app)

		// 运行应用
		port := viper.GetString("server.http.addr")
		if err := app.Run(port); err != nil {
			log.Fatalf("%+v\n", err)
		}
	},
}
```

在 cmd/root.go 文件中，通过 cobra.OnInitialize(initConfig)对项目的配置进行初始化。

```go
func init() {
	// 初始化配置文件
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// 初始化配置文件
func initConfig() {
	// 从系统环境变量中读取运行环境
	meta.EnvMode = env.Mode()

	// viper目前仅支持单文件
	viper.SetConfigFile(meta.CfgFile + "/" + meta.EnvMode + "/common.yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
		log.Fatal(err)
	}
}
```

此时，运行项目就可以启动 web 服务，访问 127.0.0.1:8080/hello 即可访问。

```shell
go run main.go http
```

### wire 解决依赖注入

安装 wire

```shell
go get -u github.com/google/wire
```

对项目目录结构进行分层，目录结构如下：

![image](assets/images/1.png)

在项目中，web 服务调用 controller 控制器，控制器调用 service 服务类，服务类调用 repository 数据仓库层，
数据仓库层调用其他包生成的服务引擎客户端（如 mysql）。

这里以控制器调用服务类举例：
在 controller 目录下创建 hello_controller.go 文件。

1. 采用面向接口编程，因此会有一个向外提供服务的 interface 类型的 HelloController， 该接口规范了 struct 类型的接收者里的方法；
2. 在接收者 helloController 中定义要调用的服务 service.HelloService；
3. 创建一个实例化当前控制器的方法，提供给上层调用；

```go
package controller

import (
	"net/http"

	"github.com/go_example/internal/utils/negotiate"

	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/service"
)

type HelloController interface {
	Say(c *gin.Context) (int, gin.Negotiate)
}

type helloController struct {
	helloSvc service.HelloService
}

func NewHelloController(helloSvc service.HelloService) HelloController {
	return &helloController{
		helloSvc: helloSvc,
	}
}

func (ctr *helloController) Say(c *gin.Context) (int, gin.Negotiate) {
	data := ctr.helloSvc.SayHello(c)
	return negotiate.JSON(http.StatusOK, data)
}
```

在 service 目录下创建 hello_service.go 文件。同 hello_controller.go 文件类似定义。

```go
package service

import (
	"github.com/gin-gonic/gin"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	"github.com/spf13/viper"
)

type HelloService interface {
	SayHello(c *gin.Context) map[string]string
}

type helloService struct {
	UserRepo MysqlRepo.UserRepository
}

func NewHelloService(userRepo MysqlRepo.UserRepository) HelloService {
	return &helloService{
		UserRepo: userRepo,
	}
}

func (svc *helloService) SayHello(c *gin.Context) map[string]string {
	user, _ := svc.UserRepo.GetUserInfo(1)
	return map[string]string{
		"hello": "world，golang",
		"env":   viper.GetString("server.mode"),
		"user":  user.Name,
	}
}
```

assembly 目录下创建 wire 文件，这里命名为 controller.go，定义了控制器的依赖关系。
必须在文件第一行加上// +build wireinject 表示不参与编译。

```go
//go:build wireinject
// +build wireinject

package assembly

import (
	"github.com/go_example/internal/controller"
	"github.com/google/wire"
)

func NewHelloController() (controller.HelloController, func(), error) {
	panic(wire.Build(
		NewHelloService,
		controller.NewHelloController,
	))
}
```

同样创建一个服务类的 service.go 文件，定义服务类的依赖关系。

```go
//go:build wireinject

package assembly

import (
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	"github.com/go_example/internal/service"
	"github.com/google/wire"
)

func NewHelloService() (service.HelloService, func(), error) {
	panic(wire.Build(
		NewMysqlGroupClient,
		MysqlRepo.NewUserRepository,
		service.NewHelloService,
	))
}
```

进入 assembly 目录下，执行 wire 命令生成 wire_gen.go 文件。

```shell
cd ./internal/assembly
wire
```

这样就解决了文件互相依赖的问题，每层更加专注实现自己的功能，不用关心依赖方的实现。

## 连接 MySQL

### 安装

```shell
go get -u github.com/go-sql-driver/mysql
go get -u xorm.io/xorm
```

### 代码解析

在 assembly 目录下，创建一个 common.go 文件，该文件定义了获取 MySQL 连接，并遵循 wire 规范。

```go
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
```

以单机数据库连接为例，定义了 SingleClient 接口，singleClient 结构体为接收者，NewSingleClient 用来返回一个客户端连接及 wire 所需要的回调方法和错误。

在获取数据库连接中，使用了 xorm 创建数据库连接。

```go
package mysqlClient

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"xorm.io/xorm"
)

//SingleClient 单实例客户端
type SingleClient interface {
	xorm.EngineInterface
	Transaction(f func(*xorm.Session) (interface{}, error)) (interface{}, error)
	GetEngine() *xorm.Engine
}

type singleClient struct {
	*xorm.Engine
}

func NewSingleClient(cfg SingleConfig) (SingleClient, func(), error) {
	engine, err := newEngine(cfg)
	if err != nil {
		return nil, nil, err
	}
	client := &singleClient{Engine: engine}
	return client, func() {
		_ = client.Close()
	}, nil
}

func (cli *singleClient) GetEngine() *xorm.Engine {
	return cli.Engine
}

func (cli *singleClient) Transaction(f func(*xorm.Session) (interface{}, error)) (interface{}, error) {
	return cli.Engine.Transaction(f)
}

type SingleConfig struct {
	MaxIdle     int    `json:"max_idle" mapstructure:"max_idle"`
	MaxOpen     int    `json:"max_open" mapstructure:"max_open"`
	MaxLifetime int    `json:"max_lifetime" mapstructure:"max_lifetime"`
	Dsn         string `json:"dsn" mapstructure:"dsn"`
	IsDebug     bool   `json:"is_debug" mapstructure:"is_debug"`
}

func newEngine(cfg SingleConfig) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", cfg.Dsn)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err := engine.Ping(); err != nil {
		return nil, errors.WithStack(err)
	}
	if cfg.IsDebug {
		engine.ShowSQL(true)
	}
	if cfg.MaxIdle > 0 {
		engine.SetMaxIdleConns(cfg.MaxIdle)
	}
	if cfg.MaxOpen > 0 {
		engine.SetMaxOpenConns(cfg.MaxOpen)
	}
	if cfg.MaxLifetime > 0 {
		engine.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)
	}
	dsnCfg, _ := mysql.ParseDSN(cfg.Dsn)
	dsnCfg.Passwd = ""
	return engine, nil
}

```

最后，将 NewMysqlSingleClient 或 NewMysqlGroupClient 方法注入到 Repository 中，修改 assembly/service.go 文件如下：

```go
//go:build wireinject

package assembly

import (
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	RedisRepo "github.com/go_example/internal/repository/redis"
	"github.com/go_example/internal/service"
	"github.com/google/wire"
)

func NewHelloService() (service.HelloService, func(), error) {
	panic(wire.Build(
		NewRedisGroupClient,
		RedisRepo.NewUserRepository,
		NewMysqlGroupClient,
		MysqlRepo.NewUserRepository,
		service.NewHelloService,
	))
}
```

这样在数据仓库层（repository）就可以使用数据库连接了。

## 连接 Redis

### 安装

```shell
go get -u github.com/go-redis/redis/v8
```

### 代码解析

Redis 连接同 MySQL 类似，目的也是提供一个 Redis 连接，注入到 repository 中。
在 assembly/common.go 文件中定义两个方法分别为连接单机和 redisCluster 集群模式；

```go
// NewRedisSingleClient 连接Redis单机
func NewRedisSingleClient() (meta.RedisSingleClient, func(), error) {
	var cfg redisClient.SingleConfig
	if err := viper.Sub("redis." + meta.RedisSingleDriverKey).Unmarshal(&cfg); err != nil {
		return nil, nil, err
	}
	return redisClient.NewSingleClient(cfg)
}

// NewRedisGroupClient 连接RedisCluster集群
func NewRedisGroupClient() (meta.RedisGroupClient, func(), error) {
	var cfg redisClient.GroupConfig
	if err := viper.Sub("redis." + meta.RedisGroupDriverKey).Unmarshal(&cfg); err != nil {
		return nil, nil, err
	}
	return redisClient.NewGroupClient(cfg)
}
```

以单机连接为例，定义了 SingleClient 接口，singleClient 结构体为接收者，NewSingleClient 返回 redis 连接及回调方法和错误

获取 redis 连接使用了 github.com/go-redis/redis/v8 包，代码如下：

```go
package redisClient

import (
	"context"

	"github.com/pkg/errors"

	"github.com/go-redis/redis/v8"
)

type SingleClient interface {
	redis.UniversalClient
}

type singleClient struct {
	redis.UniversalClient
}

func NewSingleClient(cfg SingleConfig) (SingleClient, func(), error) {
	client, err := NewClient(cfg)
	if err != nil {
		return nil, nil, err
	}
	sc := &singleClient{UniversalClient: client}
	return sc, func() {
		client.Close()
	}, nil
}

type SingleConfig struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	Password string `json:"password" mapstructure:"password"`
	DB       int    `json:"db" mapstructure:"db"`
}

func NewClient(cfg SingleConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return client, err
}
```

最后，将 NewRedisSingleClient 和 NewRedisGroupClient 方法注入到 Repository 中，修改 assembly/service.go 文件如下：

```go
//go:build wireinject

package assembly

import (
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	RedisRepo "github.com/go_example/internal/repository/redis"
	"github.com/go_example/internal/service"
	"github.com/google/wire"
)

func NewHelloService() (service.HelloService, func(), error) {
	panic(wire.Build(
		NewRedisGroupClient,
		RedisRepo.NewUserRepository,
		NewMysqlGroupClient,
		MysqlRepo.NewUserRepository,
		service.NewHelloService,
	))
}
```

这样在数据仓库层（repository）就可以使用 Redis 连接了。

## JWT 登录验证

JWT 的出现就是为了解决传统 Session+Cookie 技术存在的各种问题，随着前后端分离的发展，JWT 被广泛使用。

### 普通验证

#### 生成 JWT

注册两个路由，分别为登录和首页，其中首页需要登录用户才可以访问，修改 internal/server/http.go 文件，代码如下：

```go
func (srv *httpServer) Register(router *gin.Engine) {
	router.GET("/hello", srv.Hello())
	// 登录接口
	router.POST("/login", srv.Login())
	// 用户主页
	router.GET("/home", middleware.JWTAuthMiddleware(), srv.Home())
}
```

当请求 127.0.0.1:8080/login 进行登录，系统根据提交的表单数据验证用户密码后生成 JWT，

```go
func (svc *loginService) Login(formData model.LoginForm) (token string, err error) {
	// 验证用户登录
	userInfo, err := svc.UserRepo.FindIdentity(formData)
	if err != nil {
		return "", errors.New("用户名密码不正确")
	}

	// 登录成功则生成Token
	return auth.GenToken(userInfo)
}
```

接下来，golang 如何生成 jwt 的，需要先安装：

```shell
go get -u github.com/golang-jwt/jwt/v4
```

- 新建 common/auth/auth.go 文件，因为我们要自定义参数，所以定义一个 CustomClaims 结构体，其中包含用户信息；
- 定义一个 token 的过期时间及加密盐；
- 定义生成 token 的方法，使用 jwt.SigningMethodHS256 加密方式生成签名对象，然后通过加密盐进行加密并返回 json 字符串;

```go
type CustomClaims struct {
	UserInfo MysqlModel.User `json:"user_info"`
	jwt.RegisteredClaims
}

const TokenExpireDuration = time.Hour * 24

// CustomSecret 用于加盐的字符串
var CustomSecret = []byte("custom_salt")

// GenToken 生成token
func GenToken(userInfo MysqlModel.User) (string, error) {
	claims := CustomClaims{
		userInfo,
		jwt.RegisteredClaims{
			Issuer:    "ken",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}
```

至此，登陆成功后会返回一个 json 字符串，我们称为 token(令牌)。

#### 解析 JWT

gin 框架支持中间价，在访问主页的路由上使用中间价，即可完成验证；

```go
router.GET("/home", middleware.JWTAuthMiddleware(), srv.Home())
```

创建 internal/middleware/jwt_auth_middlerware.go 文件，接收请求头里的 AUTHORIZATION 进行解析 JWT，解析成功将 JWT 中的 UserInfo 信息写入到 gin.context 中；

```go
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "request header authorization is empty",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"message": "request header authorization is invalid",
			})
			c.Abort()
			return
		}

		claims, err := auth.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "request header authorization parse token err",
			})
			c.Abort()
			return
		}

		c.Set("userInfo", claims.UserInfo)
		c.Next()
	}
}
```

接下来，就是对 token 进行解析，可以在 common/auth/auth.go 文件中看到，将 token 使用加密盐解析，最终得到 CustomClaims；

```go
// ParseToken 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return CustomSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token invalid")
}
```

至此，中间件若成功解析并写入到 gin 的上下文后，会进入到 srv.Home()的业务逻辑中。
这样就实现了普通的 JWT 登录验证。

### JWT 续期（Refresh Token）

JWT 是一次签发带有过期时间的令牌 Token，当登录成功后，系统会签发一个 token，过期时间假设为 30 分钟，若用户登录成功后在 30 分钟后操作，解析 JWT 将会失效，那么就无法进入主页。
为了解决这个问题，可以采用 refresh token 的方式。

在请求接口中，我们使用的 token 称为 access_token，该 token 带有用户信息；当 access_token 过期后，我们使用 refresh_token 重新获取新的 access_token,
具体流程如下：

- 客户端使用用户名密码进行认证
- 服务端生成有效时间较短的 Access Token（例如 10 分钟），和有效时间较长的 Refresh Token（例如 7 天）
- 客户端访问需要认证的接口时，携带 Access Token
- 如果 Access Token 没有过期，服务端鉴权后返回给客户端需要的数据
- 如果携带 Access Token 访问需要认证的接口时鉴权失败（例如返回 401 错误），则客户端使用 Refresh Token 向刷新接口申请新的 Access Token
- 如果 Refresh Token 没有过期，服务端向客户端下发新的 Access Token
- 客户端使用新的 Access Token 访问需要认证的接口

![image](assets/images/jwt.jpg)

对普通验证代码进行改造，生成 token 时，返回一个 access_token 和 refresh_token，解析 token 时使用对应的解密盐；

```go
package auth

import (
	"time"

	"github.com/pkg/errors"

	MysqlModel "github.com/go_example/internal/model/mysql"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	UserInfo MysqlModel.User `json:"user_info"`
	jwt.RegisteredClaims
}

const bearerSchema = "Bearer "
const accessTokenExpire = time.Minute * 30
const refreshTokenExpire = time.Hour * 24 * 30

// accessTokenSecret 用于加盐的字符串
var accessTokenSecret = []byte("custom_access_salt")

// refreshTokenSecret 用于加盐的字符串
var refreshTokenSecret = []byte("custom_refresh_salt")

// GenToken 生成token
func GenToken(userInfo MysqlModel.User) (string, string, error) {
	// 生成access_token
	accessClaims := CustomClaims{
		userInfo,
		jwt.RegisteredClaims{
			Issuer:    "ken",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExpire)),
		},
	}

	// 使用指定的签名方法创建签名对象
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	accessTokenSign, err := accessToken.SignedString(accessTokenSecret)
	if err != nil {
		return "", "", err
	}

	// 生成refresh_token，不包含自定义信息
	refreshClaims := CustomClaims{
		MysqlModel.User{
			Id: userInfo.Id,
		},
		jwt.RegisteredClaims{
			Issuer:    "ken",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenExpire)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenSign, err := refreshToken.SignedString(refreshTokenSecret)
	if err != nil {
		return "", "", err
	}

	return bearerSchema + accessTokenSign, bearerSchema + refreshTokenSign, nil
}

// ParseToken 解析token
func ParseToken(tokenString string, grantType string) (*CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		secret := accessTokenSecret
		if grantType == "refresh_token" {
			secret = refreshTokenSecret
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token invalid")
}
```

注册一个/auth/refresh_token 路由，用于刷新 token；

```go
func (srv *httpServer) Register(router *gin.Engine) {
	router.GET("/hello", srv.Hello())
	// 登录接口
	router.POST("/auth/login", srv.Login())
	// 刷新token接口
	router.GET("/auth/refresh_token", srv.RefreshToken())
	// 用户主页
	router.GET("/home", middleware.JWTAuthMiddleware(), srv.Home())
}
```

RefreshToken 方法，从头信息 AUTHORIZATION 中读取 refresh_token，若解析 refresh_token 成功，可以得到用户 id,查询用户信息，生成新的 access_token 及 refresh_token,客户端后期使用新的 access_token 进行请求；

```go
func (ctr *authController) RefreshToken(c *gin.Context) (int, gin.Negotiate) {
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		return negotiate.JSON(http.StatusOK, gin.H{"message": "令牌为空"})
	}

	parts := strings.SplitN(authorization, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return negotiate.JSON(http.StatusOK, gin.H{"message": "令牌格式错误"})
	}

	claims, err := auth.ParseToken(parts[1], "refresh_token")
	if err != nil {
		return negotiate.JSON(http.StatusOK, gin.H{"message": "刷新令牌失败，请重新登录"})
	}

	// 重新生成令牌
	userInfo, err := ctr.authSvc.GetUserInfo(claims.UserInfo.Id)
	if err != nil {
		return negotiate.JSON(http.StatusOK, gin.H{"message": "用户信息错误，请重新登录"})
	}
	accessToken, refreshToken, err := auth.GenToken(userInfo)
	if err != nil {
		return negotiate.JSON(http.StatusOK, gin.H{"message": "生成令牌失败，请重新登录"})
	}

	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
			"message":      "OK",
		},
	})
}
```

经过 Refresh Token 机制处理，实现了 JWT 的续期。

### 密钥对加密

前面对 token 加密都是通过自定义的密码盐进行加密，安全级别较低，使用密钥对加密 token，更加安全可靠！
创建 assets/certs/1652421664 目录，执行下面命令，生成密钥对：

```shell
# 生成私钥
openssl genrsa -out rsa_private_key.pem 1024
# 生成公钥
openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
```

在配置文件中定义密钥列表及当前使用的 key

```yaml
certs:
  curKey: 1652421664
  keys: [1652421664]
```

定义 SetCerts()方法读取证书内容到全局变量 certs 中，该方法将在 cmd/http.go 文件中的 RUN 方法调用。

```go
type cert struct {
	PrivateKey []byte
	PublicKey  []byte
}

// 证书密钥
var certs = make(map[string]cert)

// 当前使用的证书key
var curKey string

// SetCerts 生成密钥
func SetCerts() {
	var certKeyConfig struct {
		CurKey string
		Keys   []string
	}
	err := viper.Sub("certs").Unmarshal(&certKeyConfig)
	if err != nil {
		panic(err)
	}

	for _, key := range certKeyConfig.Keys {
		privateKey, err := ioutil.ReadFile("assets/certs/" + key + "/rsa_private_key.pem")
		if err != nil {
			panic(err)
		}
		publicKey, err := ioutil.ReadFile("assets/certs/" + key + "/rsa_public_key.pem")
		if err != nil {
			panic(err)
		}
		cert := cert{
			PrivateKey: privateKey,
			PublicKey:  publicKey,
		}
		certs[key] = cert
	}

	// 读取当前的使用的证书key
	curKey = viper.GetString("certs.curKey")
}
```

接下来只需要将之前使用的密码盐加密替换成从 pem 文件中读取密钥对进行加解密，注意这里 jwt.SigningMethodRS512 加密类型，必须和生成密钥对的类型一致。

```go
// GenToken 生成token
func GenToken(userInfo MysqlModel.User) (string, string, error) {
	// 读取当前使用的私钥证书
	secret, err := jwt.ParseRSAPrivateKeyFromPEM(certs[curKey].PrivateKey)
	if err != nil {
		panic(err)
	}
	// 生成access_token
	accessClaims := CustomClaims{
		userInfo,
		jwt.RegisteredClaims{
			Issuer:    "ken",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExpire)),
		},
	}

	// 使用指定的签名方法创建签名对象
	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS512, accessClaims)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	accessTokenSign, err := accessToken.SignedString(secret)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	// 生成refresh_token，不包含自定义信息
	refreshClaims := CustomClaims{
		MysqlModel.User{
			Id: userInfo.Id,
		},
		jwt.RegisteredClaims{
			Issuer:    "ken",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenExpire)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS512, refreshClaims)
	refreshTokenSign, err := refreshToken.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	return bearerSchema + accessTokenSign, bearerSchema + refreshTokenSign, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 读取当前使用的公钥证书
		secret, err := jwt.ParseRSAPublicKeyFromPEM(certs[curKey].PublicKey)
		if err != nil {
			panic(err)
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token invalid")
}
```

### 单点登录

由于 JWT 令牌是一次签发，在有效时间内一直有效，因此在多个设备登录后，本地的 token 是有效的，为了保证单点登录，可通过 Redis 记录用户最新一次的 token 信息。
生成或重新刷新 token 会将 redis hash 数据更新，解析 token 时，会检查 token 是否为用户的最新 token 信息。

在 jwt.go 文件中的 GenToken()方法中增加保存到 redis 的操作：

```go
// 将token写入redis hash中
err = saveAuthTokenRedis(userInfo.Id, accessTokenSign, refreshTokenSign)
if err != nil {
    return "", "", err
}
```

在 ParseToken()方法中获取 redis 的认证 token，并和当前使用的 token 判断；

```go
if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
    // 判断token是否有效，若无效返回错误
    err = checkAuthTokenRedis(claims.UserInfo.Id, tokenString, grantType)
    if err != nil {
        return nil, err
    }
    return claims, nil
}
```

### 单点登录优化（代码依赖）

在 v1.4.3 版本中，单独创建了 driver/目录，实现 redis 连接，这样就存在两套 Redis 连接（通过 wire 注入形式的连接和 driver 自己编写的连接），为了优化该问题，避免循环调用，
对代码进行了以下调整：

1. 在 httpServer 注入一个 authService，并在中间件中传入该服务类；

```go
type httpServer struct {
	helloCtr    controller.HelloController
	loginCtr    controller.AuthController
	homeCtr     controller.HomeController
	authService service.AuthService
}

func NewHttpServer(
	helloCtr controller.HelloController,
	loginCtr controller.AuthController,
	homeCtr controller.HomeController,
	authService service.AuthService,
) HttpServer {
	return &httpServer{
		helloCtr:    helloCtr,
		loginCtr:    loginCtr,
		homeCtr:     homeCtr,
		authService: authService,
	}
}

func (srv *httpServer) Register(router *gin.Engine) {
    router.GET("/hello", srv.Hello())
    // 登录接口
    router.POST("/auth/login", srv.Login())
    // 刷新token接口
    router.GET("/auth/refresh_token", srv.RefreshToken())
    // 用户主页
    router.GET("/home", middleware.JWTAuthMiddleware(srv.authService), srv.Home())
}
```

2. 中间件将 authService 服务类传入到 auth.ParseToken 方法中；

```go
func JWTAuthMiddleware(authService service.AuthService) func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "令牌为空",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"message": "令牌格式错误",
			})
			c.Abort()
			return
		}

		claims, err := auth.ParseToken(authService, parts[1], "access_token")
		if err != nil {
			if err.Error() == "账号已在其他设备登录" {
				c.JSON(http.StatusOK, gin.H{
					"message": "账号已在其他设备登录",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message":     "当前登录已失效，请尝试请求refresh_token获取新令牌",
					"refresh_url": "/auth/refresh_token",
				})
			}
			c.Abort()
			return
		}

		c.Set("userInfo", claims.UserInfo)
		c.Next()
	}
}
```

3. 在 auth/jwt.go 文件中，便可以使用 authService 下的验证令牌方法及生成令牌方法；

```go
// GenToken 生成token
func GenToken(authService service.AuthService, userInfo MysqlModel.User) (string, string, error) {
	// 读取当前使用的私钥证书
	secret, err := jwt.ParseRSAPrivateKeyFromPEM(certs[curKey].PrivateKey)
	if err != nil {
		panic(err)
	}
	// 生成access_token
	accessClaims := CustomClaims{
		userInfo,
		jwt.RegisteredClaims{
			Issuer:    "ken",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExpire)),
		},
	}

	// 使用指定的签名方法创建签名对象
	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS512, accessClaims)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	accessTokenSign, err := accessToken.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	// 生成refresh_token，不包含自定义信息
	refreshClaims := CustomClaims{
		MysqlModel.User{
			Id: userInfo.Id,
		},
		jwt.RegisteredClaims{
			Issuer:    "ken",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenExpire)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS512, refreshClaims)
	refreshTokenSign, err := refreshToken.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	// 将token写入redis hash中
	err = authService.SaveAuthTokenRedis(userInfo.Id, accessTokenSign, refreshTokenSign)
	if err != nil {
		return "", "", err
	}

	return bearerSchema + accessTokenSign, bearerSchema + refreshTokenSign, nil
}

// ParseToken 解析token
func ParseToken(authService service.AuthService, tokenString string, grantType string) (*CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 读取当前使用的公钥证书
		secret, err := jwt.ParseRSAPublicKeyFromPEM(certs[curKey].PublicKey)
		if err != nil {
			panic(err)
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// 判断token是否有效，若无效返回错误
		err = authService.CheckAuthTokenRedis(claims.UserInfo.Id, tokenString, grantType)
		if err != nil {
			return nil, err
		}
		return claims, nil
	}

	return nil, errors.New("token invalid")
}
```

最后，在 authService 中会依赖 redis.UserRepository，该仓库会依赖 redisClient 连接，这样便实现了整个依赖注入的过程。

## xlsx 文件导入导出

Excelize 是 Go 语言编写的用于操作 Office Excel 文档基础库。

### 安装

```shell
go get -u github.com/xuri/excelize/v2
```

### xlsx 导出

在 common/excelHandler 目录下创建 excel_export.go 文件，定义了一个 Export()公共方法用于导出文件；
按行设置表头及文件内容，指定请求头进行下载；

```go
package excelHandler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/xuri/excelize/v2"
)

type ExcelExportHandler interface {
	Export(*gin.Context, [][]interface{}, []string, string) error
}
type excelExportHandler struct {
	file *excelize.File
}

func NewExcelExportHandler() (ExcelExportHandler, func(), error) {
	file := excelize.NewFile()
	handler := &excelExportHandler{
		file: file,
	}
	return handler, func() {
		defer file.Close()
	}, nil
}

// Export 文件导出
func (excel *excelExportHandler) Export(c *gin.Context, rows [][]interface{}, headerArr []string, fileName string) error {
	// 获取当前sheet
	sheetName := excel.file.GetSheetName(0)

	// 设置表头
	err := excel.file.SetSheetRow(sheetName, "A1", &headerArr)
	if err != nil {
		return err
	}

	// 设置文件内容
	err = excel.setExcelContent(rows, sheetName, 2)
	if err != nil {
		return err
	}

	// 指定请求头
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	c.Header("Content-Type", "application/octet-stream")

	return excel.file.Write(c.Writer)
}

// 设置文档内容
func (excel *excelExportHandler) setExcelContent(rows [][]interface{}, sheetName string, startRowIndex int) (err error) {
	for i, row := range rows {
		rowIndex := startRowIndex + i
		if err = excel.file.SetSheetRow(sheetName, "A"+strconv.Itoa(rowIndex), &row); err != nil {
			return err
		}
	}
	return nil
}
```

新建一个 service/excel_user_service.go 文件，组装数据调用公共导入方法，代码如下：

```go
// ExportUser 导出用户信息
func (svc *excelUserService) ExportUser(c *gin.Context, userList []MysqlModel.User) error {
	// 文件名
	fileName := "hello.xlsx"
	// 表头
	headerArr := []string{"用户Id", "用户名", "密码", "性别"}

	// 表格数据
	rows := make([][]interface{}, 0, 100)
	for _, user := range userList {
		row := []interface{}{user.Id, user.Username, user.Password, user.Gender}
		rows = append(rows, row)
	}

	// 导出文件
	return svc.excelExportHandler.Export(c, rows, headerArr, fileName)
}
```

### xlsx 导入

通过 router.LoadHTMLGlob("views/index.html")告诉 gin 加载文件，添加上传文件页面路由，并定义了文件导入处理路由；

```go
// Excel文件导出
router.GET("/excel/export", srv.Export())
// 告诉gin框架去哪加载讲台⽂件此处可以使⽤正则表达式
router.LoadHTMLGlob("views/index.html")
// Excel文件导入
router.GET("/excel/import_index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
})
router.POST("/excel/import", srv.Import())
```

views/index.html 文件如下：

```html
<!DOCTYPE html>
<html lang="zh-CN">
  <head>
    <title>上传⽂件⽰例</title>
  </head>
  <body>
    <form
      action="http://127.0.0.1:8080/excel/import"
      method="post"
      enctype="multipart/form-data"
    >
      <input type="file" name="uploadFile" />
      <input type="submit" value="上传" />
    </form>
  </body>
</html>
```

在 excel_user_service.go 文件中，定义 ImportUser()方法，代码如下：

```go
// ImportUser 导入文件
func (svc *excelUserService) ImportUser(c *gin.Context, file io.Reader) ([]MysqlModel.User, error) {
	userList := make([]MysqlModel.User, 0, 100)

	// 表头（用于检测文件内容是否符合要求）
	headerArr := []string{"用户Id", "用户名", "密码", "性别"}

	// 读取文件数据
	rows, err := svc.excelImportHandler.Import(c, file, 0, headerArr)
	if err != nil {
		return userList, err
	}

	// 格式化数据
	for i, row := range rows {
		if i == 0 {
			continue
		}
		user := MysqlModel.User{
			Id:       cast.ToInt(row[0]),
			Username: row[1],
			Password: row[2],
			Gender:   cast.ToInt(row[3]),
		}
		userList = append(userList, user)
	}
	return userList, err
}
```

调用了文件导入公共方法，新增文件 common/excelHandler/excel_import.go 文件，该文件从上传文件 uploadFile 文件打开 Excel，读取文件内容，并根据表头判断文件内容格式，
代码如下：

```go
package excelHandler

import (
	"io"

	"github.com/gin-gonic/gin"

	"github.com/go_example/internal/utils/tools"

	"github.com/pkg/errors"

	"github.com/xuri/excelize/v2"
)

type ExcelImportHandler interface {
	Import(*gin.Context, io.Reader, int, []string) ([][]string, error)
}

type excelImportHandler struct {
}

func NewExcelImportHandler() ExcelImportHandler {
	return &excelImportHandler{}
}

func (excel *excelImportHandler) Import(c *gin.Context, uploadFile io.Reader, sheetIndex int, headerArr []string) ([][]string, error) {
	file, err := excelize.OpenReader(uploadFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 获取当前sheet
	sheetName := file.GetSheetName(sheetIndex)

	// 按行读取文件内容
	rows, err := file.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, errors.New("解析数据为空")
	}

	if len(rows[0]) != len(headerArr) {
		return nil, errors.New("文件数据格式错误")
	}

	// 检查是否包含表头数据
	for _, header := range headerArr {
		if !tools.IsContain(rows[0], header) {
			return nil, errors.New("文件数据格式错误")
		}
	}

	return rows, err
}
```

## 优雅关机处理

热更新时为了不影响正在访问的用户，需要进行优雅关机处理，将 Gin 通过一个 groutine 启动服务，不影响下面的优雅关机处理；

在优雅关机处理中，通过通道接收指定的终止信号量，并读取通道实现是否执行关机操作；

在关机处理中，定义了一个带超时时间的上下文，旧请求应该在超时时间内完成响应，超时后执行关闭服务；

```go
Run: func(cmd *cobra.Command, args []string) {
    // 读取证书内容
    auth.SetCerts()

    // 实例化依赖注入服务
    httpSrv, clean, err := assembly.NewHttpServer()
    if err != nil {
        log.Fatalf("%+v\n", err)
    }
    defer clean()

    // 设置gin的运行环境
    gin.SetMode(meta.EnvMode)

    // 初始化engine
    app := gin.Default()

    // 注册路由
    httpSrv.Register(app)

    addr := viper.GetString("server.http.addr")

    // 自定义server
    srv := &http.Server{
        Addr:    addr,
        Handler: app,
    }

    // 启动监听服务
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %+v\n", err)
        }
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server Shutdown:", err)
    }
    log.Println("Server exiting")
},
```

### 测试效果

在 hello_controller.go 文件中，增加一个超时时间

```go
func (ctr *helloController) Say(c *gin.Context) (int, gin.Negotiate) {
	time.Sleep(5 * time.Second)
	data := ctr.helloSvc.SayHello(c)
	return negotiate.JSON(http.StatusOK, data)
}
```

运行服务，访问 127.0.0.1:8080/hello，立即执行 Ctrl+C 停止 http 服务，此时可以看到 http 服务不会立即停止，访问正常响应后服务关闭。

## WebSocket 服务

Socket 是实现 TCP/IP 连接的 API 接口，调用 Socket 提供的接口，可以实现一次 TCP 连接，客户端和服务端全双工数据交互，主要应用场景如游戏。

### 创建子入口构建 Socket 服务

通过 cobra 生成 socket 服务入口，在项目根目录下执行以下命令：

```shell
cobra add socket
```

同 http 服务构建类似，创建或编辑以下文件，并执行 wire 重新生成 wire_gen.go 文件：

- internal/server/socket.go
- internal/controller/socket_controller.go
- internal/assembly/controller.go
- internal/assembly/server.go

### Socket 服务端

在 socket_controller.go 文件中，通过 github.com/gorilla/websocket 库升级 http 连接为 websocket 连接，并提供服务端服务。

```go
package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type SocketController interface {
	SocketServer(c *gin.Context)
}

type socketController struct {
}

func NewSocketController() SocketController {
	return &socketController{}
}

func (ctr *socketController) SocketServer(c *gin.Context) {
	// 升级为WebSocket
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalf("upGrader,err is %+v", err)
		return
	}
	defer ws.Close()

	// 处理socket请求
	for {
		// 读取客户端的数据为字符串消息类型
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			break
		}

		// 响应客户端的数据也为字符串消息类型
		resData := "recv data:" + string(message)
		ws.WriteMessage(messageType, []byte(resData))
	}
}
```

### 浏览器客户端

在 http 服务中，增加 html 页面，通过前端发起 websocket 连接。

在路由增加加载的 html 文件，代码如下：

```go
func (srv *httpServer) Register(router *gin.Engine) {
    // 告诉gin框架去哪加载讲台⽂件此处可以使⽤正则表达式
    router.LoadHTMLGlob("views/*.html")
    // Excel文件导入
    router.GET("/excel/import_index", func (c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
    })
    // websocket连接
    router.GET("/socket/index2", func (c *gin.Context) {
        c.HTML(http.StatusOK, "socket2.html", nil)
    })
	// ......
}
```

views/socket2.html 文件内容如下：

```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <script>
      window.addEventListener("load", function (evt) {
        var output = document.getElementById("output");
        var input = document.getElementById("input");
        var ws;
        var print = function (message) {
          var d = document.createElement("div");
          d.textContent = message;
          output.appendChild(d);
          output.scroll(0, output.scrollHeight);
        };
        document.getElementById("open").onclick = function (evt) {
          if (ws) {
            return false;
          }
          ws = new WebSocket("ws://192.168.163.1:30000/socket_server");
          ws.onopen = function (evt) {
            print("OPEN");
          };
          ws.onclose = function (evt) {
            print("CLOSE");
            ws = null;
          };
          ws.onmessage = function (evt) {
            print("RESPONSE: " + evt.data);
          };
          ws.onerror = function (evt) {
            print("ERROR: " + evt.data);
          };
          return false;
        };
        document.getElementById("send").onclick = function (evt) {
          if (!ws) {
            return false;
          }
          print("SEND: " + input.value);
          ws.send(input.value);
          return false;
        };
        document.getElementById("close").onclick = function (evt) {
          if (!ws) {
            return false;
          }
          ws.close();
          return false;
        };
      });
    </script>
  </head>
  <body>
    <table>
      <tr>
        <td valign="top" width="50%">
          <p>
            Click "Open" to create a connection to the server, "Send" to send a
            message to the server and "Close" to close the connection. You can
            change the message and send multiple times.
          </p>

          <p></p>
          <form>
            <button id="open">Open</button>
            <button id="close">Close</button>
            <p>
              <input id="input" type="text" value="Hello world!" />
              <button id="send">Send</button>
            </p>
          </form>
        </td>
        <td valign="top" width="50%">
          <div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
        </td>
      </tr>
    </table>
  </body>
</html>
```

访问 127.0.0.1:8080/socket/index2 即可展示页面，点击发送 socket 请求；

以上使用的数据格式为普通字符串类型，一般使用采用 json 数据格式，因此需要对服务端代码进行如下修改：

```go
for {
    // 从ws中读取数据
    var reqData map[string]string
    err = ws.ReadJSON(&reqData)
    if err != nil {
        break
    }

    message := reqData["name"] + "_" + reqData["age"]

    _, resData := negotiate.JSON(http.StatusOK, gin.H{
        "data": gin.H{
            "message": "recv msg：" + string(message),
        },
    })
    // 写入ws数据
    ws.WriteJSON(resData.Data)
}
```

客户端代码也需要进行修改，发送为 json 数据，接收也是 json 数据，如下所示：

```html
<html>
  <head>
    <title>test</title>
  </head>
  <body>
    test ....
  </body>
  <script>
    var ws = new WebSocket("ws://192.168.163.1:30000/socket_server");
    var data = `{"name":"张三","age":"18"}`;

    //连接打开时触发
    ws.onopen = function (evt) {
      console.log("Connection open ...");
      ws.send(data);
    };
    //接收到消息时触发
    ws.onmessage = function (evt) {
      console.log("Received Message: " + evt.data);
    };
    //连接关闭时触发
    ws.onclose = function (evt) {
      console.log("Connection closed.");
    };
  </script>
</html>
```

### 客户端

当需要连接其他项目提供的 socket 服务时，需要编写客户端代码，代码如下：

```go
package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":30000", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/socket_server"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	inputReader := bufio.NewReader(os.Stdin)
	// 发送数据
	go func() {
		for {
			input, _ := inputReader.ReadString('\n')
			inputInfo := strings.Trim(input, "\r\n")
			err = c.WriteMessage(websocket.TextMessage, []byte(inputInfo))
			if err != nil {
				log.Printf("WriteMessage err:%+v", err)
				break
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// 通过发送一个关闭消息，在等待1秒后关闭连接
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
```

### WebSocket 服务增加心跳检测

WebSocket 服务启动后，若客户端不关闭连接且长时间不通信，服务端 TCP 连接会一直保持。当 Socket 连接数不断增多时，对服务器资源是个巨大的消耗，因此需要及时关闭连接。

golang 并没有提供针对长连接的方法，因此可以通过心跳检测策略来实现：

- 当客户端有数据请求时，则重新设置服务端连接的超时时间为 30 秒；
- 当客户端超过 30 秒未请求，则会超时并关闭服务端 TCP 连接；

通过定义一个 channel，定义一个 goroutine 若有数据则写入通道，另定义一个 goroutine 从通道里取值，若能取出，则重新设置连接超时时间，代码如下：

```go
func (ctr *socketController) SocketServer(c *gin.Context) {
	// 升级为WebSocket
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalf("upGrader,err is %+v", err)
		return
	}
	defer conn.Close()

	// 处理socket请求
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		resData := "recv data:" + string(message)
		fmt.Println(resData)
		conn.WriteMessage(messageType, []byte(resData))

		// 定义一个chan
		heartChan := make(chan byte)
		go heartBeating(message[:1], heartChan)
		go heartHandler(conn, heartChan, 30)
	}
}

// 如果有消息则写入通道
func heartBeating(msg []byte, heartChan chan byte) {
	for _, v := range msg {
		heartChan <- v
	}
}

// 保活
func heartHandler(conn *websocket.Conn, heartChan chan byte, timeout int) {
	select {
	case <-heartChan:
		conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		conn.SetWriteDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	}
}

```

### WebSocket 增加客户端来源检查

Socket 服务端启动后，只需要知道 IP 和 Port 后，就可以连接，这样就不安全，可以增加一个来源检查。以下代码通过 IP 地址进行判断：

```go
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		clientIp := tools.GetClientIp(r)
		if len(meta.SocketWhiteIpList) > 0 {
			for _, ip := range meta.SocketWhiteIpList {
				pattern := strings.ReplaceAll(strings.ReplaceAll(ip, ".", "\\."), "*", ".*")
				match, _ := regexp.MatchString(pattern, clientIp)
				if match {
					return true
				}
			}
			return false
		}
		return true
	},
}
```

## gRPC 服务

gRPC 是一种实现 rpc 框架，gRPC 能让我们更容易编写跨语言的分布式代码，使用 protobuf 协议效率更高、更规范。

### 编写 proto 文件

gRPC 使用 protobuf 协议进行通信，在项目目录下创建 common/protobuf/hello/hello.proto 文件，代码如下：

```protobuf
syntax = "proto3";

package hello;

option go_package="common/protobuf/hello;pb";

message HelloRequest{
  int32 id=1;
}

message HelloResponse{
  int32 id = 1;
  string name = 2;
}

service HelloService{
  rpc SayHello(HelloRequest) returns (HelloResponse);
}
```

### 生成 pb.go 文件

从https://github.com/protocolbuffers/protobuf/releases下载对应的protoc编译器，解压拷贝到$GOPATH/bin目录下；
安装生成 go 语言代码插件

```shell
go get -u github.com/golang/protobuf/protoc-gen-go
```

在项目根目录下执行命令：

```shell
protoc  --go_out=plugins=grpc:. common/protobuf/hello/hello.proto
```

这样就在 common/protobuf/hello 目录下生成了 hello.pb.go 文件。

### 创建 gRPC 服务入口

使用 cobra 增加子命令：

```shell
cobra add grpc
```

在 internal/server 目录下创建一个 grpc.go 文件，该文件定义了一个 grpcServer，包含注册服务到 grpc 及服务的一个实现方法，代码如下：

```go
package server

import (
	"context"

	pb "github.com/go_example/common/protobuf/hello"
	"google.golang.org/grpc"
)

type GrpcServer interface {
	SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error)
	Register(server *grpc.Server)
}
type grpcServer struct {
	pb.UnimplementedHelloServiceServer
}

func NewGrpcServer() GrpcServer {
	return &grpcServer{}
}

// Register 注册服务到grpc
func (srv *grpcServer) Register(server *grpc.Server) {
	pb.RegisterHelloServiceServer(server, &grpcServer{})
}

// SayHello grpc服务
func (srv *grpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := "world"
	if in.Id != 1 {
		name = "gRPC"
	}
	return &pb.HelloResponse{
		Id:   in.Id,
		Name: "hello " + name,
	}, nil
}
```

### 启动 grpc 服务

将 protobuf 定义的 service 服务注册到 grpc 中，并启动服务。代码如下：

```go
/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"google.golang.org/grpc"

	"github.com/go_example/internal/assembly"
	"github.com/spf13/cobra"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "A brief description of your command",
	Long:  `grpc server`,
	Run: func(cmd *cobra.Command, args []string) {
		grpcSrv, cleanup, err := assembly.NewGrpcServer()
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		defer cleanup()

		// 1.监听端口
		addr := viper.GetString("server.grpc.addr")
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("Listen err:%+v", err)
		}

		// 2.创建一个grpc服务
		app := grpc.NewServer()

		// 3.注册服务
		grpcSrv.Register(app)

		// 4.启动服务
		go func() {
			err = app.Serve(listen)
			if err != nil {
				log.Fatalf("Serve err:%+v", err)
			}
		}()

		// 优雅关闭服务
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
		<-quit
		app.GracefulStop()
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}
```

对 internal/assembly/server.go 文件增加 NewGrpcServer 方法后，执行 wire 重新生成文件：

```go
func NewGrpcServer() (server.GrpcServer, func(), error) {
	panic(wire.Build(
		server.NewGrpcServer,
	))
}
```

这样就完成了 grpc 服务端，执行 go run main.go grpc 即可。

### gRPC 客户端

创建一个 controller，对其进行 wire 注入，配置 http 服务的路由，最终代码如下：

```go
package controller

import (
	"net/http"

	"github.com/spf13/cast"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/utils/negotiate"

	pb "github.com/go_example/common/protobuf/hello"
	"google.golang.org/grpc"
)

type GrpcClientController interface {
	HelloGrpc(c *gin.Context) (int, gin.Negotiate)
}

type grpcClientController struct {
}

func NewGrpcClientController() GrpcClientController {
	return &grpcClientController{}
}

func (ctr *grpcClientController) HelloGrpc(c *gin.Context) (int, gin.Negotiate) {
	idStr := c.DefaultQuery("id", "1")
	// 连接服务
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return negotiate.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"name": "",
			},
		})
	}

	// 创建grpc client
	client := pb.NewHelloServiceClient(conn)

	// 调用服务
	resp, err := client.SayHello(c, &pb.HelloRequest{Id: cast.ToInt32(idStr)})
	if err != nil {
		return negotiate.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"name": "",
			},
		})
	}
	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"name": resp.Name,
		},
	})
}
```

启动 http 服务，访问对应的路由即可连接 gRPC 服务。

### TLS 加密认证

在上面的示例中，服务端未使用加密认证，客户端通过指定 insecure.NewCredentials()跳过加密认证，这样安全性不高，为此这里使用 TLS 证书加密。
在 1.15 版本后，需要执行 SAN 证书通信。

#### 生成证书

copy 一份 openssl.cnf 文件到要生成证书的目录下，对文件做如下修改：

```text
[ CA_default ]
...
# 打开下面这行
copy_extensions = copy
...
[ req ]
...
# 打开下面这行
req_extensions = v3_req
...
[ V3_req ]
...
# 增加如下一行
subjectAltName = @alt_names
...

# 增加[ alt_names ]配置段，这里设置域名
[ alt_names ]
DNS.1 = *.org.example.com
DNS.2 = *.example.com
```

生成根证书

```shell
# 生成根证书私钥
openssl genrsa -out ca.key 2048
# 生成CA证书请求
openssl req -new -key ca.key -out ca.csr -subj "/C=cn/OU=myorg/O=mytest/CN=myname"
# 自签名得到根证书
openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.crt
```

生成服务端证书

```shell
# 生成服务端证书私钥
openssl genrsa -out server.key 2048
# 根据私钥server.key生成证书请求文件server.csr
openssl req -new -nodes -key server.key -out server.csr -subj "/C=cn/OU=myserver/O=servercomp/CN=servername" -config ./openssl.cnf -extensions v3_req
# 请求CA对证书请求文件签名，生成最终的证书文件
openssl x509 -req -days 365 -in server.csr -out server.pem -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req
```

#### 代码实现

下面将使用到 ca.crt、server.key、server.pem 文件，复制到/assets/certs/grpc_tls/目录下；
服务端代码修改

```go
creds, _ := credentials.NewServerTLSFromFile("./assets/certs/grpc_tls/server.pem", "./assets/certs/grpc_tls/server.key")
app := grpc.NewServer(grpc.Creds(creds))
```

客户端代码修改

```go
creds, err := credentials.NewClientTLSFromFile("./assets/certs/grpc_tls/server.pem", "a.example.com")
conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds))
```

运行服务端和客户端即可。

### TLS 加密证书认证升级

同样的方式生成客户端证书，得到 client.key、client.crs、client.pem 文件，复制文件到 grpc_tls 目录下；

```Bash
openssl genrsa -out client.key 2048
openssl req -new -nodes -key client.key -out client.csr -subj "/C=cn/OU=myclient/O=clientcomp/CN=clientname" -config ./openssl.cnf -extensions v3_req
openssl x509 -req -days 365 -in client.csr -out client.pem -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req
```

服务端：

```Go
// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
certificate, err := tls.LoadX509KeyPair("./assets/certs/grpc_tls/server.pem", "./assets/certs/grpc_tls/server.key")
if err != nil {
  log.Fatalf("tls.LoadX509KeyPair err:%+v", err)
}
// 创建一个新的、空的 CertPool
certPool := x509.NewCertPool()
ca, err := ioutil.ReadFile("./assets/certs/grpc_tls/ca.crt")
if err != nil {
  log.Fatalf("ioutil.ReadFile err:%+v", err)
}
// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
if ok := certPool.AppendCertsFromPEM(ca); !ok {
  log.Fatalf("certPool.AppendCertsFromPEM err:%+v", err)
}
// credentials.NewTLS:构建基于 TLS 的 TransportCredentials 选项
creds := credentials.NewTLS(&tls.Config{ // Config 结构用于配置 TLS 客户端或服务器
  Certificates: []tls.Certificate{certificate}, // 设置证书链，允许包含一个或多个
  // tls.RequireAndVerifyClientCert 表示 Server 也会使用 CA 认证的根证书对 Client 端的证书进行校验
  // ClientAuth: tls.RequireAndVerifyClientCert, // 要求必须校验客户端的证书
  ClientAuth:   tls.RequestClientCert,          // 请求客户端证书，握手期间不要求发送证书(服务发现时健康检查使用这个参数)
  ClientCAs:  certPool,                       // 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
})

app := grpc.NewServer(grpc.Creds(creds))
```

客户端（基本与服务端获取证书一样）：

```Go
certificate, err:= tls.LoadX509KeyPair("./assets/certs/grpc_tls/client.pem", "./assets/certs/grpc_tls/client.key")
if err != nil {
  log.Fatalf("tls.LoadX509KeyPair err:%+v", err)
}
// 创建一个新的、空的 CertPool
certPool := x509.NewCertPool()
ca, err:= ioutil.ReadFile("./assets/certs/grpc_tls/ca.crt")
if err != nil {
  log.Fatalf("ioutil.ReadFile err:%+v", err)
}
// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
if ok := certPool.AppendCertsFromPEM(ca); !ok {
  log.Fatalf("certPool.AppendCertsFromPEM err:%+v", err)
}

creds := credentials.NewTLS(&tls.Config{
  Certificates: []tls.Certificate{certificate},
  ServerName:   "www.example.com",
  RootCAs:      certPool,
})

conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds))
```

这样我们使用 CA 颁发的根证书对客户端、服务端的证书进行了签发。进一步的提高了两者的通讯安全。

### gRPC 基于 Token 认证

gRPC 还为每个 gRPC 方法调用提供了认证支持，这样就基于用户 Token 对不同的方法访问进行权限管理。
要实现对每个 gRPC 方法进行认证，需要实现`grpc.PerRPCCredentials`接口，在 internal/lib/auth 下创建一个 grpc.go 文件，代码如下：

```Go
package auth

import (
  "context"
  "fmt"

  "google.golang.org/grpc/metadata"
)

type Authentication interface {
  GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
  RequireTransportSecurity() bool
  Auth(ctx context.Context) error
}

type authentication struct {
  User     string
  Password string
}

func NewAuthentication(user, password string) Authentication {
  return &authentication{
    User:     user,
    Password: password,
  }
}

// GetRequestMetadata 返回地认证信息
func (auth *authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
  return map[string]string{"user": auth.User, "password": auth.Password}, nil
}

// RequireTransportSecurity 返回true需要证书认证
func (auth *authentication) RequireTransportSecurity() bool {
  return true
}

// Auth 进行权限判断
func (auth *authentication) Auth(ctx context.Context) error {
  md, ok := metadata.FromIncomingContext(ctx)
  if !ok {
    return fmt.Errorf("missing credentials")
  }
  var appId, appKey string
  if val, ok := md["user"]; ok {
    appId = val[0]
  }
  if val, ok := md["password"]; ok {
    appKey = val[0]
  }
  if appId != auth.User || appKey != auth.Password {
    return fmt.Errorf("invaild token")
  }
  return nil
}

```

客户端在之前基础上修改：

```Go
grpcAuth := auth.NewAuthentication("root", "root123")
conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(grpcAuth))
```

服务端在提供的方法修改：

```Go
// SayHello grpc服务
func (srv *grpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
  grpcAuth := auth.NewAuthentication("root", "root123")
  if err := grpcAuth.Auth(ctx); err != nil {
    return nil, err
  }
  name := "world"
  if in.Id != 1 {
    name = "gRPC"
  }
  return &pb.HelloResponse{
    Id:   in.Id,
    Name: "hello " + name,
  }, nil
}
```

## 服务注册与服务发现

Consul 提供服务注册与服务发现 API 接口，可实现 gRPC 分布式服务负载均衡。

### 安装

官方提供了 golang 版本的库，通过下面命令进行安装：

```shell
go get github.com/hashicorp/consul/api
```

创建 common/consulClient/consulClient.go 文件，通过 consul 官方提供的 api 包调用对应的服务注册与发现接口。

```go
package consulClient

import (
	"fmt"

	consulApi "github.com/hashicorp/consul/api"
)

type ConsulClient interface {
	RegisterService(serviceName string, ip string, port int) error
	DeregisterService(serviceId string) error
}

type consulClient struct {
	Client *consulApi.Client
}

func NewClient(addr string) (ConsulClient, error) {
	cfg := consulApi.DefaultConfig()
	cfg.Address = addr
	cli, err := consulApi.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &consulClient{
		Client: cli,
	}, nil
}

// RegisterService 注册服务
func (cli *consulClient) RegisterService(serviceName string, ip string, port int) error {
	// 服务健康检查
	check := &consulApi.AgentServiceCheck{
		Interval:                       "10s",
		Timeout:                        "10s",
		GRPC:                           fmt.Sprintf("%s:%d", ip, port),
		GRPCUseTLS:                     true,
		TLSSkipVerify:                  true,
		DeregisterCriticalServiceAfter: "30s", //check失败后30秒删除本服务
	}
	return cli.Client.Agent().ServiceRegister(&consulApi.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port),
		Name:    serviceName,
		Tags:    []string{"my_grpc"},
		Port:    port,
		Address: ip,
		Check:   check,
	})
}

// DeregisterService 注销服务
func (cli *consulClient) DeregisterService(serviceId string) error {
	return cli.Client.Agent().ServiceDeregister(serviceId)
}
```

在 cmd/grpc.go 文件中，添加服务注册与服务注销，同时增加健康检查。

```go
Run: func(cmd *cobra.Command, args []string) {
    app := grpc.NewServer(grpc.Creds(creds))

    // 开启健康检查
    healthServer := health.NewServer()
    healthServer.SetServingStatus(meta.HEALTHCHECK_SERVICE, healthpb.HealthCheckResponse_SERVING)
    healthpb.RegisterHealthServer(app, healthServer)

    // 3.注册服务
    grpcSrv.Register(app)

    // 注册服务到consul
    consulClient, serviceIdArr, err := grpcSrv.RegisterConsul()
    if err != nil {
        log.Fatalf("RegisterConsul err%+v\n", err)
    }

    // 4.启动服务
    go func() {
        err = app.Serve(listen)
        if err != nil {
        log.Fatalf("Serve err:%+v\n", err)
        }
    }()

    // 优雅关闭服务
    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
    <-quit

    app.GracefulStop()

    // 注销consul服务
    for _, serviceId := range serviceIdArr {
        if err := consulClient.DeregisterService(serviceId); err != nil {
            log.Fatalf("consulClient.DeregisterService err:%+v", err)
        }
    }
}
```

客户端修改连接地址为 consul 地址，并指定负载均衡及服务健康检查。

```go
conn, err := grpc.Dial("consul://192.168.163.131:8500/hello?wait=10s", grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"HealthCheckConfig": {"ServiceName": "%s"}}`, meta.HEALTHCHECK_SERVICE)), grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(grpcAuth))
```

## 内存缓存

项目选用 go-cache 作为内存缓存。 ###安装

```shell
go get github.com/patrickmn/go-cache
```

### 代码实现

在 meta.go 文件中定义全局变量

```go
// CacheDriver go-cache缓存对象
var CacheDriver = cache.New(5*time.Minute, 10*time.Minute)
```

在 internal/repository/cache/创建 cache_repository.go，文件定义了获取 key 的方法及其他公共方法。

```go
package cache

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

// GetCacheKey 获取缓存key
func GetCacheKey(key string, vals ...interface{}) (string, error) {
	cacheKey := viper.GetString("cacheKeys." + key)
	if cacheKey == "" {
		return "", errors.New("key不存在")
	}
	return fmt.Sprintf(cacheKey, vals...), nil
}
```

创建一个 user_repository.go 文件，包含读取和写入 go-cache。

```go
package cache

import (
	"errors"
	"time"

	"github.com/go_example/internal/meta"
	MysqlModel "github.com/go_example/internal/model/mysql"
)

type UserRepository interface {
	SetUserInfo(uid int, userInfo MysqlModel.User) error
	GetUserInfo(uid int) (MysqlModel.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (repo *userRepository) SetUserInfo(uid int, userInfo MysqlModel.User) error {
	cacheKey, err := GetCacheKey("userInfo", uid)
	if err != nil {
		return err
	}
	meta.CacheDriver.Set(cacheKey, userInfo, time.Hour)
	return nil
}

func (repo *userRepository) GetUserInfo(uid int) (MysqlModel.User, error) {
	cacheKey, err := GetCacheKey("userInfo", uid)
	if err != nil {
		return MysqlModel.User{}, err
	}
	data, isExist := meta.CacheDriver.Get(cacheKey)
	if !isExist {
		return MysqlModel.User{}, errors.New("key不存在")
	}
	return data.(MysqlModel.User), nil
}
```

在 hello_service.go 文件中使用，代码如下：

```go
func (svc *helloService) SayHello(c *gin.Context) map[string]string {
	uid := cast.ToInt(c.DefaultQuery("id", "0"))
	// 先从gocache中读取数据
	user, err := svc.userCacheRepo.GetUserInfo(uid)
	if err != nil {
		user, err = svc.userRepo.GetUserInfoById(uid)
		if err == nil {
			// 写入数据到gocache中
			_ = svc.userCacheRepo.SetUserInfo(uid, user)
		}
	}
	value := svc.userRedisRepo.GetValue("aa")
	return map[string]string{
		"hello": "world，golang",
		"env":   viper.GetString("server.mode"),
		"user":  user.Username,
		"value": value,
	}
}
```

## 实现 MongoDB 连接

安装 MongoDB Go 驱动包

```shell
go get go.mongodb.org/mongo-driver/mongo
```

添加 mongodb 连接配置，如果是分片集群则为多个地址以逗号隔开，若为 replicas 集群在后面还要加/?replicaSet=myRepl

```yaml
mongodb:
  # addr格式：mongodb://username:password@addr1:port2,addr2:port2
  addr: "mongodb://192.168.163.131:27017"
  max_open: 20
```

创建 common/mongoClient/mongo.go 文件，定义了连接 Mongo 的方法，由于 mongo 包中没有提供方法的 interface，需要将 mongo 包 client.go 定义的方法定义复制到 SingleClient interface，代码如下：

```go
package mongoClient

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	StartSession(opts ...*options.SessionOptions) (mongo.Session, error)
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
	ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error)
	ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error)
	UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
	UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo.SessionContext) error) error
	Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error)
	NumberSessionsInProgress() int
}
type mongoClient struct {
	*mongo.Client
}

type MongoConfig struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	MaxOpen  uint64 `json:"max_open" mapstructure:"max_open"`
}

func NewMongoClient(cfg MongoConfig) (MongoClient, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 得到连接信息
	clientOption := options.Client().ApplyURI(cfg.Addr)
	if cfg.MaxOpen > 0 {
		clientOption.SetMaxPoolSize(cfg.MaxOpen)
	}

	// 连接
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return &mongoClient{Client: client}, func() {
		_ = client.Disconnect(ctx)
	}, nil
}

```

在 internal/respository/mongodb 目录下创建 user_repository.go 文件，该文件依赖 mongo 连接，定义了 mongo CRUD 文档的具体操作。

```go
package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/go_example/internal/meta"
	MongoModel "github.com/go_example/internal/model/mongodb"
)

type UserRepository interface {
	SetUserInfo(int, string, string) error
	GetUserInfo(int) (MongoModel.User, error)
}

type userRepository struct {
	client     meta.MongoClient
	database   string
	collection string
}

func NewUserRepository(client meta.MongoClient) UserRepository {
	return &userRepository{
		client:     client,
		database:   "test",
		collection: "user",
	}
}

func (repo *userRepository) SetUserInfo(id int, username, password string) error {
	userDocument := MongoModel.User{
		Id:       id,
		Username: username,
		Password: password,
	}
	_, err := repo.client.Database(repo.database).Collection(repo.collection).InsertOne(context.TODO(), userDocument)
	if err != nil {
		return nil
	}
	return nil
}

func (repo *userRepository) GetUserInfo(uid int) (MongoModel.User, error) {
	var userInfo MongoModel.User

	err := repo.client.Database(repo.database).Collection(repo.collection).FindOne(context.TODO(), bson.D{}).Decode(&userInfo)
	if err != nil {
		return userInfo, err
	}
	return userInfo, nil
}

```

最后，在 hello_service.go 文件中使用，代码如下：

```go
// mongodb操作数据
if user.Username != "" {
    _ = svc.userMongoRepo.SetUserInfo(uid, user.Username, user.Password)
}
userInfo, _ := svc.userMongoRepo.GetUserInfo(uid)
```

这样就可以愉快地使用 mongodb 了。

## pprof 性能分析器

安装

```shell
go get -u github.com/felixge/fgprof
```

采用 fgprof 包代替官方提供的 pprof，仅需在 cmd/http.go 文件中添加如下代码：

```go
import(
    _ "net/http/pprof"
    "github.com/felixge/fgprof"
)
// ......
Run: func(cmd *cobra.Command, args []string) {
    // pprof性能分析
    if !env.IsReleasing() {
        http.DefaultServeMux.Handle("/debug/fgprof", fgprof.Handler())
        go func() {
            log.Println(http.ListenAndServe(":6060", nil))
        }()
    }
}
```

打开浏览器访问http://127.0.0.1:6060/debug/pprof/ 即可查看概览。若要显示图形化界面，需要安装 Graphviz，安装完成后执行如下代码：

```Bash
go tool pprof -http=:6061 http://localhost:6060/debug/pprof/profile?seconds=60
```

等待 60 秒后，会自动跳转到图形界面。

## RPC 熔断和降级

微服务架构，服务间调用，当一个服务发生故障，若不对其进行熔断或降级处理，容易导致整个系统崩溃，造成雪崩。
这里使用 hystrix-go 来实现服务熔断和降级。

### 安装

```shell
go get github.com/afex/hystrix-go
```

### 代码实现

在 gRPC 调用方，增加拦截器，在拦截器中实现断路器的作用。增加一个拦截器方法，并在 grpc.Dail()方法增加拦截器。

```go
// UnaryClientInterceptor grpc拦截器
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		hystrix.ConfigureCommand(method, hystrix.CommandConfig{
			Timeout:                1000,
			MaxConcurrentRequests:  10,
			RequestVolumeThreshold: 10,
			SleepWindow:            5000,
			ErrorPercentThreshold:  50,
		})
		return hystrix.Do(method, func() (err error) {
			return invoker(ctx, method, req, reply, cc, opts...)
		}, func(err error) error {
			// 因为这里是在调用方实现熔断，若服务不可用可以发邮件通知或什么都不做，实现降级
			return nil
		})
	}
}
// ...省略代码
conn, err := grpc.Dial("consul://"+consulAddr+"/hello?wait=10s", grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"HealthCheckConfig": {"ServiceName": "%s"}}`, meta.HEALTHCHECK_SERVICE)), grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(grpcAuth), grpc.WithUnaryInterceptor(UnaryClientInterceptor()))
```

## 统一错误码

当前返回错误码与成功返回公用一个结构体，存在冗余等问题。为了更加清晰的分辨错误返回和成功返回，使用两种结构体分别处理，并统一错误码。

### 错误码服务

在目录 internal/lib/errorAssets 下，创建 error_list.go，该文件规定了哪些错误码

```go
package errorAssets

import "github.com/ken-house/go-contrib/prototype/errorAssets"

var (
	// ERR_PARAM 公共错误
	ERR_PARAM     = errorAssets.NewError(10000, "参数错误")
	ERR_SYSTEM    = errorAssets.NewError(10001, "系统错误")
	ERR_CERT      = errorAssets.NewError(10002, "证书错误")
	ERR_CALL_FUNC = errorAssets.NewError(10003, "调用方法出错")
	ERR_DIAL      = errorAssets.NewError(10004, "连接错误")

	// ERR_LOGIN 登录注册认证相关错误
	ERR_LOGIN         = errorAssets.NewError(20000, "用户名或密码不正确")
	ERR_LOGIN_REMOTE  = errorAssets.NewError(20001, "账号已在其他设备登录")
	ERR_LOGIN_FAILURE = errorAssets.NewError(20002, "当前登录已失效，请尝试请求refresh_token获取新令牌")
	ERR_REFRESH_TOKEN = errorAssets.NewError(20003, "刷新令牌失败，请重新登录")

	// ERR_EXPOSE 导入导出相关错误
	ERR_EXPORT     = errorAssets.NewError(30000, "导出失败")
	ERR_IMPORT     = errorAssets.NewError(30001, "导入失败")
	ERR_FILE_PARSE = errorAssets.NewError(30002, "文件解析失败")
)
```

创建 error_no.go 文件，定义了返回错误的结构体及方法。

```go
package errorAssets

type Level int

// 提示错误显示效果
const (
	LevelToast Level = 1 // toast提示
	LevelPopup Level = 2 // 弹窗提示
)

type ErrorNo interface {
	GetTitle() string
	GetCode() int
	ToastError() errorNo
	PopupError(popupTitle, popupContent string, popupStyle int) errorNo
}

type errorNo struct {
	Error error `json:"error"`
}

type error struct {
	Code         int    `json:"code,string"`   // 业务编码
	Level        Level  `json:"level,string"`  // 弹窗提示类型
	Message      string `json:"message"`       // level=1 toast标题
	PopupTitle   string `json:"popup_title"`   // level=2 弹框标题
	PopupContent string `json:"popup_content"` // level=2 弹框内容
}

func NewError(code int, title string) ErrorNo {
	return &errorNo{
		Error: error{
			Code:    code,
			Level:   LevelToast,
			Message: title,
		},
	}
}

func (err *errorNo) GetTitle() string {
	return err.Error.Message
}

func (err *errorNo) GetCode() int {
	return err.Error.Code
}

// ToastError toast提示返回结构
func (err *errorNo) ToastError() errorNo {
	return errorNo{
		error{
			Code:    err.Error.Code,
			Level:   LevelToast,
			Message: err.Error.Message,
		},
	}
}

// PopupError 弹窗返回结构
func (err *errorNo) PopupError(popupTitle, popupContent string, popupStyle int) errorNo {
	return errorNo{
		error{
			Code:         err.Error.Code,
			Level:        LevelPopup,
			PopupTitle:   popupTitle,
			PopupContent: popupContent,
		},
	}
}
```

### 错误码调用

```go
return negotiate.JSON(http.StatusOK, errorAssets.ERR_REFRESH_TOKEN.ToastError())
```

## 引用自有包

将项目目录下的 common 目录替换成 github.com/ken-house/go-contrib 库，在 internal/lib 目录下增加 errorAssets 错误码包，实现项目可自定义错误码。

## zap 高性能日志

zap 是 uber 开源的 Go 高性能日志库，是非常快的、结构化的，分日志级别的 Go 日志库。

### 默认使用

使用 zap 提供的配置

```go
// SimpleLogger 使用zap包自带的配置文件
func SimpleLogger(outPutPaths []string) {
	var logger *zap.Logger
	var err error

	config := zap.NewProductionConfig()
	config.Encoding = "console"

	// 增加自定义日志记录位置
	if len(outPutPaths) > 0 {
		outPutPathsArr := config.OutputPaths
		for _, filePath := range outPutPaths {
			_, err = tools.FileNotExistAndCreate(filePath)
			if err != nil {
				log.Fatalln(err)
			}
			outPutPathsArr = append(outPutPathsArr, filePath)
		}

		config.OutputPaths = outPutPathsArr
		config.ErrorOutputPaths = outPutPathsArr
	}

	// 更改时间编码
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig = encoderConfig

	// 保证生产环境和其他环境日志存储格式一致，仅日志等级不同
	if env.IsReleasing() {
		logger, err = config.Build()
	} else {
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		config.Development = true

		logger, err = config.Build()
	}
	if err != nil {
		log.Fatalln(err)
	}
	defer logger.Sync()

	// 注册全局的单例的logger
	zap.ReplaceGlobals(logger)
	// 改变全局的标准库的log的输出，将其通过zap.Logger来输出
	zap.RedirectStdLog(logger)
}
```

### 自定义配置

支持日志文件切割归档

```go
// CustomLogger 自定义zap日志，支持日志切割归档
func CustomLogger(lumberjackLogger *lumberjack.Logger, outPutFile string) {
	encoder := getEncoder()
	writeSyncer := getWriteSyncer(lumberjackLogger, outPutFile)

	logLevel := zapcore.DebugLevel
	if env.IsReleasing() {
		logLevel = zapcore.InfoLevel
	}
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(logLevel))

	defer logger.Sync()

	// 注册全局的单例的logger
	zap.ReplaceGlobals(logger)
	// 改变全局的标准库的log的输出，将其通过zap.Logger来输出
	zap.RedirectStdLog(logger)
}

// 获取编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 日志写入目标，使用lumberjack进行日志切割
func getWriteSyncer(lumberjackLogger *lumberjack.Logger, outPutFile string) zapcore.WriteSyncer {
	if lumberjackLogger != nil { // 使用lumberjack进行日志切割
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberjackLogger))
	} else { // 不使用切割
		// 确保目录存在，不存在则创建目录
		if outPutFile == "" {
			outPutFile = fmt.Sprintf("./logs/log_%s.log", time.Now().Format("20060102"))
		}
		file, err := tools.FileNotExistAndCreate(outPutFile)
		if err != nil {
			log.Fatalln(err)
		}
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(file))
	}
}
```

在 cmd/root.go 中，初始化日志。

```go
// 初始化日志
func initLog() {
    // 普通日志记录
    outputFile := fmt.Sprintf("./logs/log_%s.log", time.Now().Format("20060102"))
    //zapLogger.SimpleLogger([]string{outputFile})

    // 支持日志文件切割
    zapLogger.CustomLogger(&lumberjack.Logger{
    Filename:   outputFile,
    MaxSize:    10,
    MaxAge:     7,
    MaxBackups: 5,
    LocalTime:  false,
    Compress:   false,
    }, "")
}
```

在其他要记录日志的地方可以直接使用 zap.L()写日志信息。

## Gin 框架参数验证

接口接收参数，控制器使用参数绑定，gin 框架自动通过 binding 标签对字段进行参数验证，使用系统自带验证器，也可以自定义验证器如 validatePassword。

```go
// LoginForm 登录表单数据
type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required,max=30"`
	Password string `form:"password" json:"password" binding:"required,validatePassword"`
}

func (ctr *authController) Login(c *gin.Context) (int, gin.Negotiate) {
    var paramData model.LoginForm
    if err := c.ShouldBind(&paramData); err != nil {
    return negotiate.JSON(http.StatusOK, errorAssets.ERR_PARAM.ToastError())
    }
	// ...
}
```

### 自定义验证器

上面示例使用了 validatePassword 自定义验证器，需要对验证器进行定义并注册。
定义验证器方法，在 internal/meta 目录下创建 validator.go，内容如下：

```go
package meta

import (
	"github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
)

// 正则匹配字符串验证器公共方法
func regexpStringValidatorFunc(pattern string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		re := regexp2.MustCompile(pattern, 0)
		isMatch, _ := re.MatchString(value)
		return isMatch
	}
}

// ValidatePhone 验证手机号
var ValidatePhone = regexpStringValidatorFunc(`^1[3-9][0-9]{9}$`)

// ValidatePassword 验证密码必须为6-16位字母+数字组合
var ValidatePassword = regexpStringValidatorFunc(`^(?=.*[0-9])(?=.*[a-zA-Z])([0-9a-zA-Z]{6,16})$`)

// ValidateUsername 验证注册用户名必须为6-30位字母或数字
var ValidateUsername = regexpStringValidatorFunc(`^[0-9A-Za-z]{6,30}$`)

// ValidateVerifyCode 验证码必须为6位数字
var ValidateVerifyCode = regexpStringValidatorFunc(`^[0-9]{6}$`)

```

还需要将自定义验证器注册到 gin 框架中，修改 cmd/root.go，改动如下：

```go
func init() {
    // 初始化配置文件
    cobra.OnInitialize(initConfig, initLog, initValidator)

    rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
// ...
// 参数验证器
func initValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validatePhone", meta.ValidatePhone)
		v.RegisterValidation("validatePassword", meta.ValidatePassword)
		v.RegisterValidation("validateUsername", meta.ValidateUsername)
		v.RegisterValidation("validateVerifyCode", meta.ValidateVerifyCode)
	}
}
```

## Jenkins 服务

Jenkins 是最流行的开源持续集成系统。

### 安装

```shell
go get github.com/bndr/gojenkins
```

### 使用

controller 目录下创建 jenkins_controller.go，代码如下：

```go
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/service"
	"github.com/go_example/internal/utils/negotiate"
	"go.uber.org/zap"
	"net/http"
)

type JenkinsController interface {
	Index(c *gin.Context) (int, gin.Negotiate)
}

type jenkinsController struct {
	jenkinsSvc service.JenkinsService
}

func NewJenkinsController(
	jenkinsSvc service.JenkinsService,
) JenkinsController {
	return &jenkinsController{
		jenkinsSvc: jenkinsSvc,
	}
}

func (ctr *jenkinsController) Index(c *gin.Context) (int, gin.Negotiate) {
	//taskList, err := ctr.jenkinsSvc.GetQueuePendingTaskList(c)
	//if err != nil {
	//	zap.L().Error("jenkinsSvc.GetQueuePendingTaskList", zap.Error(err))
	//}
	//for _, taskId := range taskList {
	//	res, err := ctr.jenkinsSvc.CancelQueueTaskById(c, taskId)
	//	if err != nil {
	//		zap.L().Error("jenkinsSvc.CancelQueueTaskById", zap.Error(err))
	//	}
	//	fmt.Printf("停止队列中的任务%d,res:%v", taskId, res)
	//}

	//nextBuildNum, status, err := ctr.jenkinsSvc.GetJobLatestStatusAndNextBuildNumByName(c, "go-test")
	//if err != nil {
	//	zap.L().Error("jenkinsSvc.GetJobLatestStatusAndNextBuildNumByName", zap.Error(err))
	//}
	//fmt.Println(nextBuildNum, status)
	log, err := ctr.jenkinsSvc.GetJobLastBuildLog(c, "go-test")
	if err != nil {
		zap.L().Error("jenkinsSvc.GetJobLastBuildLog", zap.Error(err))
	}

	//resMap, err := ctr.jenkinsSvc.GetAllJobLatestStatus(c)
	//if err != nil {
	//	zap.L().Error("jenkinsSvc.ValidateAllJobStatus", zap.Error(err))
	//}
	//err := ctr.jenkinsSvc.CreateJobFolder(c, "xudttest")
	//fmt.Println(err)
	//err := ctr.jenkinsSvc.CreateJobFolder(c, "folder1_test", "xudttest")
	//fmt.Println(err)
	//job, err := ctr.jenkinsSvc.CreateJobInFolder(c, "./assets/jenkins/job.xml", "job_test", "xudttest", "folder1_test")
	//fmt.Println(job, err)
	//job := ctr.jenkinsSvc.RenameJob(c, "xudttest/job/folder1_test/job/job_test", "job_test1")
	//fmt.Printf("%+v", job)
	//taskId, err := ctr.jenkinsSvc.BuildJob(c, "go-test", nil)
	//if err != nil {
	//	zap.L().Error("jenkinsSvc.BuildJob", zap.Error(err))
	//}

	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			//"taskNum": taskList,
			//"status":  status,
			"log": log,
			//"resMap": resMap,
			//"taskId": taskId,
		},
	})
}
```

在 service/目录下创建 jenkins_service.go 文件，该文件仅封装了常用服务：

```go
package service

import (
	"github.com/bndr/gojenkins"
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/meta"
	"io/ioutil"
)

type JenkinsService interface {
	GetQueuePendingTaskList(c *gin.Context) (taskIdList []int64, err error)
	CancelQueueTaskById(c *gin.Context, taskId int64) (bool, error)
	CancelQueueAllTask(c *gin.Context) error
	GetJobLatestStatusAndNextBuildNumByName(c *gin.Context, jobName string) (int, int, error)
	GetAllJobLatestStatus(c *gin.Context) (map[string]int, error)
	GetJobLastBuildLog(c *gin.Context, jobName string) (string, error)
	CreateJobFolder(c *gin.Context, folderName string, parents ...string) error
	CreateJobInFolder(c *gin.Context, configPath string, jobName string, parentIDs ...string) (*gojenkins.Job, error)
	RenameJob(c *gin.Context, oldName string, newName string) *gojenkins.Job
	BuildJob(c *gin.Context, jobName string, params map[string]string) (int64, error)
}

type jenkinsService struct {
	jenkinsClient meta.JenkinsClient
}

func NewJenkinsService(jenkinsClient meta.JenkinsClient) JenkinsService {
	return &jenkinsService{
		jenkinsClient: jenkinsClient,
	}
}

// GetQueuePendingTaskList 获取队列中等待的任务列表
func (svc *jenkinsService) GetQueuePendingTaskList(c *gin.Context) (taskIdList []int64, err error) {
	taskIdList = make([]int64, 0, 100)
	queue, err := svc.jenkinsClient.GetQueue(c)
	if err != nil {
		return taskIdList, err
	}
	for _, task := range queue.Raw.Items {
		taskIdList = append(taskIdList, task.ID)
	}
	return taskIdList, nil
}

// CancelQueueAllTask 停止队列中所有等待任务
func (svc *jenkinsService) CancelQueueAllTask(c *gin.Context) error {
	queue, err := svc.jenkinsClient.GetQueue(c)
	if err != nil {
		return err
	}
	for _, task := range queue.Raw.Items {
		if _, err = queue.CancelTask(c, task.ID); err != nil {
			return err
		}
	}
	return nil
}

// CancelQueueTaskById 根据任务ID停止队列中的等待任务
func (svc *jenkinsService) CancelQueueTaskById(c *gin.Context, taskId int64) (bool, error) {
	queue, err := svc.jenkinsClient.GetQueue(c)
	if err != nil {
		return false, err
	}
	// todo 这里返回结果有问题 - 已反馈给作者
	return queue.CancelTask(c, taskId)
}

// GetAllJobLatestStatus 获取所有job最近编译状态
func (svc *jenkinsService) GetAllJobLatestStatus(c *gin.Context) (map[string]int, error) {
	resMap := make(map[string]int, 0)
	jobList, err := svc.jenkinsClient.GetAllJobs(c)
	if err != nil {
		return nil, err
	}
	for _, job := range jobList {
		resMap[job.GetName()] = getJobStatus(job)
	}
	return resMap, nil
}

// GetJobLatestStatusByName 根据job的名称获取某个job最近一次的编译状态
func (svc *jenkinsService) GetJobLatestStatusAndNextBuildNumByName(c *gin.Context, jobName string) (int, int, error) {
	job, err := svc.jenkinsClient.GetJob(c, jobName)
	if err != nil {
		return 0, 0, err
	}
	return int(job.Raw.NextBuildNumber), getJobStatus(job), nil
}

// getJobStatus 根据job编译详情判断状态
// 1 成功 2 失败 0 未编译或其他
func getJobStatus(job *gojenkins.Job) int {
	if job.Raw.LastBuild.Number == 0 { // 未编译
		return 0
	}
	if job.Raw.LastBuild.Number == job.Raw.LastSuccessfulBuild.Number { // 编译成功
		return 1
	}
	if job.Raw.LastBuild.Number == job.Raw.LastFailedBuild.Number { // 编译失败
		return 2
	}
	return 0
}

// GetJobLastBuildLog 获取指定job最近一次编译日志
func (svc *jenkinsService) GetJobLastBuildLog(c *gin.Context, jobName string) (string, error) {
	job, err := svc.jenkinsClient.GetJob(c, jobName)
	if err != nil {
		return "", err
	}
	build, err := job.GetLastBuild(c)
	if err != nil {
		return "", err
	}
	// 实时日志
	resp, err := build.GetConsoleOutputFromIndex(c, 0)
	if err != nil {
		return "", err
	}
	return resp.Content, nil
}

// CreateJobFolder 创建job的文件夹
// 前提条件：安装Folders Plugin
func (svc *jenkinsService) CreateJobFolder(c *gin.Context, folderName string, parents ...string) error {
	_, err := svc.jenkinsClient.CreateFolder(c, folderName, parents...)
	if err != nil {
		return err
	}
	return nil
}

// CreateJobInFolder 创建一个job工程项目
func (svc *jenkinsService) CreateJobInFolder(c *gin.Context, configPath string, jobName string, parentIDs ...string) (*gojenkins.Job, error) {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	config := string(buf)
	return svc.jenkinsClient.CreateJobInFolder(c, config, jobName, parentIDs...)
}

// RenameJob 重命名job工程项目
// oldName为job的完整路径（/job后面的完整路径），例如：xudttest/job/folder1_test/job/job_test
// newName为新的job名称：例如job_test1
func (svc *jenkinsService) RenameJob(c *gin.Context, oldName string, newName string) *gojenkins.Job {
	return svc.jenkinsClient.RenameJob(c, oldName, newName)
}

// BuildJob 编译job，返回队列id
func (svc *jenkinsService) BuildJob(c *gin.Context, jobName string, params map[string]string) (int64, error) {
	return svc.jenkinsClient.BuildJob(c, jobName, params)
}
```

jenkins_service.go 文件依赖 jenkinsClient，这个在 assembly/common.go 文件中定义：

```go
func NewJenkinsClient() (meta.JenkinsClient, error) {
	var cfg jenkinsClient.JenkinsConfig
	if err := viper.Sub("jenkins").Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return jenkinsClient.NewJenkinsClient(cfg)
}
```

引入公共库 prototype/jenkinsClient/jenkinsClient.go 提供了一系列 jenkins 服务：

```go
package jenkinsClient

import (
	"context"
	"github.com/bndr/gojenkins"
)

type JenkinsClient interface {
	Info(ctx context.Context) (*gojenkins.ExecutorResponse, error)
	SafeRestart(ctx context.Context) error
	GetQueue(ctx context.Context) (*gojenkins.Queue, error)
	GetQueueUrl() string
	GetQueueItem(ctx context.Context, id int64) (*gojenkins.Task, error)
	GetAllJobs(ctx context.Context) ([]*gojenkins.Job, error)
	GetAllJobNames(ctx context.Context) ([]gojenkins.InnerJob, error)
	GetJob(ctx context.Context, id string, parentIDs ...string) (*gojenkins.Job, error)
	GetSubJob(ctx context.Context, parentId string, childId string) (*gojenkins.Job, error)
	CreateJob(ctx context.Context, config string, options ...interface{}) (*gojenkins.Job, error)
	UpdateJob(ctx context.Context, job string, config string) *gojenkins.Job
	RenameJob(ctx context.Context, job string, name string) *gojenkins.Job
	CopyJob(ctx context.Context, copyFrom string, newName string) (*gojenkins.Job, error)
	DeleteJob(ctx context.Context, name string) (bool, error)
	BuildJob(ctx context.Context, name string, params map[string]string) (int64, error)
	GetBuildFromQueueID(ctx context.Context, queueid int64) (*gojenkins.Build, error)
	GetAllNodes(ctx context.Context) ([]*gojenkins.Node, error)
	GetNode(ctx context.Context, name string) (*gojenkins.Node, error)
	CreateNode(ctx context.Context, name string, numExecutors int, description string, remoteFS string, label string, options ...interface{}) (*gojenkins.Node, error)
	DeleteNode(ctx context.Context, name string) (bool, error)
	GetFolder(ctx context.Context, id string, parents ...string) (*gojenkins.Folder, error)
	CreateFolder(ctx context.Context, name string, parents ...string) (*gojenkins.Folder, error)
	CreateJobInFolder(ctx context.Context, config string, jobName string, parentIDs ...string) (*gojenkins.Job, error)
	GetLabel(ctx context.Context, name string) (*gojenkins.Label, error)
	GetAllBuildIds(ctx context.Context, job string) ([]gojenkins.JobBuild, error)
	GetBuild(ctx context.Context, jobName string, number int64) (*gojenkins.Build, error)
	GetArtifactData(ctx context.Context, id string) (*gojenkins.FingerPrintResponse, error)
	GetPlugins(ctx context.Context, depth int) (*gojenkins.Plugins, error)
	UninstallPlugin(ctx context.Context, name string) error
	HasPlugin(ctx context.Context, name string) (*gojenkins.Plugin, error)
	InstallPlugin(ctx context.Context, name string, version string) error
	ValidateFingerPrint(ctx context.Context, id string) (bool, error)
	GetAllViews(ctx context.Context) ([]*gojenkins.View, error)
	GetView(ctx context.Context, name string) (*gojenkins.View, error)
	CreateView(ctx context.Context, name string, viewType string) (*gojenkins.View, error)
	Poll(ctx context.Context) (int, error)
}

type jenkinsClient struct {
	*gojenkins.Jenkins
}

// JenkinsConfig Jenkins连接配置
type JenkinsConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

func NewJenkinsClient(cfg JenkinsConfig) (JenkinsClient, error) {
	ctx := context.Background()
	client, err := gojenkins.CreateJenkins(nil, cfg.Host, cfg.Username, cfg.Password).Init(ctx)
	if err != nil {
		return nil, err
	}
	return &jenkinsClient{
		Jenkins: client,
	}, nil
}
```

## 阿里云短信服务

使用阿里云短信服务接口，需要在阿里云短信服务商购买并获取到相应的配置。

### 使用

在 internal/assembly/common.go 中添加方法获取到阿里云短信客户端

```go
// NewAlibabaSmsClient alibaba短信连接
func NewAlibabaSmsClient() (meta.AlibabaSmsClient, error) {
	var cfg alibabaSmsClient.ClientConfig
	if err := viper.Sub("alibaba_sms").Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return alibabaSmsClient.CreateClient(cfg)
}
```

在 internal/service 下新增 sms_service.go 文件，实现发送短信验证码的功能，代码如下：

```go
// SendCode 发送短信验证码
func (svc *smsService) SendCode(ctx *gin.Context, phone string) (code string, err error) {
	var params alibabaSmsClient.SendSmsParams
	if err := viper.Sub("alibaba_sms_code").Unmarshal(&params); err != nil {
		return "", err
	}
	params.Phone = phone
	code = tools.GetRandomString(6, 1)
	templateParam := struct {
		Code string `json:"code"`
	}{
		Code: code,
	}
	templateParamByte, err := json.Marshal(templateParam)
	if err != nil {
		return "", err
	}
	params.TemplateParam = string(templateParamByte)
	err = svc.smsClient.SendCode(params)
	if err != nil {
		zap.L().Error("smsService.SendCode err", zap.Error(err), zap.String("phone", params.Phone))
		return "", err
	}
	return code, nil
}
```

控制器层可直接调用，实现短信发送。

### 短信客户端实现

```go
package alibabaSmsClient

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/pkg/errors"
)

type AlibabaSmsClient interface {
	SendCode(params SendSmsParams) (err error)
}

type alibabaSmsClient struct {
	Client *dysmsapi20170525.Client
}

// ClientConfig 连接客户端配置
type ClientConfig struct {
	Endpoint        string `json:"endpoint" mapstructure:"endpoint"`
	AccessKeyId     string `json:"access_key_id" mapstructure:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" mapstructure:"access_key_secret"`
}

// SendSmsParams 发送短信参数结构体
type SendSmsParams struct {
	Phone         string `json:"phone" mapstructure:"phone"`
	SignName      string `json:"sign_name" mapstructure:"sign_name"`
	TemplateCode  string `json:"template_code" mapstructure:"template_code"`
	TemplateParam string `json:"template_param" mapstructure:"template_param"`
}

// CreateClient 使用AK&SK初始化账号Client
func CreateClient(cfg ClientConfig) (AlibabaSmsClient, error) {
	config := &openapi.Config{
		// 访问的域名
		Endpoint: &cfg.Endpoint,
		// 您的 AccessKey ID
		AccessKeyId: &cfg.AccessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: &cfg.AccessKeySecret,
	}
	client, err := dysmsapi20170525.NewClient(config)
	return &alibabaSmsClient{
		Client: client,
	}, err
}

// SendCode 发送短信验证码
func (cli *alibabaSmsClient) SendCode(params SendSmsParams) (err error) {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(params.Phone),
		SignName:      tea.String(params.SignName),
		TemplateCode:  tea.String(params.TemplateCode),
		TemplateParam: tea.String(params.TemplateParam),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (err error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				err = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, err = cli.Client.SendSmsWithOptions(sendSmsRequest, runtime)
		if err != nil {
			return err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		errMsg := util.AssertAsString(error.Message)
		return errors.New(*errMsg)
	}
	return err
}
```

## CRUD 代码自动生成

### Makefile

编写 Crud_Makefile 文件，其中相关变量需要根据项目实际情况进行修改，目前仅支持单个单词的控制器，并要求表名遵循相应规则，若有更高需求，可生成后再做相应修改，代码如下：

```makefile
goModule = github.com/go_example
name = test
tableName = test
tableModel = Test
# 数据库连接
db = "root:root@tcp(127.0.0.1:3306)/go_example?charset=utf8mb4"
# 模板文件目录
templatePath = "./assets/templates/crud"
# 对name进行处理为首字母大写和全小写
nameLitter = `echo $(name) | awk '{print tolower($0)}'`
nameUpper = `echo $(name) | awk '{print substr(toupper($0),1,1)substr(tolower($0),2)}'`
# 目标文件目录
targetControllerPath = "./internal/controller/$(nameLitter)_controller.go"
targetServicePath = "./internal/service/$(nameLitter)_service.go"
targetRepositoryPath = "./internal/repository/mysql/$(nameLitter)_repository.go"
targetModelPath = "./internal/model/common.go"
targetControllerWirePath = "./internal/assembly/controller.go"
targetServiceWirePath = "./internal/assembly/service.go"
cleanTemplatePath = "./assets/templates/crud/clean.template"

.PHONY: crud
all: controller service repository model model_db wire wire_tmp
	@echo "请手动添加$(templatePath)/wire/http.tmp文件内容到以下三个文件："
	@echo "./internal/assembly/server.go"
	@echo "./internal/server/http.go"
	@echo "手动添加完代码后，请切换到项目internal/assembly/目录，执行wire命令"

.PHONY: controller
controller:
	@# 生成controller
	@# 复制文件到指定位置，并改名
	@cp $(templatePath)/controller.template $(targetControllerPath)
	@# 替换文件中的宏
	@sed -i '' 's#{{PROJECT_MODULE}}#$(goModule)#g' $(targetControllerPath)
	@sed -i '' 's#{{CONTROLLER_NAME_LITTER}}#'$(nameLitter)#g $(targetControllerPath)
	@sed -i '' 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(targetControllerPath)
	@sed -i '' 's#{{TABLE_MODEL}}#'$(tableModel)#g $(targetControllerPath)
	@echo "generate controller success"

.PHONY: service
service:
	@# 生成service
	@# 复制文件到指定位置，并改名
	@cp $(templatePath)/service.template $(targetServicePath)
	@# 替换文件中的宏
	@sed -i '' 's#{{PROJECT_MODULE}}#$(goModule)#g' $(targetServicePath)
	@sed -i '' 's#{{CONTROLLER_NAME_LITTER}}#'$(nameLitter)#g $(targetServicePath)
	@sed -i '' 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(targetServicePath)
	@sed -i '' 's#{{TABLE_MODEL}}#'$(tableModel)#g $(targetServicePath)
	@echo "generate service success"

.PHONY: repository
repository:
	@# 生成repository
	@# 复制文件到指定位置，并改名
	@cp $(templatePath)/repository.template $(targetRepositoryPath)
	@# 替换文件中的宏
	@sed -i '' 's#{{PROJECT_MODULE}}#$(goModule)#g' $(targetRepositoryPath)
	@sed -i '' 's#{{CONTROLLER_NAME_LITTER}}#'$(nameLitter)#g $(targetRepositoryPath)
	@sed -i '' 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(targetRepositoryPath)
	@sed -i '' 's#{{TABLE_NAME}}#'$(tableName)#g $(targetRepositoryPath)
	@sed -i '' 's#{{TABLE_MODEL}}#'$(tableModel)#g $(targetRepositoryPath)
	@echo "generate repository success"

.PHONY: model
model:
	@# 生成请求model
	@sed 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/model.template >> $(targetModelPath)
	@echo "generate model success"

.PHONY: model_db
model_db:
	@# 生成数据库model
	@xorm reverse mysql $(db) ./internal/model/mysql/templates/goxorm ./internal/model/mysql $(tableName)
	@echo "generate model_db success"

.PHONY: wire
wire:
	@# 生成wire依赖文件
	@sed 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/wire/controller.template >> $(targetControllerWirePath)
	@sed 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/wire/service.template >> $(targetServiceWirePath)
	@echo "generate wire success"

.PHONY: wire_tmp
wire_tmp:
	@# 请手动复制以下http.tmp文件内容到对应文件下
	@cp $(templatePath)/wire/http.template $(templatePath)/wire/http.tmp
	@sed -i '' 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/wire/http.tmp
	@sed -i '' 's#{{CONTROLLER_NAME_LITTER}}#'$(nameLitter)#g $(templatePath)/wire/http.tmp
	@echo "generate wire_tmp success"

.PHONY: clean
clean:
	@# 删除controller
	@rm -f $(targetControllerPath)
	@echo "delete controller success"
	@# 删除service
	@rm -f $(targetServicePath)
	@echo "delete service success"
	@# 删除repository
	@rm -f $(targetRepositoryPath)
	@echo "delete repository success"
	@# 删除数据库表模型
	@rm ./internal/model/mysql/$(tableName).go
	@echo "delete db_model success"
	@# 请手动删除./assets/templates/crud/clean.template对应的内容
	@rm -f $(cleanTemplatePath)
	@sed 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/model.template >> $(cleanTemplatePath)
	@sed 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/wire/controller.template >> $(cleanTemplatePath)
	@sed 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/wire/service.template >> $(cleanTemplatePath)
	@echo "请手动删除项目中./assets/templates/crud/clean.template文件内容对应以下文件："
	@echo "./internal/model/common.go"
	@echo "./internal/assembly/controller.go"
	@echo "./internal/assembly/service.go"
	@echo "若手动添加过其他代码，请删除，例如以下两个文件："
	@echo "./internal/assembly/server.go"
	@echo "./internal/server/http.go"
```

### Makefile 执行命令

生成 CRUD 代码，根据输出提示进行操作。

```shell
# name：控制器的名称，一般为业务名，例如：user、department等；
# tableName：对应数据库表名；
# tableModel：数据库表生成表模型名，需要对应tableName，例如tableName为user_data，那么tableModel为UserData；
# -f Crud_Makefile：表示指定Crud_Makefile文件；
make all name=user tableName=user_data tableModel=UserData -f Crud_Makefile
```

删除 CRUD 代码，根据输出提示进行操作。

```shell
make clean name=user tableName=user_data tableModel=UserData -f Crud_Makefile
```

## Nacos 配置中心

Nacos 是阿里巴巴推出来的一个新开源项目。一个更易于构建云原生应用的动态服务发现、配置管理和服务管理平台。

### 配置文件修改

configs/debug 新增 config_center.yaml 配置文件，定义了连接 nacos 配置中心相关的配置信息：

```yaml
config_center:
  server_ip_list:
    - "127.0.0.1"
  server_http_port: 8848
  server_grpc_port: 9848
  namespace_id: "debug"
  timeout: 5000
  log_level: "debug"
  log_path: "./nacos/logs"
  cache_path: "./nacos/cache"
  group: "go_example"
  data_id: "debug-go_example-common.yaml"
```

### root.go 初始化项目配置

支持本地调试可选择本地配置文件或配置中心获取项目配置，生产及测试环境使用配置中心配置。将配置信息解析到 meta.GlobalConfig 全局变量，并实现客户端监听，当配置发生变化，客户端通过 OnChange 回调修改 meta.GlobalConfig 全局变量。

```go
func initConfig() {
	// 从系统环境变量中读取运行环境
	meta.EnvMode = env.Mode()
	if env.IsDebugging() && !meta.DebugUseConfigCenter { // 本地调试若不使用配置中心则直接读取common.yaml文件
		viper.SetConfigFile(meta.CfgFile + "/" + meta.EnvMode + "/common.yaml")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalln(err)
		}
		if err := viper.Unmarshal(&meta.GlobalConfig); err != nil {
			log.Fatalln(err)
		}
	} else { // 测试环境、生产环境从配置中心读取
		viper.SetConfigFile(meta.CfgFile + "/" + meta.EnvMode + "/config_center.yaml")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalln(err)
		}

		// 从配置中心读取项目配置
		var cfg nacosClient.Config
		if err := viper.Sub("config_center").Unmarshal(&cfg); err != nil {
			log.Fatalln(err)
		}
		configCenterClient, clean, err := nacosClient.NewClient(cfg)
		if err != nil {
			log.Fatalln(err)
		}
		defer clean()
		globalConfigStr, err := configCenterClient.GetConfig(vo.ConfigParam{
			DataId: cfg.DataId,
			Group:  cfg.Group,
		})
		if err != nil {
			log.Fatalln(err)
		}
		// 将读取到的配置信息转为全局配置
		setGlobalConfigFromData(globalConfigStr)

		// 监听实现自动感知
		configCenterClient.ListenConfig(vo.ConfigParam{
			DataId: cfg.DataId,
			Group:  cfg.Group,
			OnChange: func(namespace, group, dataId, data string) {
				setGlobalConfigFromData(data)
			},
		})
	}
}

// 从文本读取到全局配置
func setGlobalConfigFromData(data string) {
    // 将读取到的配置信息转为全局配置
    viper.SetConfigType("yaml")
    err := viper.ReadConfig(bytes.NewBuffer([]byte(data)))
    if err != nil {
        log.Fatalln(err)
    }
    if err = viper.Unmarshal(&meta.GlobalConfig); err != nil {
        log.Fatalln(err)
    }
}
```

### 从全局变量中读取配置信息

新增全局变量结构体，对应文件为：internal/model/global_config.go，并在 meta.go 中定义全局变量。

```go
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

```

将从 viper 全局变量中读取改为从 meta.GlobalConfig 全局变量。如下示例：

```go
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
```

### Nacos 配置中心服务

```go
package nacosClient

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type ConfigCenterClient interface {
	config_client.IConfigClient
}

type configCenterClient struct {
	config_client.IConfigClient
}

// Config Nacos连接配置
type Config struct {
	ServerIpList   []string `json:"server_ip_list" mapstructure:"server_ip_list"`
	ServerHttpPort uint64   `json:"server_http_port" mapstructure:"server_http_port"`
	ServerGrpcPort uint64   `json:"server_grpc_port" mapstructure:"server_grpc_port"`
	NamespaceId    string   `json:"namespace_id" mapstructure:"namespace_id"`
	Timeout        uint64   `json:"timeout" mapstructure:"timeout"`
	LogLevel       string   `json:"log_level" mapstructure:"log_level"`
	LogPath        string   `json:"log_path" mapstructure:"log_path"`
	CachePath      string   `json:"cache_path" mapstructure:"cache_path"`
	Group          string   `json:"group" mapstructure:"group"`
	DataId         string   `json:"data_id" mapstructure:"data_id"`
}

func NewClient(cfg Config) (ConfigCenterClient, func(), error) {
	// Nacos服务端配置
	serverConfigList := getServerConfig(cfg)
	// Nacos客户端配置
	clientConfig := getClientConfig(cfg)

	client, err := clients.NewConfigClient(vo.NacosClientParam{
		ServerConfigs: serverConfigList,
		ClientConfig:  &clientConfig,
	})

	if err != nil {
		panic(err)
	}

	return &configCenterClient{client}, func() {
		client.CloseClient()
	}, nil
}

// 获取服务端配置
func getServerConfig(cfg Config) []constant.ServerConfig {
	serverConfigList := make([]constant.ServerConfig, 0, 10)
	for _, ip := range cfg.ServerIpList {
		serverConfigList = append(serverConfigList, *constant.NewServerConfig(ip, cfg.ServerHttpPort, constant.WithGrpcPort(cfg.ServerGrpcPort)))
	}
	return serverConfigList
}

// 获取客户端配置
func getClientConfig(cfg Config) constant.ClientConfig {
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(cfg.NamespaceId),
		constant.WithTimeoutMs(cfg.Timeout),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(cfg.LogPath),
		constant.WithCacheDir(cfg.CachePath),
		constant.WithLogLevel(cfg.LogLevel),
	)
	return clientConfig
}
```

## Nacos 服务注册与服务发现

在 Nacos 配置中心的基础上，对 Nacos 提供的服务进行拆分为两个服务，配置中心与服务注册：

- 公共配置对应 github.com/ken-house/go-contrib/prototype/nacosClient/common.go
- 配置中心对应 github.com/ken-house/go-contrib/prototype/nacosClient/configCenterClient.go
- 服务注册对应 github.com/ken-house/go-contrib/prototype/nacosClient/serviceClient.go

### 服务注册与服务发现

获得服务注册与发现的客户端连接，代码如下：

```go
package nacosClient

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type ServiceClient interface {
	naming_client.INamingClient
	FindHealthInstanceAddress(clusters []string, serviceName string, groupName string) (string, error)
}

type serviceClient struct {
	naming_client.INamingClient
}

func NewServiceClient(cfg Config) (ServiceClient, func(), error) {
	// Nacos服务端配置
	serverConfigList := getServerConfig(cfg)
	// Nacos客户端配置
	clientConfig := getClientConfig(cfg)

	client, err := clients.NewNamingClient(vo.NacosClientParam{
		ServerConfigs: serverConfigList,
		ClientConfig:  &clientConfig,
	})

	if err != nil {
		panic(err)
	}

	return &serviceClient{client}, func() {
		client.CloseClient()
	}, nil
}

// FindHealthInstanceAddress 根据serviceId查找到一个健康的服务实例，获取其地址
func (cli *serviceClient) FindHealthInstanceAddress(clusters []string, serviceName string, groupName string) (string, error) {
	serviceInfo, err := cli.INamingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		Clusters:    clusters,
		ServiceName: serviceName,
		GroupName:   groupName,
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d", serviceInfo.Ip, serviceInfo.Port), nil
}
```

### 服务注册

在 cmd/grpc.go 文件中调用 RegisterNacos()方法，将服务注册到 nacos。

```go
func (srv *grpcServer) RegisterNacos() (nacosClient.ServiceClient, []vo.RegisterInstanceParam, error) {
	ip := tools.GetOutBoundIp()
	if ip == "" {
		return nil, nil, errors.New("GetOutBoundIp fail")
	}
	port := meta.GlobalConfig.Server.Grpc.Port
	serviceArr := make([]vo.RegisterInstanceParam, 0, 10)
	for serviceName, _ := range serviceNameArr {
		param := vo.RegisterInstanceParam{
			Ip:          ip,
			Port:        cast.ToUint64(port),
			Weight:      10,
			Enable:      true,
			Healthy:     true,
			Metadata:    map[string]string{"appname": "go_example"},
			ServiceName: serviceName,
			GroupName:   "go_example",
			Ephemeral:   true,
		}
		_, err := srv.nacosServiceClient.RegisterInstance(param)
		if err != nil {
			return nil, nil, err
		}
		serviceArr = append(serviceArr, param)
	}
	return srv.nacosServiceClient, serviceArr, nil
}
```

### 服务发现

通过健康检查获取到一个健康的 Nacos 服务节点，连接 Nacos 创建 grpc 客户端。

```go
	// nacos 服务发现
	nacosAddr, err := ctr.nacosServiceClient.FindHealthInstanceAddress(nil, "hello", "go_example")
	if err != nil {
		zap.L().Panic("nacosServiceClient.FindHealthInstanceAddress err", zap.Error(err))
	}
	conn, err := grpc.Dial(nacosAddr, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"HealthCheckConfig": {"ServiceName": "%s"}}`, meta.HEALTHCHECK_SERVICE)), grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(grpcAuth), grpc.WithUnaryInterceptor(UnaryClientInterceptor()))
	if err != nil {
		log.Printf("Dial err:%+v", err)
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_DIAL.ToastError())
	}
```

## Kafka 服务

Kafka 是一个开源的分布式事件流平台，在大数据场景主要采用 Kafka 作为消息队列。

### 安装

```shell
go get github.com/Shopify/sarama
```

### 生产者

kafka 生产者分为同步和异步，同步需要等待 broker 节点根据 acks 响应所有步骤完成返回；异步则将消息推送到队列后即返回。

common.go 文件整理了 kafka 相关配置及公共方法。

```go
package kafkaClient

import "github.com/Shopify/sarama"

// Config kafka连接及配置信息
type Config struct {
	ServerAddrList []string       `json:"server_addr_list" mapstructure:"server_addr_list"` // kafka地址
	ProducerConfig ProducerConfig `json:"producer_config"  mapstructure:"producer_config"`  // 生产者配置
	ConsumerConfig ConsumerConfig `json:"consumer_config" mapstructure:"consumer_config"`   // 消费者配置
}

// ProducerConfig kafka生产者配置参数
type ProducerConfig struct {
	Ack                   int `json:"acks" mapstructure:"acks"`                                       // 应答类型 0  1 -1
	PartitionerType       int `json:"partitioner_type" mapstructure:"partitioner_type"`               // 分区算法
	FlushMessageNum       int `json:"flush_message_num" mapstructure:"flush_message_num"`             // 达到多少条消息才发送
	FlushMessageFrequency int `json:"flush_message_frequency" mapstructure:"flush_message_frequency"` // 达到多少秒消息才发送
}

// ConsumerConfig kafka消费者配置参数
type ConsumerConfig struct {
}

// 指定分区算法
func setPartition(config *sarama.Config, partitionerType int) {
	switch partitionerType {
	case 1: // 随机算法
		config.Producer.Partitioner = sarama.NewRandomPartitioner
	case 2: // robin算法
		config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	case 3: // 按消息内容计算分区
		config.Producer.Partitioner = sarama.NewManualPartitioner
	default: // 按key计算分区
		config.Producer.Partitioner = sarama.NewHashPartitioner
	}
}
```

#### 同步生产者

默认提供了发送单个消息和多个消息的方法，这里对这两个方法进行封装。

```go
package kafkaClient

import (
	"time"

	"github.com/Shopify/sarama"
)

type ProducerSyncClient interface {
	sarama.SyncProducer
	SendOne(topic string, key string, message string, partition int32) (int32, int64, error)
	SendMany(topic string, key string, messageList []string, partition int32) error
}

type producerSyncClient struct {
	sarama.SyncProducer
}

// NewProducerSyncClient 同步生产者
func NewProducerSyncClient(cfg Config) (ProducerSyncClient, func(), error) {
	config := sarama.NewConfig()
	// 指定应答方式
	config.Producer.RequiredAcks = sarama.RequiredAcks(cfg.ProducerConfig.Ack)
	// 设置达到多少条消息才发送到kafka
	config.Producer.Flush.Messages = cfg.ProducerConfig.FlushMessageNum
	// 设置间隔多少秒才发送到kafka
	if cfg.ProducerConfig.FlushMessageFrequency > 0 {
		config.Producer.Flush.Frequency = time.Duration(cfg.ProducerConfig.FlushMessageFrequency) * time.Second
	}
	// 成功交付的消息将在success channel返回 必须指定为true
	config.Producer.Return.Successes = true
	// 指定分区算法
	setPartition(config, cfg.ProducerConfig.PartitionerType)
	// 建立同步生产者连接
	productClient, err := sarama.NewSyncProducer(cfg.ServerAddrList, config)
	if err != nil {
		return nil, nil, err
	}

	return &producerSyncClient{productClient}, func() {
		defer productClient.Close()
	}, nil
}

// SendOne 单条消息发送
func (cli *producerSyncClient) SendOne(topic string, key string, message string, partition int32) (int32, int64, error) {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(message),
		Partition: partition,
	}
	if key != "" {
		msg.Key = sarama.StringEncoder(key)
	}
	return cli.SendMessage(msg)
}

// SendMany 多条消息发送
func (cli *producerSyncClient) SendMany(topic string, key string, messageList []string, partition int32) error {
	msgList := make([]*sarama.ProducerMessage, 0, 100)
	for _, message := range messageList {
		msg := &sarama.ProducerMessage{
			Topic:     topic,
			Value:     sarama.StringEncoder(message),
			Partition: partition,
		}
		if key != "" {
			msg.Key = sarama.StringEncoder(key)
		}
		msgList = append(msgList, msg)
	}
	return cli.SendMessages(msgList)
}
```

在 controller 层，通过如下代码进行调用：

```go
// ProducerSync 同步发送
func (ctr kafkaController) ProducerSync(ctx *gin.Context) (int, gin.Negotiate) {
	message := "hello one message"
	msg := &sarama.ProducerMessage{
		Topic: "first",
		Value: sarama.StringEncoder(message),
	}
	// 单条消息发送
	partition, offset, err := ctr.kafkaProducerSyncClient.SendMessage(msg)
	fmt.Println(partition, offset)
	// 单条消息发送
	partition, offset, err = ctr.kafkaProducerSyncClient.SendOne("first", "", message, 0)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_SYSTEM.ToastError())
	}
	fmt.Println(partition, offset)
	messageList := []string{
		"hello world sync",
		"hi",
	}
	// 多条消息发送
	err = ctr.kafkaProducerSyncClient.SendMany("first", "", messageList, 0)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_SYSTEM.ToastError())
	}

	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"message": "OK",
		},
	})
}
```

#### 异步生产者

异步生产者为单独的对象连接，并封装了发送消息方法。

```go
package kafkaClient

import (
	"time"

	"github.com/Shopify/sarama"
)

type ProducerAsyncClient interface {
	sarama.AsyncProducer
	SendOne(topic string, key string, message string, partition int32)
}

type producerAsyncClient struct {
	sarama.AsyncProducer
}

// NewProducerAsyncClient 异步生产者
func NewProducerAsyncClient(cfg Config) (ProducerAsyncClient, func(), error) {
	config := sarama.NewConfig()
	// 指定应答方式
	config.Producer.RequiredAcks = sarama.RequiredAcks(cfg.ProducerConfig.Ack)
	// 设置达到多少条消息才发送到kafka
	config.Producer.Flush.Messages = cfg.ProducerConfig.FlushMessageNum
	// 设置间隔多少秒才发送到kafka
	if cfg.ProducerConfig.FlushMessageFrequency > 0 {
		config.Producer.Flush.Frequency = time.Duration(cfg.ProducerConfig.FlushMessageFrequency) * time.Second
	}
	// 成功交付的消息将在success channel返回 必须指定为true
	config.Producer.Return.Successes = true
	// 指定分区算法
	setPartition(config, cfg.ProducerConfig.PartitionerType)
	// 建立同步生产者连接
	productClient, err := sarama.NewAsyncProducer(cfg.ServerAddrList, config)
	if err != nil {
		return nil, nil, err
	}

	return &producerAsyncClient{productClient}, func() {
		defer productClient.Close()
	}, nil
}

// SendOne 单条消息发送
func (cli *producerAsyncClient) SendOne(topic string, key string, message string, partition int32) {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(message),
		Partition: partition,
	}
	if key != "" {
		msg.Key = sarama.StringEncoder(key)
	}
	cli.Input() <- msg
}
```

在 controller 层调用：

```go
// ProducerAsync 异步发送
func (ctr kafkaController) ProducerAsync(ctx *gin.Context) (int, gin.Negotiate) {
	message := "hello world async cc"
	ctr.kafkaProducerAsyncClient.SendOne("first", "", message, 0)
	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"message": "OK",
		},
	})
}
```

### 消费

kafka 提供消费者，这里封装消费一个 topic 的方法：

```go
package kafkaClient

import (
	"sync"

	"github.com/Shopify/sarama"
)

type ConsumerClient interface {
	sarama.Consumer
	ConsumeTopic(topic string, isNew int64, ConsumerFunc func(message *sarama.ConsumerMessage)) error
}

type consumerClient struct {
	sarama.Consumer
}

func NewConsumerClient(cfg Config) (ConsumerClient, func(), error) {
	config := sarama.NewConfig()
	client, err := sarama.NewConsumer(cfg.ServerAddrList, config)
	if err != nil {
		return nil, nil, err
	}

	return &consumerClient{client}, func() {
		defer client.Close()
	}, nil
}

// ConsumeTopic 消费整个Topic
func (cli *consumerClient) ConsumeTopic(topic string, isNew int64, ConsumerFunc func(message *sarama.ConsumerMessage)) error {
	partitionList, err := cli.Partitions(topic) // 通过topic获取到所有的分区
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	for _, partition := range partitionList {
		partitionConsumer, err := cli.ConsumePartition(topic, partition, isNew)
		if err != nil {
			return err
		}
		defer partitionConsumer.AsyncClose()

		wg.Add(1)
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range partitionConsumer.Messages() {
				defer wg.Done()
				ConsumerFunc(msg)
			}
		}(partitionConsumer)
	}
	wg.Wait()
	return nil
}
```

在 controller 层调用：

```go
// Consumer 启动消费
func (ctr kafkaController) Consumer(ctx *gin.Context) (int, gin.Negotiate) {
	var ConsumerFunc = func(msg *sarama.ConsumerMessage) {
		fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
	err := ctr.kafkaConsumerClient.ConsumeTopic("first", sarama.OffsetOldest, ConsumerFunc)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_SYSTEM.ToastError())
	}
	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"message": "OK",
		},
	})
}
```

### Docker 容器化部署服务

#### 编写 Dockerfile 文件

在项目目录下新建 Dockerfile 文件，文件内容如下：

```dockerfile
FROM golang:1.18-alpine3.16

# 维护者
MAINTAINER Ken

# 设置环境变量
ENV GO111MODULE=on CGO_ENABLE=0 GOOS=linux GO_ARCH=amd64 GOPROXY=https://goproxy.cn,direct

# 将本地项目文件拷贝到容器当前目录
COPY . /build

# 切换容器的目录到
WORKDIR /build

# 执行go build生成可执行文件example-http-server
RUN go build -o example-http-server main.go


# 切换到容器下的目标目录
WORKDIR /dist

# 将需要加载的本地配置文件copy到容器下的/dist目录
COPY ./configs ./configs
COPY ./assets  ./assets
COPY ./views   ./views

# 向容器添加卷
VOLUME ["/dist/configs","/dist/nacos","/dist/logs"]

# 将容器内的可执行文件拷贝到容器的当前目录
RUN cp /build/example-http-server .

# 暴露端口
EXPOSE 6060
EXPOSE 8080


# 启动容器时运行项目，并指定参数http
ENTRYPOINT ["/dist/example-http-server","http"]

```

#### 构建镜像

生成的镜像会放在 docker 本地镜像仓库

```shell
docker build . -t example-http-server:latest
```

#### 运行容器

需要在挂在的配置目录下，创建配置文件，配置文件中的 IP 地址要改为可访问地址。

```dockerfile
docker run -d -p 8888:8080 -p 6666:6060 -v ~/dockerVolumes/exampleVolume/http/logs:/dist/logs -v ~/dockerVolumes/exampleVolume/http/nacos:/dist/nacos -v ~/dockerVolumes/exampleVolume/http/configs:/dist/configs --name example-http-server example-http-server
```

打开浏览器，访问：http://127.0.0.1:8888/hello接口访问。

### K8s 容器编排部署

#### 推送镜像到 Docker Hub

在上一节，我们将项目打包成镜像，并可启动容器运行。接下来，我们将镜像推送到 Docker Hub。

```shell
# 对镜像打标签
docker tag example-http-server xudengtang/example-http-server:1.0

# 登陆Docker Hub，若未注册，需先注册。
docker login

# 推送镜像
docker push xudengtang/example-http-server:1.0
```

这样镜像就推送到 Docker Hub，并且是 Public。

#### Pod 部署

这里通过 Deployment 控制器对 Pod 进行创建。在项目目录下创建/deploy/debug/example-http-server-deployment.yaml,代码如下：

```yaml
apiVersion: apps/v1
kind: Deployment # 类型为deployment
metadata: # deployment的详细信息
  name: example-http-server-deploy
  namespace: dev
spec: # deployment描述
  replicas: 3 # 每个pod生成3个副本
  selector: # 对标签为app=example-http-server-pod的pod进行操作
    matchLabels:
      app: example-http-server-pod
  template: # pod的模板设置
    metadata: # pod的详细信息
      labels: # 给pod打上标签app=example-http-server-pod
        app: example-http-server-pod
    spec: # pod描述
      containers: # 容器信息
        - name: example-http-server
          image: xudengtang/example-http-server:1.0
          volumeMounts: # 挂载目录
            - name: configs-volume
              mountPath: /dist/configs
            - name: logs-volume
              mountPath: /dist/logs
            - name: nacos-volume
              mountPath: /dist/nacos
      volumes: # 定义存储卷
        - name: configs-volume
          hostPath:
            path: /Users/zonst/dockerVolumes/exampleVolume/http/configs
        - name: logs-volume
          hostPath:
            path: /Users/zonst/dockerVolumes/exampleVolume/http/logs
        - name: nacos-volume
          hostPath:
            path: /Users/zonst/dockerVolumes/exampleVolume/http/nacos
```

执行创建命令：

```shell
kubectl apply -f ./deploy/debug/example-http-server-deployment.yaml
# 查看pod列表
kubectl get pods -n dev
```

到这里，项目在 Pod 下的容器中运行成功，但外部还无法访问。

### Service 服务

K8s 提供 Service 服务对 Pod 端口进行代理，这里使用 NodePort，以便外部可以通过节点 IP 进行访问；创建 example-http-server-service.yaml，代码如下：

```yaml
apiVersion: v1
kind: Service # 类型为service
metadata: # service详细信息
  name: example-http-server-service
  namespace: dev
spec: # service描述
  selector: # 对标签为app=example-http-server-pod的pod进行代理
    app: example-http-server-pod
  type: NodePort # 指定service类型为NodePort
  ports: # 端口映射
    - port: 8080 # 端口名称
      nodePort: 30000 # service的NodePort端口（范围为30000-32767）
      targetPort: 8080 # pod节点暴露的端口
```

执行如下命令创建 Service:

```shell
kubectl apply -f ./deploy/debug/example-http-server-service.yaml
```

此时，可以通过 127.0.0.1:30000/hello 访问服务；

### K8s 容器化部署全面升级

这里以 go_example 项目为例。由于 go_example 项目使用 cobra 做项目脚手架，支持一个项目多个服务。在前面构建镜像时，仅实现了一个 http 服务，因此需要对代码进行如下调整。

### Makefile 文件

```Makefile
.PHONY: build
build:
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -ldflags "-s -w" -o ./bin/go-example main.go

.PHONY: push
push:
  docker build . -t xudengtang/go-example:latest
  docker push xudengtang/go-example:latest

.PHONY: clean
clean:
  rm -f ./bin/go-example
  docker image rm xudengtang/go-example:latest
```

本地执行命令生成可执行文件到项目目录下的 bin/go_example

```Bash
make build
```

### Dockerfile 文件

上面通过 make build 已经生成了可执行文件，修改 Dockerfile 文件，基于 alpine 基础镜像（最小化的 Liunx 镜像），将可执行文件 COPY 到镜像中，同时加载一些所需静态文件，添加卷等。

```Docker
FROM alpine:latest

# 维护者
MAINTAINER Ken

# 切换到容器下的目标目录
WORKDIR /

# 将本地编译好的可执行文件复制到容器中
COPY ./bin/go-example go-example

# 将需要加载的本地配置文件copy到容器下的/dist目录
COPY ./views   ./views
COPY ./assets/jenkins ./assets/jenkins

# 向容器添加卷
VOLUME ["/logs", "/nacos"]

# 容器运行时执行命令
ENTRYPOINT ["/go-example"]
```

本地执行如下命令，将打包镜像并推送到 Docker Hub。

```Bash
make push
```

### ConfigMap 创建配置

创建 example-server-configmap.yaml 文件将原 configs/debug/common.yaml 及 config_center.yaml 配置文件内容，通过 k8s 提供的 ConfigMap 配置存储到容器中。

```YAML
apiVersion: v1
kind: ConfigMap
metadata:
  name: example-server-configmap
  namespace: dev
data:
  config_center.yaml: |-
    config_center:
      server_ip_list:
        - "10.0.98.64"
      server_http_port: 8848
      server_grpc_port: 9848
      namespace_id: "debug"
      timeout: 5000
      log_level: "debug"
      log_path: "./nacos/logs"
      cache_path: "./nacos/cache"
      group: "go_example"
      data_id: "debug-go_example-common.yaml"
  common.yaml: |-
    server:
      http:
        name: "http"
        addr: ""
        port: "8080"
      http_pprof:
        name: "http_pprof"
        addr: ""
        port: "6060"
      socket:
        name: "socket"
        addr: ""
        port: "30000"
      grpc:
        name: "grpc"
        addr: ""
        port: "9090"

    certs:
      cur_key: 1652421664
      keys: [1652421664]

    mysql:
      group:
        max_open: 20
        max_idle: 5
        max_lifetime: 10
        master:
          dsn: "root:root@tcp(10.0.98.64:3306)/go_example?charset=utf8mb4"
        slaves:
          - dsn: "root:root@tcp(10.0.98.64:3306)/go_example?charset=utf8mb4"
      single:
        dsn: "root:root@tcp(10.0.98.64:3306)/go_example?charset=utf8mb4"
        max_open: 20
        max_idle: 5
        max_lifetime: 10
    redis:
      single:
        addr: "10.0.98.64:6379"
        password: "redis"
        db: 0
        pool_size: 20
      group:
        addrs:
          - 10.0.98.64:7000
          - 10.0.98.64:7001
          - 10.0.98.64:7002
          - 10.0.98.64:7003
          - 10.0.98.64:7004
          - 10.0.98.64:7005
        password: ""
        pool_size: 20
    mongodb:
      # 格式：mongodb://username:password@addr1:port2,addr2:port2
      addr: "mongodb://10.0.98.64:27017"
      max_open: 20
    consul:
      host: "10.0.98.64"
      port: "8500"
    jenkins:
      host: "http://10.0.98.64:8088/"
      username: "admin"
      password: "11a8c7f5ad7e4d9bcbbe08eceb415eb7ad"
    alibaba_sms:
      end_point: "dysmsapi.aliyuncs.com"
      access_key_id: "LTAI4GFM5Sqw5AiWwAPq4che"
      access_key_secret: "QfaWjC3BJkOJRP0v3ewntdOQi8nKVZ"
    alibaba_sms_code:
      template_code: "SMS_204125567"
      sign_name: "小区闲置物品信息交流平台" # 短信签名
    kafka:
      server_addr_list:
        - "10.0.98.64:9092"
        - "10.0.98.64:9093"
        - "10.0.98.64:9094"
      producer_config:
        ack: -1 # 0 不需要应答 1 leader应答 -1 leader和follower都需要应答
        partitioner_policy: # "hash" # hash、random、robin、manual；若为consum需要完善go-contrib包中的自定义分区器方法
        batch_message_num: 0 # 累计多少条消息打包发送
        linger_ms: 0 # 打包发送间隔时间
        message_max_bytes: 0 # 一条消息最大字节数，默认为1000000，约等于1M
        compression_type: 0 # 压缩方式 0 不压缩 1 zip压缩 2 Snappy压缩 3 LZ4压缩 4 ZSTD压缩
        idempotent_enabled: true # 是否开启幂等 要求ack=-1,max_open_requests=1,retry_max>0
        max_open_requests: 1 # 若开启幂等，这里要求设置为1（该包存在的问题暂未解决）
        retry_max: 3 # 发送失败重试次数
      consumer_config:
        group_id: "test12" # 消费者组id
        balance_strategy: "range" # 消费者组分区再平衡策略 range、roundrobin、sticky
        fetch_min_bytes: 1 # 每批次抓取数据最小字节数
        fetch_max_bytes: 0 # 每批次抓取数据最大字节数
        max_wait_time_ms: 500 # 每批次数据达到的超时时间（毫秒）
        from_beginning: true # 是否从0开始消费（貌似不起作用）
        offset_auto_commit_enabled: true # offset是否设置为自动提交
        offset_auto_commit_interval: 0 # offset自动提交的时间间隔
        max_poll_records: 500 # 一次拉取返回消息的最大条数
      consumer_reset_offset_config: # 指定offset消费
        enabled: false
        list:
          - topic: "second"
            partition: 0
            offset: 10

```

执行如下命令，生成项目配置：

```YAML
kubectl apply -f ./deploy/debug/example-server-configmap.yaml

# 可通过命令查看
kubectl get configmap -n dev
```

### Secret 创建密钥配置

对于一些比较敏感信息，如密码、密钥等。k8s 提供了 Secret 来实现配置，创建 example-server-secret.yaml 文件，代码如下：

```YAML
apiVersion: v1
kind: Secret
metadata:
  name: example-server-secret-jwt-certs
  namespace: dev
type: Opaque
stringData:
  rsa_private_key.pem: |-
    -----BEGIN RSA PRIVATE KEY-----
    MIICXQIBAAKBgQC2mn74yFAoDo40HGB88hp4PK3Y1CKsGiKtZDnF5E0atp3l3dWx
    c7tEdqioh9wO2EkSkvPrEvD4ToSTuDopCCkpZOVfkOWEvCkvVfWT4ygWERq8d9tX
    ua1DrXmOS/k/iseJGS57ZMrjPWMk/NQNkWBx8F7Koy1BtFaR5c9vtaBLUQIDAQAB
    AoGANk3n6NCZlfAONk0iHMZbPNq+0Lb1wLzdBzfBPqrQi1s3xf6c6HMEhC3NpqEQ
    vpFv2cg+JfCyHQYsw0mMclts8dPoIEjm5uLcxsUBr2rF4fjhqOgSKgo3XGGIO7Fh
    FMtElCzTke6x2vBSUZtgqUMnvljm56X1lfeSD4SHOQ0vF+kCQQDZHHlxxe44En+Y
    VDnP4dDY8Rw2/KW/G0QOt5WqQQOIS/uRpa8dH87LCZRodfB+ApcQY283C+BZCWWT
    gpc1coYzAkEA10+v6c5j+c7ywhBc4VGBDVj3k/lHZZoHfUgkRWWqeE+wgKD0z5do
    ljfPFGCTFRZUINNktegCAg/zM/maGrj8awJABDhze+y3Fsv1QwbQIORo5O66966t
    oitrhUW4WRjucqXXfBtWbiCCKpxDEEzKogh221CzhsRqVhDCsr1lJxrU1QJBAMrB
    NMP04nq04i/TntNjoCaEmCgVf6Fy6giVLerX8S74xBps9/yg9RE4nt3uj7M04qZC
    yzztrS5p/Lj+HJt7hd8CQQCymJCfdwu+yMqcYX8DbKsFOJWdM7Hmuw9IwdcCCAnf
    M2pG06S0jIaVieHKZF80Pnns4ZUXnIg0Idla02nIQbP/
    -----END RSA PRIVATE KEY-----
  rsa_public_key.pem: |-
    -----BEGIN PUBLIC KEY-----
    MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC2mn74yFAoDo40HGB88hp4PK3Y
    1CKsGiKtZDnF5E0atp3l3dWxc7tEdqioh9wO2EkSkvPrEvD4ToSTuDopCCkpZOVf
    kOWEvCkvVfWT4ygWERq8d9tXua1DrXmOS/k/iseJGS57ZMrjPWMk/NQNkWBx8F7K
    oy1BtFaR5c9vtaBLUQIDAQAB
    -----END PUBLIC KEY-----
---
apiVersion: v1
kind: Secret
metadata:
  name: example-server-secret-grpc-certs
  namespace: dev
type: Opaque
stringData:
  ca.crt: |-
    -----BEGIN CERTIFICATE-----
    MIIC+jCCAeICCQCwvuN/60JlaTANBgkqhkiG9w0BAQsFADA/MQswCQYDVQQGEwJj
    bjEOMAwGA1UECwwFbXlvcmcxDzANBgNVBAoMBm15dGVzdDEPMA0GA1UEAwwGbXlu
    YW1lMB4XDTIyMDYwNjA4NDcyNloXDTMyMDYwMzA4NDcyNlowPzELMAkGA1UEBhMC
    Y24xDjAMBgNVBAsMBW15b3JnMQ8wDQYDVQQKDAZteXRlc3QxDzANBgNVBAMMBm15
    bmFtZTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAPgyGuSEWjNw77j4
    ZX+eF6cmt/OYuwFLTo+iIAzueZMi2HEu/tcLHU1qcd38fG8Kq9smvJlwVo7TiO6L
    4ZlRb6FOMZURZ9DgcSek/COqLn7lXejLxQAT0NpKqvsKYwz7Rz0zvyB/KmcWbEZI
    31Xw0bS7yp0zSmHyq9nnow57ifn6gKTD4AbuORPJFpOMrZaa03tGiVsFAe6Jdl7F
    5RcK+m4iiTzEz3gxyWmO7yrsI0QYTO79rDRV53DcVEJqisfJgB+tFrQMBgaOTcA7
    55g4xhZtakQpzHGuLdFSOTYaEtuOMPqAHmycsOvjh/5jJyuWOFmUwRBuYEcxEL5E
    yK9znEsCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAzoNqrUq+pPH3GQz8p2FjlLWQ
    w6iNKaIWiEQzDF8FaTiSnPi520uFupfDazSCmn/o+iz8wqCKzKo5pEx+8iUQCxPX
    vA5QX4eCz1cfp94rNqdfqYgACoVAKjStGuQU9wOdjSQSOlfCHtE+Z74MsOIOYR47
    tEnUbxYfiryMnNRafS4rI3XI7Oq8CWgzm1T5gym+UObuxTDIAdIDBHDQExqmdw4Q
    zFmmv7iA7lrEx2FrVHHm23MaGmodXYY1cpSE743HzCbxKKa1iZAlA00UUgTjgYF/
    LzX1/ZOoPhbQTBoAAltcl9q8hS9/G/FYf5flR6FnmcSJujPPZKUzDU9lWoKCqw==
    -----END CERTIFICATE-----
  client.key: |-
    -----BEGIN RSA PRIVATE KEY-----
    MIIEowIBAAKCAQEAk8E7/x3VvyI1mD4NIZEvsEGlQ1kUfutmliokggS2Kp7LaT9S
    AyRS2zrUQBOrhMLbtr3FcawA8oPOEgp81BRbdNlBXF+RzUoV109fuu6cA/owhece
    aCnRutYkGPZUb7dAYfo6Pai18bQkw5t/iIBnuwQ6lG9ICVYJznX588y5cexMQYn2
    vnGGobPFP2uhyU/roaiW8hyCFMKaLnzdqUgRdtN7LdZ1vbFvbWL+l6i701NSehNb
    ztTLpe2pFpgaZG/dZ2oWfxPsBz1zvxRUKXeFjv57Oh00G7DbwH4NR3YiAa7PsAsQ
    g9XyX7vMgq1SKBy85bWLQ4lXdtsRCW5O0wi+FwIDAQABAoIBAD/OKWeA7cXtMvG8
    khUVyl816fUMadJpO1LLfhp7iGMdUb/wbS0YKYOsgu3s0JuzPghLKRFucrZHftzB
    lKhM5jWugBAw4Riizo29JBzoX41EJGIerT3UGVKymMIC1fGsU+w7j50C1xc6cMnb
    w5YhxSm4qMdiytvlQLTcs4S5urdb6JSseVfmRNOu1VoKkDwZMtp4NqpBwtmRohZS
    EtzTwGYoW6rPHHWl33BWcCfQUikFILnEYk1e3ikCduHT/pZUAjjk1TxEgmYvL2l9
    QViB2cjHl/3ZBEf6kDoEqhjTDRbInO8pJIBzzNIHGzRr64/JdxorPjtmWD0ON+IY
    jq+dNZkCgYEAwwuCzsZxoeuG+JaY0XhEsGqMosDWLsKcILfpS5B0HIaWOEhtKyEM
    Y/rNWZ3iDz3Ou6SLjQbQUnBs632Uu7lwzQ5mX5NPKCaGRtZ1NXr7OwOP4ovOgVK7
    7MOlQL0bCa3ydq5sTcIRa2wjcn2G6LvnBwLswqy9ZRpzQVdI8/diOb0CgYEAwe5L
    tCLhGscX51z+rrcVY9VFbUOIsTPOkoB63Y5xS1DblEvPR0IWBhBE05E/6azRG74H
    bDRUAbxQKFemkSqtywXigcScyVpo89xlF/1vEqoHqMvSvG/LqBLmrIYKLHkiDNOB
    QcwWBEMUYg6QvkFQtCATBG3VOMS3JqYHpFxzsmMCgYBFHfckvz3VW1lhDjzpPe1Z
    duiPZDaCK4SYM5QAqxf6Y9cn6sKE3hk99MqAT2qkbAquhAjP2PJJxH1UlhP20igf
    AlNU2ybZ240J3UV7xjDxILoEVfa/7UmbBVmyL1E9hN6/H+BRTcxOfOuyzrjYEH42
    BdtQwVX8Vfsx1HSNv8SXIQKBgD0YvQLwZijQhohedOJITdXY7/l/aEYjgzYGgi7g
    3HKU88ihpZjWiie44mYhAZi12F88HAyL49d0I3DCRHjBkGVA4l5tudN5x0zkp3jK
    YD7qEGWuaSLyuZ+m7MN3qybdtsuBqtw21sza8UX/PryhyTImVLd1ftJWrzje39N5
    SP3vAoGBAKpsMeA+mgg59K4/nd9VHqYytqnn8t24wjUBW/KMuLql0z1156b7rYjS
    eZrr9yOIgHVih5VBm8s8UHojdzos9sr/xXzwGiyCH3UlGMQfaKszy3ru903dAHNi
    1eGbfVTTfVCLaSQ+Uw99k8gMCmawPpWWC1CUSB4SQsUb/YBUA5IH
    -----END RSA PRIVATE KEY-----
  client.pem: |-
    -----BEGIN CERTIFICATE-----
    MIIDUzCCAjugAwIBAgIJANR7ri/2L29aMA0GCSqGSIb3DQEBCwUAMD8xCzAJBgNV
    BAYTAmNuMQ4wDAYDVQQLDAVteW9yZzEPMA0GA1UECgwGbXl0ZXN0MQ8wDQYDVQQD
    DAZteW5hbWUwHhcNMjIwNjA3MDEwMTIxWhcNMjMwNjA3MDEwMTIxWjBKMQswCQYD
    VQQGEwJjbjERMA8GA1UECwwIbXljbGllbnQxEzARBgNVBAoMCmNsaWVudGNvbXAx
    EzARBgNVBAMMCmNsaWVudG5hbWUwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK
    AoIBAQCTwTv/HdW/IjWYPg0hkS+wQaVDWRR+62aWKiSCBLYqnstpP1IDJFLbOtRA
    E6uEwtu2vcVxrADyg84SCnzUFFt02UFcX5HNShXXT1+67pwD+jCF5x5oKdG61iQY
    9lRvt0Bh+jo9qLXxtCTDm3+IgGe7BDqUb0gJVgnOdfnzzLlx7ExBifa+cYahs8U/
    a6HJT+uhqJbyHIIUwpoufN2pSBF203st1nW9sW9tYv6XqLvTU1J6E1vO1Mul7akW
    mBpkb91nahZ/E+wHPXO/FFQpd4WO/ns6HTQbsNvAfg1HdiIBrs+wCxCD1fJfu8yC
    rVIoHLzltYtDiVd22xEJbk7TCL4XAgMBAAGjRzBFMAkGA1UdEwQCMAAwCwYDVR0P
    BAQDAgXgMCsGA1UdEQQkMCKCESoub3JnLmV4YW1wbGUuY29tgg0qLmV4YW1wbGUu
    Y29tMA0GCSqGSIb3DQEBCwUAA4IBAQCCJZyrt8HHAXSEgcKz/1YJN/I6l8A6tbE/
    p6STxSU5V7T4NEyR9bCgqNKw15Ai/fnIcBMpFnWm7EE4Ku1HpRfO+BcjkwVajBaw
    HDzyPDRz/zpeo/v9W5Zpq/XfQoyDXbMhy42GaPv3gRco1igvD4ZmOvTTbWqH21xY
    HxZWuBmxYIDRdc/S81LSfts6Ccs25RJqrN1szRwXuSdjKR872aNo+pTvklHBSKvS
    eunfdr0qivF29lAKg+hs/hUAbVPMKrYX9CzdXeB3mqjipXn/1g/60/568vicFWD/
    A9XEDBRN689e9+WqNc5e1Y877y7HtgnNP15dXNNxlwshOXAZnFXc
    -----END CERTIFICATE-----
  server.key: |-
    -----BEGIN RSA PRIVATE KEY-----
    MIIEowIBAAKCAQEA1rtE2QaU1NfsIjYMZii9FJLmom0de4asGLcu80ESGkI2cpja
    NfVYdZrXqY1zi4D8bUwP1Vj7n3CN8RpQ8inveg1x3zoCDN7vIIaZagqNzb6e7ie+
    EVTqy9dbTRj/7O/+GzmEX3tCDtcahlU8p1D0ik67fruYyCd7583r/1o+E9FbDfyb
    rFPvZCMutlx9Z1vzpxg8lxD9lD1CHkHNvILKgU76hYAt9kzsH4cwzVxbu8I3DGw3
    dwYPXDkz0Yrxx18ioAEcVT6Nrpn5amf4e16XVGg+2OV6/WXrzQunElM7+suW8AVT
    IgL+5yNLxkUAplWxXj3VLEBm74MwzmOJRKpSGQIDAQABAoIBADf9HDZ6QNQ+eJNz
    BEie1q5Gn1w6lzVeKAFJwjQrA01Vdlzq/fqTvjzr4s+S0k3NlAPrkTVPvkwUTO5Q
    rkp2JLruRoKt3Vr9xdv1PIsptU+wJEB0nFpoidAclvHDUve0m8z9ckZzHQZBj+j2
    0JnPwfNOFBU3VN7bbLNEPgSWwxjnQtvP25NuTy57MTzuffw0PBZlcnVlJLkdq5qm
    vHiRm4e4I9faXQypRAoyY+ipuWcP/ZIf8DnoT63w6ZMtqXeidlPQVVMnkEoOe54Z
    W4+NcKO5PwRZG5xrnvYlMkzJ8YQUx3KhtDjMgAnqBGh6QEdx75c56DOXoo+YO7cy
    ay+dPAECgYEA/8Qd6aTIauTt5U4xGuOOczR4Z8z+01hWCOfeTQaB3X98aq7kmhrw
    iBCh2xF2EwRmxBTcpaE27tBTcGrpPsjtVqt+IoZzLvwvJBVIr2yvN4nM10e6zjxQ
    4WFuS+X1/bfoBC7SgJen6xbxHl5S0fKCenEnO2q/U3L3Ng6CLJvIyB0CgYEA1u2L
    aKHDkX0imCWJbGwKBY/YRJEl6JUJMLLCE+W7l0tbBJKPR4bkSlu+Ssgo48pkfQC0
    7X4wyADcyCI1Z6YOWgDgUnHWWDNk/toU1WNndhDuTMG0Q3FksL731hwe8k5M3mOU
    CcD/I8bvzjUmTlFGZ7AP5cdCbBE9MDYUn0QUqS0CgYB+LfbD6bOmaMKkueqn5VPF
    qkEslImzIoVvZ1TYkzlGWTq8NwQ4ssmvluUC+p5Ry34XhmNEo0BCmSxQ7a2cd+v/
    MiPMTKBOscNf/dAaG3hl4enY215hmIgeaL7NWun8/dnNFnxrTGZfKLZT0Rzyzr1o
    QqKkQYmaVcF4Pbgi4cgwUQKBgAMVLSa3C0MO9GwSPcJcWlhyE7wICO8pFBSqu7Xl
    Wad14TxTjdKM0MFwR25zKIL76C0S0YWGhdSYjpIgS3vNGe+Yi58jGtvxAMXXAu18
    m4NKAuD+9kLxYWxXkN51UMxQ6zFVZE4vgnlU5SOkB4SgDc527ANztDQwmxwR6oHO
    TC2RAoGBAOvnz8LpEfUQm4JKPM27SOsjpylElyQ45sVZFYqVvnrBew3jFW4gJGGb
    qp+oar89dPursmObAbyTSc+wutQbOnLfmoM61zwibak7H0PJLnn3RQA/lZDGCK7C
    Fvlb5x6niC8C+aGamE557PqghyWpVHDoMD0gtPv14XSwOh4GsTfB
    -----END RSA PRIVATE KEY-----
  server.pem: |-
    -----BEGIN CERTIFICATE-----
    MIIDUzCCAjugAwIBAgIJANR7ri/2L29YMA0GCSqGSIb3DQEBCwUAMD8xCzAJBgNV
    BAYTAmNuMQ4wDAYDVQQLDAVteW9yZzEPMA0GA1UECgwGbXl0ZXN0MQ8wDQYDVQQD
    DAZteW5hbWUwHhcNMjIwNjA2MDg0NzUzWhcNMjMwNjA2MDg0NzUzWjBKMQswCQYD
    VQQGEwJjbjERMA8GA1UECwwIbXlzZXJ2ZXIxEzARBgNVBAoMCnNlcnZlcmNvbXAx
    EzARBgNVBAMMCnNlcnZlcm5hbWUwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK
    AoIBAQDWu0TZBpTU1+wiNgxmKL0UkuaibR17hqwYty7zQRIaQjZymNo19Vh1mtep
    jXOLgPxtTA/VWPufcI3xGlDyKe96DXHfOgIM3u8ghplqCo3Nvp7uJ74RVOrL11tN
    GP/s7/4bOYRfe0IO1xqGVTynUPSKTrt+u5jIJ3vnzev/Wj4T0VsN/JusU+9kIy62
    XH1nW/OnGDyXEP2UPUIeQc28gsqBTvqFgC32TOwfhzDNXFu7wjcMbDd3Bg9cOTPR
    ivHHXyKgARxVPo2umflqZ/h7XpdUaD7Y5Xr9ZevNC6cSUzv6y5bwBVMiAv7nI0vG
    RQCmVbFePdUsQGbvgzDOY4lEqlIZAgMBAAGjRzBFMAkGA1UdEwQCMAAwCwYDVR0P
    BAQDAgXgMCsGA1UdEQQkMCKCESoub3JnLmV4YW1wbGUuY29tgg0qLmV4YW1wbGUu
    Y29tMA0GCSqGSIb3DQEBCwUAA4IBAQDr9poKB3GGlHAa0hzizb7hSM9bh32zdvhY
    1V5fBssmsXuJ9P2iRlNdGaEjN2JzJ4cYJbhbHWqYnQ8/H96xL2GvXC33FojCyWTk
    NvqkxCPMeLIJjvCXv+ONstAX13rFxauqFKICtBjUywc158oMs+faCniTAnlafEC+
    3ZWvh/Frj831LB4SaiPcgZ7aF/GpQXHOskx2nC17H/Y7aoJMggq4kiTbrBXFGrMe
    LqrmfsKcBhuuVTTRSeNM++jwJBn60lS/4d63kv18j4W7eYqe+It3dS2xCS4/6HiL
    Oyo2SVWBea2obkJuX8UXPG1XPCh6eXURjo6pvzDN5Tminr3IpuLg
    -----END CERTIFICATE-----
```

执行如下命令，生成项目密钥配置：

```YAML
kubectl apply -f ./deploy/debug/example-server-secret.yaml

# 可通过命令查看
kubectl get secret -n dev
```

### HTTP 服务下的 Deployment

修改/deploy/debug/example-http-server-deployment.yaml 文件代码，代码如下：

```YAML
apiVersion: apps/v1
kind: Deployment # 类型为deployment
metadata: # deployment的详细信息
  name: example-http-server-deploy
  namespace: dev
spec: # deployment描述
  replicas: 1 # 每个pod生成3个副本
  selector: # 对标签为app=example-http-server-pod的pod进行操作
    matchLabels:
      app: example-http-server-pod
  template: # pod的模板设置
    metadata: # pod的详细信息
      labels: # 给pod打上标签app=example-http-server-pod
        app: example-http-server-pod
    spec: # pod描述
      restartPolicy: Always # pod重启策略，默认为Always
      containers: # 容器信息
        - name: example-http-server
          image: xudengtang/go-example:latest
          imagePullPolicy: Always # 设置镜像拉取策略
          env: # 设置环境变量
            - name: RUN_MODE
              value: "debug"
          command: ["/go-example", "http"] # 容器运行后执行命令
          ports: # 容器暴露端口
            - name: http-port
              containerPort: 8080
              protocol: TCP
            - name: http-pprof-port
              containerPort: 6060
              protocol: TCP
          volumeMounts: # 挂载目录
            - name: configs-volume
              mountPath: /configs/debug
            - name: jwt-certs-volume
              mountPath: /assets/certs/1652421664
            - name: grpc-certs-volume
              mountPath: /assets/certs/grpc_tls
            - name: logs-volume
              mountPath: /logs
            - name: nacos-volume
              mountPath: /nacos
      volumes: # 定义存储卷
        - name: configs-volume
          configMap:
            name: example-server-configmap
        - name: jwt-certs-volume
          secret:
            secretName: example-server-secret-jwt-certs
        - name: grpc-certs-volume
          secret:
            secretName: example-server-secret-grpc-certs
        - name: logs-volume
          hostPath:
            path: /Users/zonst/dockerVolumes/exampleVolume/http/logs
        - name: nacos-volume
          hostPath:
            path: /Users/zonst/dockerVolumes/exampleVolume/http/nacos

```

#### 容器启动运行命令

在上面代码中，指定了 command，即当容器启动时会先执行/go-example http，这样服务就运行起来了。

```YAML
command: ["/go-example", "http"] # 容器运行后执行命令
```

同时还暴露容器的端口

```YAML
ports: # 容器暴露端口
  - name: http-port
    containerPort: 8080
    protocol: TCP
  - name: http-pprof-port
    containerPort: 6060
    protocol: TCP
```

#### 挂载卷

这里定义了三种卷

- config-volume：这里是 configMap，通过如下代码实现将上面定义的 configMap 配置生成到/configs/debug 目录下，即 common.yaml 和 config_center.yaml 文件；
- jwt-certs-volume 及 grpc-certs-volume：这里是 secret，同样将定义的配置生成到对应容器目录下；
- logs-volume 及 nacos-volume：这里是 HostPath，是将节点上的目录挂载到容器指定的目录下；

```YAML
    volumeMounts: # 挂载目录
        - name: configs-volume
          mountPath: /configs/debug
        - name: jwt-certs-volume
          mountPath: /assets/certs/1652421664
        - name: grpc-certs-volume
          mountPath: /assets/certs/grpc_tls
        - name: logs-volume
          mountPath: /logs
        - name: nacos-volume
          mountPath: /nacos
  volumes: # 定义存储卷
    - name: configs-volume
      configMap:
        name: example-server-configmap
    - name: jwt-certs-volume
      secret:
        secretName: example-server-secret-jwt-certs
    - name: grpc-certs-volume
      secret:
        secretName: example-server-secret-grpc-certs
    - name: logs-volume
      hostPath:
        path: /Users/zonst/dockerVolumes/exampleVolume/http/logs
    - name: nacos-volume
      hostPath:
        path: /Users/zonst/dockerVolumes/exampleVolume/http/nacos
```

运行命令部署 pod

```Bash
kubectl apply -f ./deploy/debug/example-http-server-deployment.yaml
```

### Ingress 进行转发请求

kubernetes 提供了 Ingress 资源对象，Ingress 只需要一个 NodePort 或者一个 LB 就可以满足暴露多个 Service 的需求。

#### Ingress 环境部署

这里使用 k8s 版本为 1.24.1，需要使用[https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.2.0/deploy/static/provider/cloud/deploy.yaml](https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.2.0/deploy/static/provider/cloud/deploy.yaml)对应的 Ingress 环境部署文件；

若 k8s 版本为 1.21 以前，可以参考[Service 详解](https://www.wolai.com/9wQir53epFBAAgQXNEc9D1)中 ingress 使用进行环境部署。

文件下载到项目目录下/deploy/ingress/controller/deploy.yaml，通过如下命令进行创建：

```Bash
kubectl apply -f ./deploy/ingress/controller/deploy.yaml

# 将生成svc中ingress-nginx-controller类型为LoadBalancer，要使用该类型，还需要配置其他环境，这里修改NodePort
kubectl edit svc ingress-nginx-controller -n ingress-nginx

# 再次查看会在ingress-nginx命名空间中生成svc服务
➜  ~ kubectl get svc -n ingress-nginx
NAME                                 TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                      AGE
ingress-nginx-controller             NodePort    10.96.233.75     <none>        80:30060/TCP,443:30806/TCP   42h
ingress-nginx-controller-admission   ClusterIP   10.104.136.230   <none>        443/TCP                      42h

```

至此，Ingress 环境已经部署好了。

#### https 密钥

在/deploy/ingress/tls 生成证书 tls.crt 和 tls.key

```YAML
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/C=CN/ST=BJ/L=BJ/O=nginx/CN=goexample.com"

```

创建证书密钥配置

```YAML
# 创建密钥
kubectl create secret tls tls-secret --key tls.key --cert tls.crt
```

#### service 服务

更改类型为 ClusterIP。

```YAML
apiVersion: v1
kind: Service # 类型为service
metadata: # service详细信息
  name: example-http-server-service
  namespace: dev
spec: # service描述
  selector: # 对标签为app=example-http-server-pod的pod进行代理
    app: example-http-server-pod
  type: ClusterIP
  ports: # 端口映射
    - port: 8080 # 端口名称
      targetPort: 8080 # pod节点暴露的端口
```

执行命令

```YAML
kubectl apply -f ./deploy/debug/example-http-server-service.yaml
```

#### Ingress 代理服务

创建 example-http-server-ingress.yaml 文件，实现 goexample.com 域名代理 example-http-server-service 服务。

```YAML
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: example-http-server-ingress
  namespace: dev
spec:
  ingressClassName: nginx
  tls:
    - hosts:
        - goexample.com
      secretName: tls-secret
  rules:
    - host: goexample.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: example-http-server-service
                port:
                  number: 8080
```

执行命令，生成 ingress。

```YAML
kubectl apply -f ./deploy/debug/example-http-server-ingress.yaml

# 查看ingress服务
➜  ~ kubectl get ingress -n dev
NAME                          CLASS   HOSTS           ADDRESS        PORTS   AGE
example-http-server-ingress   nginx   goexample.com   10.96.233.75   80      41h
```

#### 请求访问

使用 Ingress 部署环境生成的 ingress-nginx-controller 下暴露的端口+域名进行访问。

```YAML
https://goexample.com:30806/hello
```

### GRPC 服务下的 Deployment

和 http 服务类似，创建 example-grpc-server-deployment.yaml 文件，代码如下：

```YAML
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example-grpc-server-deploy
  namespace: dev
spec:
  replicas: 1
  selector:
    matchLabels:
      app: example-grpc-server-pod
  template:
    metadata:
      labels:
        app: example-grpc-server-pod
    spec:
      restartPolicy: Always
      containers:
        - name: example-grpc-server
          image: xudengtang/go-example:latest
          imagePullPolicy: Always
          env:
            - name: RUN_MODE
              value: "debug"
          command: ["./go-example", "grpc"]
          ports:
            - name: grpc-port
              containerPort: 9090
              protocol: TCP
          volumeMounts: # 挂载目录
            - name: configs-volume
              mountPath: /configs/debug
            - name: jwt-certs-volume
              mountPath: /assets/certs/1652421664
            - name: grpc-certs-volume
              mountPath: /assets/certs/grpc_tls
            - name: logs-volume
              mountPath: /logs
            - name: nacos-volume
              mountPath: /nacos
      volumes: # 定义存储卷
        - name: configs-volume
          configMap:
            name: example-server-configmap
        - name: jwt-certs-volume
          secret:
            secretName: example-server-secret-jwt-certs
        - name: grpc-certs-volume
          secret:
            secretName: example-server-secret-grpc-certs
        - name: logs-volume
          hostPath:
            path: /Users/zonst/dockerVolumes/exampleVolume/http/logs
        - name: nacos-volume
          hostPath:
            path: /Users/zonst/dockerVolumes/exampleVolume/http/nacos
```

运行命令部署 pod

```Bash
kubectl apply -f ./deploy/debug/example-grpc-server-deployment.yaml
```

测试请求：[https://goexample.com:30806/grpc/hello](https://goexample.com:30806/grpc/hello)

### Crontab 定时任务

创建 example-cronjob.yaml，代码如下：

```YAML
apiVersion: batch/v1
kind: CronJob
metadata:
  name: example-cronjob
  namespace: dev
  labels:
    controller: cronjob
spec:
  schedule: "*/1 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          restartPolicy: Never
          containers:
            - name: example-cronjob-test
              image: xudengtang/go-example:latest
              imagePullPolicy: Always # 设置镜像拉取策略
              env:
                - name: RUN_MODE
                  value: "debug"
              command: ["/go-example", "cronjob"]
              volumeMounts: # 挂载目录
                - name: configs-volume
                  mountPath: /configs/debug
                - name: logs-volume
                  mountPath: /logs
                - name: nacos-volume
                  mountPath: /nacos
          volumes: # 定义存储卷
            - name: configs-volume
              configMap:
                name: example-server-configmap
            - name: logs-volume
              hostPath:
                path: /Users/zonst/dockerVolumes/exampleVolume/http/logs/cronjob
            - name: nacos-volume
              hostPath:
                path: /Users/zonst/dockerVolumes/exampleVolume/http/nacos/cronjob

```

运行命令创建 cronjob

```YAML
kubectl apply -f ./deploy/debug/example-cronjob.yaml
```

### Job 一次性任务

创建 example-job.yaml 文件，代码如下：

```YAML
apiVersion: batch/v1
kind: Job
metadata:
  name: example-job
  namespace: dev
  labels:
    controller: job
spec:
  template:
    spec:
      restartPolicy: Never
      containers:
        - name: example-job-test
          image: xudengtang/go-example:latest
          imagePullPolicy: Always # 设置镜像拉取策略
          env:
            - name: RUN_MODE
              value: "debug"
          command: [ "/go-example", "job" ]
          volumeMounts: # 挂载目录
            - name: configs-volume
              mountPath: /configs/debug
            - name: logs-volume
              mountPath: /logs
            - name: nacos-volume
              mountPath: /nacos
      volumes: # 定义存储卷
        - name: configs-volume
          configMap:
            name: example-server-configmap
        - name: logs-volume
          hostPath:
            path: /Users/zonst/dockerVolumes/exampleVolume/http/logs/job
        - name: nacos-volume
          hostPath:
            path: /Users/zonst/dockerVolumes/exampleVolume/http/nacos/job
```

运行命令创建 job

```YAML
kubectl apply -f ./deploy/debug/example-job.yaml
```

### 消费者下的 Deployment

注意：由于 kafka 是通过 docker 进行部署的，代码部署到 k8s 后，需要访问到 kafka 集群，此时需要修改/etc/hosts 中 kafka 节点的 IP 地址为本机内网 IP 地址。

创建 example-consumer-server-deployment.yaml，代码如下：

```YAML
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example-consumer-server-deploy
  namespace: dev
spec:
  replicas: 1
  selector:
    matchLabels:
      app: example-consumer-server-pod
  template:
    metadata:
      labels:
        app: example-consumer-server-pod
    spec:
      restartPolicy: Always
      containers:
        - name: example-consumer-server
          image: xudengtang/go-example:latest
          imagePullPolicy: Always
          env:
            - name: RUN_MODE
              value: "debug"
          command: ["./go-example", "kafka_consumer"]
          volumeMounts: # 挂载目录
            - name: configs-volume
              mountPath: /configs/debug
            - name: logs-volume
              mountPath: /logs
            - name: nacos-volume
              mountPath: /nacos
      volumes: # 定义存储卷
        - name: configs-volume
          configMap:
            name: example-server-configmap
        - name: logs-volume
          hostPath:
            path: /Users/zonst/dockerVolumes/exampleVolume/http/logs/consumer
        - name: nacos-volume
          hostPath:
            path: /Users/zonst/dockerVolumes/exampleVolume/http/nacos/consumer
```

运行命令部署 pod

```Bash
kubectl apply -f ./deploy/debug/example-consumer-server-deployment.yaml
```

测试请求：[https://goexample.com:30806/kafka/producer-async](https://goexample.com:30806/kafka/producer-async)，查看 pod 日志。

### Socket 服务下的 Deployment

#### deployment 配置

和 http 服务的 deployment 配置类似，端口和执行命令不同。

```YAML
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example-socket-server-deploy
  namespace: dev
spec:
  replicas: 1
  selector:
    matchLabels:
      app: example-socket-server-pod
  template:
    metadata:
      labels:
        app: example-socket-server-pod
    spec:
      restartPolicy: Always
      containers:
        - name: example-socket-server
          image: xudengtang/go-example:latest
          imagePullPolicy: Always
          env:
            - name: RUN_MODE
              value: "debug"
          command: ["./go-example", "socket"]
          ports:
            - name: socket-port
              containerPort: 30000
              protocol: TCP
          volumeMounts: # 挂载目录
            - name: configs-volume
              mountPath: /configs/debug
            - name: jwt-certs-volume
              mountPath: /assets/certs/1652421664
            - name: grpc-certs-volume
              mountPath: /assets/certs/grpc_tls
            - name: logs-volume
              mountPath: /logs
            - name: nacos-volume
              mountPath: /nacos
      volumes: # 定义存储卷
        - name: configs-volume
          configMap:
            name: example-server-configmap
        - name: jwt-certs-volume
          secret:
            secretName: example-server-secret-jwt-certs
        - name: grpc-certs-volume
          secret:
            secretName: example-server-secret-grpc-certs
        - name: logs-volume
          hostPath:
            path: /Users/zonst/dockerVolumes/exampleVolume/http/logs/socket
        - name: nacos-volume
          hostPath:
            path: /Users/zonst/dockerVolumes/exampleVolume/http/nacos/socket
```

#### NodePort 类型 Service

与 http 服务类似，若直接通过 NodePort 类型的 service，可以直接创建

```YAML
apiVersion: v1
kind: Service # 类型为service
metadata: # service详细信息
  name: example-socket-server-service
  namespace: dev
spec: # service描述
  selector: # 对标签为app=example-socket-server-pod的pod进行代理
    app: example-socket-server-pod
  type: NodePort
  ports: # 端口映射
    - port: 30000 # 端口名称
      nodePort: 30000
      targetPort: 30000 # pod节点暴露的端口
```

修改项目目录下的 views/socket2.html 中的 socket 地址为：127.0.0.1:30000/socket_server 即可。

本地访问[http://127.0.0.1:8080/socket/index2](http://127.0.0.1:8080/socket/index2)，点击 OPEN - SEND 即可发起 socket 请求。

#### Ingress 代理服务

同样可以实现 ingress 代理服务。修改 service 为 ClusterIP 类型

```YAML
apiVersion: v1
kind: Service # 类型为service
metadata: # service详细信息
  name: example-socket-server-service
  namespace: dev
spec: # service描述
  selector: # 对标签为app=example-socket-server-pod的pod进行代理
    app: example-socket-server-pod
  type: ClusterIP
  ports: # 端口映射
    - port: 30000 # 端口名称
      targetPort: 30000 # pod节点暴露的端口
```

创建一个 socket 服务的 ingress，新建 example-socket-server-ingress.yaml 文件，代码如下：

```YAML
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: example-socket-server-ingress
  namespace: dev
spec:
  ingressClassName: nginx
  rules:
    - host: goexample.socket.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: example-socket-server-service
                port:
                  number: 30000
```

注意这里绑定的 host 不能和其他 ingress 重名。

修改项目目录下的 views/socket2.html 中的 socket 地址为：goexample.socket.com:30060/socket_server 即可。

本地访问[http://127.0.0.1:8080/socket/index2](http://127.0.0.1:8080/socket/index2)，点击 OPEN - SEND 即可发起 socket 请求。

### 其他补充

#### pod 动态扩缩容

以 http 服务为例，k8s 的 deployment 支持动态扩缩容

```Bash
kubectl scale deploy example-http-server-deploy --replicas=5  -n dev
```

#### pod 滚动更新

以 http 服务为例，k8s 的 deployment 支持滚动更新。

```Bash
kubectl rollout restart deploy example-http-server-deploy -n dev

# 查询更新状态
kubectl rollout status deploy example-http-server-deploy -n dev

```

#### pod 回退

k8s 的 deployment 支持版本回退是基于 deployment.yaml 配置信息进行回退的。k8s 会记录每次部署的配置信息，回退到指定版本，则重新执行指定版本的 deployment.yaml 配置信息。

因此，若要实现版本回退，需要设置 deployment 部署时拉取不同的镜像版本，例如已部署了版本 1.0、2.0、3.0，当前版本为 3.0。若此时要回退到其他版本，则可以执行如下命令：

```Bash
# 查看升级历史记录
[root@k8s-master01 ~]# kubectl rollout history deploy pc-deployment -n dev

# 版本回滚
# 这里直接使用--to-revision=1回滚到了1版本， 如果省略这个选项，就是回退到上个版本，就是2版本
[root@k8s-master01 ~]# kubectl rollout undo deployment pc-deployment --to-revision=1 -n dev
```

#### 金丝雀发布

Deployment 控制器支持控制更新过程中的控制，如“暂停(pause)”或“继续(resume)”更新操作。

```Bash
# 1.先更新一小部分pod后，立即停止更新；
[root@k8s-master01 ~]#  kubectl set image deploy pc-deployment nginx=nginx:1.17.4 -n dev && kubectl rollout pause deployment pc-deployment  -n dev
deployment.apps/pc-deployment image updated
deployment.apps/pc-deployment paused

# 2.观察更新状态
[root@k8s-master01 ~]# kubectl rollout status deploy pc-deployment -n dev　
Waiting for deployment "pc-deployment" rollout to finish: 2 out of 4 new replicas have been updated...

# 3.查看已更新的pod，看有没有出现错误等信息，若没有异常，则继续更新其他pod
[root@k8s-master01 ~]# kubectl rollout resume deploy pc-deployment -n dev
deployment.apps/pc-deployment resumed

```

## 单元测试

一般对 service 中的方法写单元测试，这里给一个小示例，项目目录下创建 internal/service/unit_test/user_service_test.go 文件，代码如下：

```go
package unit_test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/meta"
	MysqlModel "github.com/go_example/internal/model/mysql"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	"github.com/go_example/internal/service"
	"github.com/ken-house/go-contrib/prototype/mysqlClient"
	"github.com/ken-house/go-contrib/prototype/nacosClient"
	"github.com/ken-house/go-contrib/utils/env"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	// 从系统环境变量中读取运行环境
	meta.EnvMode = env.Mode()
	// 获取Nacos配置
	getNacosConfig()

	if env.IsDebugging() && !meta.DebugUseConfigCenter { // 本地调试若不使用配置中心则直接读取common.yaml文件
		viper.SetConfigFile("../../../" + meta.CfgFile + "/" + meta.EnvMode + "/common.yaml")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalln(err)
		}
		if err := viper.Unmarshal(&meta.GlobalConfig); err != nil {
			log.Fatalln(err)
		}
	} else { // 测试环境、生产环境从配置中心读取
		configCenterClient, clean, err := nacosClient.NewConfigClient(meta.NacosConfig)
		if err != nil {
			log.Fatalln(err)
		}
		defer clean()
		globalConfigStr, err := configCenterClient.GetConfig(vo.ConfigParam{
			DataId: meta.NacosConfig.DataId,
			Group:  meta.NacosConfig.Group,
		})
		if err != nil {
			log.Fatalln(err)
		}
		// 将读取到的配置信息转为全局配置
		setGlobalConfigFromData(globalConfigStr)

		// 监听实现自动感知
		configCenterClient.ListenConfig(vo.ConfigParam{
			DataId: meta.NacosConfig.DataId,
			Group:  meta.NacosConfig.Group,
			OnChange: func(namespace, group, dataId, data string) {
				setGlobalConfigFromData(data)
			},
		})
	}
}

// 从文本读取到全局配置
func setGlobalConfigFromData(data string) {
	// 将读取到的配置信息转为全局配置
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatalln(err)
	}
	if err = viper.Unmarshal(&meta.GlobalConfig); err != nil {
		log.Fatalln(err)
	}
}

// 读取nacos配置
func getNacosConfig() {
	viper.SetConfigFile("../../../" + meta.CfgFile + "/" + meta.EnvMode + "/config_center.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	// 从配置中心读取项目配置
	if err := viper.Sub("config_center").Unmarshal(&meta.NacosConfig); err != nil {
		log.Fatalln(err)
	}
}

func NewService() service.UserService {
	eg, _, _ := mysqlClient.NewGroupClient(meta.GlobalConfig.Mysql.Group)
	userMysqlRepo := MysqlRepo.NewUserRepository(eg)
	return service.NewUserService(userMysqlRepo)
}

// 单元测试函数
func TestInsertUserList(t *testing.T) {
	userList := []MysqlModel.User{
		{
			Username: "zhangsan",
			Password: "zhangsan",
			Gender:   1,
		},
		{
			Username: "lisi",
			Password: "lisi",
			Gender:   2,
		},
	}
	userService := NewService()
	err := userService.InsertUserList(userList)
	if err != nil {
		panic(err)
	}

	user, err := userService.GetUserInfoByFormData(&gin.Context{}, "zhangsan", "zhangsan")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "zhangsan", user.Username)
	assert.Equal(t, "zhangsan", user.Password)
}
```

执行单元测试命令：

```shell
➜  unit_test git:(master) ✗ go test -v -run TestInsertUserList
=== RUN   TestInsertUserList
[xorm] [info]  2022/10/11 15:37:41.585564 PING DATABASE mysql
[xorm] [info]  2022/10/11 15:37:41.636659 PING DATABASE mysql
--- PASS: TestInsertUserList (0.10s)
PASS
ok      github.com/go_example/internal/service/unit_test        0.825s
```

## 完全自动化生成 CRUD 代码

之前编写的 Crud_Makefile 自动生成 CRUD 代码，需要手动复制代码，执行 wire 后才可以进行编译。且不支持多个单词组合的名称。
本次升级为完全自动化，仅需执行一条命令，就可以直接编译。

### 安装 gsed

由于 Mac 和 Linux 有一些差异，在使用 sed 命令时，需要做特殊处理，为了和 linux 命令一致，安装 gsed，命令如下：

```shell
brew install gnu-sed
```

### 升级 Makefile 文件

```makefile
# 以下内容根据实际情况修改，也可指定参数来覆盖默认值
# 项目名
goModule = github.com/go_example
# 必须为蛇行
name = test
# 必须为蛇行
tableName = test
# 数据库连接
db = "root:root@tcp(127.0.0.1:3306)/go_example?charset=utf8mb4"

# 生成路由组名称
routerName = $(subst _,-,$(name))
# 生成表模型
tableModel = `echo $(tableName) | awk -F '_' '{ print substr(toupper($$1),1,1) substr($$1,2) substr(toupper($$2),1,1) substr($$2,2) substr(toupper($$3),1,1) substr($$3,2) substr(toupper($$4),1,1) substr($$4,2) substr(toupper($$5),1,1) substr($$5,2) substr(toupper($$6),1,1) substr($$6,2) substr(toupper($$7),1,1) substr($$7,2) substr(toupper($$8),1,1) substr($$8,2) substr(toupper($$9),1,1) substr($$9,2) substr(toupper($$10),1,1) substr($$10,2)}'`
# 对name转换为大驼峰
nameUpper = `echo $(name) | awk -F '_' '{ print substr(toupper($$1),1,1) substr($$1,2) substr(toupper($$2),1,1) substr($$2,2) substr(toupper($$3),1,1) substr($$3,2) substr(toupper($$4),1,1) substr($$4,2) substr(toupper($$5),1,1) substr($$5,2) substr(toupper($$6),1,1) substr($$6,2) substr(toupper($$7),1,1) substr($$7,2) substr(toupper($$8),1,1) substr($$8,2) substr(toupper($$9),1,1) substr($$9,2) substr(toupper($$10),1,1) substr($$10,2)}'`
# 对name转换为小驼峰
nameLitter = `echo $(name) | awk -F '_' '{ print substr(tolower($$1),1,1) substr($$1,2) substr(toupper($$2),1,1) substr($$2,2) substr(toupper($$3),1,1) substr($$3,2) substr(toupper($$4),1,1) substr($$4,2) substr(toupper($$5),1,1) substr($$5,2) substr(toupper($$6),1,1) substr($$6,2) substr(toupper($$7),1,1) substr($$7,2) substr(toupper($$8),1,1) substr($$8,2) substr(toupper($$9),1,1) substr($$9,2) substr(toupper($$10),1,1) substr($$10,2)}'`
# 模板文件目录
templatePath = "./assets/templates/crud"
# 目标文件目录
targetControllerPath = "./internal/controller/$(name)_controller.go"
targetServicePath = "./internal/service/$(name)_service.go"
targetRepositoryPath = "./internal/repository/mysql/$(name)_repository.go"
targetModelPath = "./internal/model/common.go"
targetControllerWirePath = "./internal/assembly/controller.go"
targetServiceWirePath = "./internal/assembly/service.go"
targetServerWirePath = "./internal/assembly/server.go"
targetServerHttpPath = "./internal/server/http.go"
cleanTemplatePath = "./assets/templates/crud/clean.template"

# 插入文本内容
contentPoint1 = "        New$(nameUpper)Controller,"
contentPoint2 = "	$(nameLitter)Ctr controller.$(nameUpper)Controller"
contentPoint3 = "	$(nameLitter)Ctr controller.$(nameUpper)Controller,"
contentPoint4 = "		$(nameLitter)Ctr: $(nameLitter)Ctr,"

.PHONY:test
test:
	@echo name=$(name)
	@echo routerName=$(routerName)
	@echo litterName=$(nameLitter)
	@echo upperName=$(nameUpper)
	@echo tableName=$(tableName)
	@echo tableModel=$(tableModel)


.PHONY: crud
all: controller service repository model model_db assembly http wire
	@echo "\r\n"
	@echo "-----------------------------------------------------------"
	@echo "Congratulation! make success!"
	@echo "-----------------------------------------------------------"

.PHONY: controller
controller:
	@# 生成controller
	@# 复制文件到指定位置，并改名
	@cp $(templatePath)/controller.template $(targetControllerPath)
	@# 替换文件中的宏
	@sed -i '' 's#{{PROJECT_MODULE}}#$(goModule)#g' $(targetControllerPath)
	@sed -i '' 's#{{CONTROLLER_NAME_LITTER}}#'$(nameLitter)#g $(targetControllerPath)
	@sed -i '' 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(targetControllerPath)
	@sed -i '' 's#{{TABLE_MODEL}}#'$(tableModel)#g $(targetControllerPath)
	@echo "$(targetControllerPath) generate success"

.PHONY: service
service:
	@# 生成service
	@# 复制文件到指定位置，并改名
	@cp $(templatePath)/service.template $(targetServicePath)
	@# 替换文件中的宏
	@sed -i '' 's#{{PROJECT_MODULE}}#$(goModule)#g' $(targetServicePath)
	@sed -i '' 's#{{CONTROLLER_NAME_LITTER}}#'$(nameLitter)#g $(targetServicePath)
	@sed -i '' 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(targetServicePath)
	@sed -i '' 's#{{TABLE_MODEL}}#'$(tableModel)#g $(targetServicePath)
	@echo "$(targetServicePath) generate success"

.PHONY: repository
repository:
	@# 生成repository
	@# 复制文件到指定位置，并改名
	@cp $(templatePath)/repository.template $(targetRepositoryPath)
	@# 替换文件中的宏
	@sed -i '' 's#{{PROJECT_MODULE}}#$(goModule)#g' $(targetRepositoryPath)
	@sed -i '' 's#{{CONTROLLER_NAME_LITTER}}#'$(nameLitter)#g $(targetRepositoryPath)
	@sed -i '' 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(targetRepositoryPath)
	@sed -i '' 's#{{TABLE_NAME}}#'$(tableName)#g $(targetRepositoryPath)
	@sed -i '' 's#{{TABLE_MODEL}}#'$(tableModel)#g $(targetRepositoryPath)
	@echo "$(targetRepositoryPath) generate success"

.PHONY: model
model:
	@# 生成请求model
	@sed 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/model.template >> $(targetModelPath)
	@echo "$(targetModelPath) generate success"

.PHONY: model_db
model_db:
	@# 生成数据库model
	@xorm reverse mysql $(db) ./internal/model/mysql/templates/goxorm ./internal/model/mysql $(tableName)
	@echo "./internal/model/mysql/$(tableName).go generate success"

.PHONY: assembly
assembly:
	@# 生成wire依赖文件
	@sed 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/wire/controller.template >> $(targetControllerWirePath)
	@sed 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/wire/service.template >> $(targetServiceWirePath)
	@gsed -i '/Point1/i\'$(contentPoint1) $(targetServerWirePath)
	@echo "$(targetControllerWirePath) generate success"
	@echo "$(targetServiceWirePath) generate success"
	@echo "$(targetServerWirePath) generate success"

.PHONY: http
http:
	@gsed -i '/Point2/i\'$(contentPoint2) $(targetServerHttpPath)
	@gsed -i '/Point3/i\'$(contentPoint3) $(targetServerHttpPath)
	@gsed -i '/Point4/i\'$(contentPoint4) $(targetServerHttpPath)

	@cp $(templatePath)/wire/router.template $(templatePath)/wire/router.tmp
	@sed -i '' 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/wire/router.tmp
	@sed -i '' 's#{{CONTROLLER_NAME_LITTER}}#'$(nameLitter)#g $(templatePath)/wire/router.tmp
	@sed -i '' 's#{{CONTROLLER_NAME}}#'$(routerName)#g $(templatePath)/wire/router.tmp
	@gsed -i '/Point5/r '$(templatePath)'/wire/router.tmp' $(targetServerHttpPath)

	@cp $(templatePath)/wire/router_func.template $(templatePath)/wire/router_func.tmp
	@sed -i '' 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/wire/router_func.tmp
	@sed -i '' 's#{{CONTROLLER_NAME_LITTER}}#'$(nameLitter)#g $(templatePath)/wire/router_func.tmp
	@cat $(templatePath)/wire/router_func.tmp >> $(targetServerHttpPath)
	@echo "$(targetServerHttpPath) generate success"

.PHONY: wire
wire:
	cd ./internal/assembly && $(GOPATH)/bin/wire

.PHONY: clean
clean:
	@# 删除controller
	@rm -f $(targetControllerPath)
	@echo "delete controller success"
	@# 删除service
	@rm -f $(targetServicePath)
	@echo "delete service success"
	@# 删除repository
	@rm -f $(targetRepositoryPath)
	@echo "delete repository success"
	@# 删除数据库表模型
	@rm ./internal/model/mysql/$(tableName).go
	@echo "delete db_model success"
	@rm -f $(templatePath)/wire/http.tmp
	@# 请手动删除./assets/templates/crud/clean.template对应的内容
	@rm -f $(cleanTemplatePath)
	@sed 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/model.template >> $(cleanTemplatePath)
	@sed 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/wire/controller.template >> $(cleanTemplatePath)
	@sed 's#{{CONTROLLER_NAME_UPPER}}#'$(nameUpper)#g $(templatePath)/wire/service.template >> $(cleanTemplatePath)
	@echo "请手动删除项目中./assets/templates/crud/clean.template文件内容对应以下文件："
	@echo $(targetModelPath)
	@echo $(targetControllerWirePath)
	@echo $(targetServiceWirePath)
	@echo "若手动添加过其他代码，请删除，例如以下两个文件："
	@echo $(targetServerWirePath)
	@echo $(targetServerHttpPath)
	@echo "最后请切换到./internal/assembly目录下执行wire命令"
```

### 其他

- 在 internal/assembly/server.go 和 internal/server/http.go 文件中增加了如下插入点：

```go
// Crud Makefile Point5
```

- 增加了 templates/crud/wire/router.template 和 templates/crud/wire/router_func.template 模板文件；

### 一键生成

```shell
make all name=global_setting tableName=global_setting  -f Crud_Makefile
```

执行成功后，可直接进行编译。

## 邮件服务

当服务出现异常或需要发送邮件通知时，可直接调用下面两个方法即可实现邮件发送。

### 邮箱开启 SMTP 服务

以 QQ 邮箱为例，需要在设置-账户开启 IMAP/SMTP 服务，并获取授权码作为服务器邮箱账户密码。
发送邮件服务器：smtp.qq.com，使用 SSL，端口号 465 或 587
账户名：您的 QQ 邮箱账户名
密码：这里需要获取授权码

### 安装

```bash
go get gopkg.in/gomail.v2
```

### 封装邮件发送方法

这里提供了两个方法 Send()和 SendSpecial()；
Send()方法用于发送邮件内容相同的邮件；
SendSpecial()方法用于发送邮件内容个性化的邮件；

```go
package emailClient

import (
	"log"

	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
)

type EmailClient interface {
	Send(receiveUserList []EmailUser, message EmailMessage) error
	SendSpecial(receiveUserList []EmailUser) error
}

type emailClient struct {
	cfg EmailConf
}

type EmailConf struct {
	SendEmail string `json:"send_email" mapstructure:"send_email"`
	Host      string `json:"host" mapstructure:"host"`
	Port      int    `json:"port" mapstructure:"port"`
	Account   string `json:"account" mapstructure:"account"`
	Password  string `json:"password" mapstructure:"password"`
}

func NewEmailClient(cfg EmailConf) (EmailClient, func(), error) {
	return &emailClient{
		cfg: cfg,
	}, nil, nil
}

// EmailUser 用户邮箱
type EmailUser struct {
	Name         string
	EmailAddress string
	EmailMessage
}

type EmailMessage struct {
	Subject            string
	ContentType        string
	Body               string
	AttachFilePathList []string
}

// 设置邮箱发送方
func (cli *emailClient) setSender(m *gomail.Message) error {
	if cli.cfg.SendEmail == "" {
		return errors.New("发送地址")
	}
	m.SetHeader("From", cli.cfg.SendEmail)
	return nil
}

// 设置邮件标题
func (cli *emailClient) setSubject(m *gomail.Message, subject string) {
	m.SetHeader("Subject", subject)
}

// 设置邮件内容
func (cli *emailClient) setBody(m *gomail.Message, contentType string, body string) {
	m.SetBody(contentType, body)
}

// 设置邮件附件
func (cli *emailClient) setAttach(m *gomail.Message, attachFilePathList []string) {
	if len(attachFilePathList) == 0 {
		return
	}
	for _, filePath := range attachFilePathList {
		m.Attach(filePath)
	}
}

// Send 发送邮箱，邮件内容相同
func (cli *emailClient) Send(receiveUserList []EmailUser, message EmailMessage) error {
	m := gomail.NewMessage()
	if err := cli.setSender(m); err != nil {
		return err
	}

	if len(receiveUserList) == 0 {
		return errors.New("接收用户列表为空")
	}
	// 接收方
	addressList := make([]string, 0, len(receiveUserList))
	for _, user := range receiveUserList {
		addressList = append(addressList, user.EmailAddress)
	}
	m.SetHeader("To", addressList...)

	// 邮件标题
	cli.setSubject(m, message.Subject)
	// 邮件内容
	cli.setBody(m, message.ContentType, message.Body)
	// 邮件附件
	cli.setAttach(m, message.AttachFilePathList)

	// 连接并发送邮箱
	dialer := gomail.NewDialer(cli.cfg.Host, cli.cfg.Port, cli.cfg.Account, cli.cfg.Password)
	if err := dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// SendSpecial 给对应用户发送专属邮件，每个邮件内容不同
func (cli *emailClient) SendSpecial(receiveUserList []EmailUser) error {
	m := gomail.NewMessage()
	if err := cli.setSender(m); err != nil {
		return err
	}

	dialer := gomail.NewDialer(cli.cfg.Host, cli.cfg.Port, cli.cfg.Account, cli.cfg.Password)
	dial, err := dialer.Dial()
	if err != nil {
		return err
	}
	for _, user := range receiveUserList {
		// 设置发送者
		if err := cli.setSender(m); err != nil {
			return err
		}
		// 设置接收者，指定邮件接收者名称，若使用SetHeader则默认使用邮箱名
		m.SetAddressHeader("To", user.EmailAddress, user.Name)
		// 邮件标题
		cli.setSubject(m, user.Subject)
		// 邮件内容
		cli.setBody(m, user.ContentType, user.Body)
		// 邮件附件
		cli.setAttach(m, user.AttachFilePathList)

		if err := gomail.Send(dial, m); err != nil {
			log.Printf("Could not send email to %q: %v", user.EmailAddress, err)
		}
		m.Reset()
	}
	return nil
}

```

调用方式可查看测试用例。

## IP 查询地区

通过 Ip 查询用户所在的省市区，方便用户画像分析及个性化处理。

### 安装

```bash
go get github.com/lionsoul2014/ip2region
```

### 封装方法

ip2region 通过提供的 ip 地址库（ip2region.xdb）进行查询，若有新版本，可以更新到最新版本，该文件可直接从仓库中复制或按照文档说明进行生成。
ip2region 库提供了多种方式加载，包括文件、vecotorIndex、内存缓存。这里主要介绍内存缓存的方式。创建 utils/ip2region/ipregion.go，代码如下：

```go
package ip2region

import (
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var regionSearcher *xdb.Searcher

func init() {
	var err error
	// 将文件读取加载到内存
	file, err := xdb.LoadContentFromFile("./data/ip2region.xdb")
	if err != nil {
		panic(err)
	}
	regionSearcher, err = xdb.NewWithBuffer(file)
	if err != nil {
		panic(err)
	}
}

func GetRegionByIpList(ipList []string) (map[string]string, error) {
	regionMap := make(map[string]string, len(ipList))
	for _, ip := range ipList {
		res, err := regionSearcher.SearchByStr(ip)
		if err != nil {
			return regionMap, err
		}
		regionMap[ip] = res
	}
	return regionMap, nil
}
```

## jsoniter 替代 encoding/json

jsoniter 使用 iterator api 遍历 json，0 内存分配，速度是 encoding/json 的 6 倍。

### 封装

为了实现项目使用 jsoniter 进行格式化和反格式化，不改变原有调用方式，这里封装了两个 json 包，在编译时看是否带有-tags=jsoniter 标识。
创建 utils/json/json.go 文件，当没有指定标识 jsoniter 时使用，代码如下：

```go
// Copyright 2017 Bo-Yi Wu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !jsoniter
// +build !jsoniter

package json

import "encoding/json"

var (
	// Marshal is exported by gin/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by gin/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by gin/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by gin/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by gin/json package.
	NewEncoder = json.NewEncoder
)

```

创建 utils/json/jsoniter.go 文件，当制定标识为 jsoniter 时使用，代码如下：

```go
// Copyright 2017 Bo-Yi Wu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build jsoniter
// +build jsoniter

package json

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal is exported by gin/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by gin/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by gin/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by gin/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by gin/json package.
	NewEncoder = json.NewEncoder
)

```

### 配置与使用

调用方式依然保持不变

```go
json.Unmarshal(data,&user)
```

在编译时，若指定-tags=jsoniter 时，则会使用 jsoniter 进行格式话。

```bash
go build -tags=jsoniter main.go
```

## Sentry 异常监控

当程序发生错误时，除了查询日志外，还可以添加 Sentry 异常监控，以页面渲染的方式展现，更加直观，且支持邮件提醒。

### 创建 sentry 项目

首先需要到 sentry.io 官网地址注册一个账号，创建项目，选择 go 项目。创建完成之后，就可以拿到 dsn 地址。

### 安装

目前官方 Sentry 推荐 sentry-go 来替代 raven-go 包。

```bash
$ go get github.com/getsentry/sentry-go@latest
```

### 封装包

封装一个 sentry 包，实现 sentry 初始化配置，提供捕获自定义错误和自定义消息，同时支持 Gin 框架中间件和捕获方法。

```go
package sentryClient

import (
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/ken-house/go-contrib/utils/env"
)

type SentryClient interface {
	CaptureException(err error)
	CaptureMessage(message string)
	Flush(duration time.Duration)
	CaptureExceptionForGin(ctx *gin.Context, err error)
	CaptureMessageForGin(ctx *gin.Context, message string)
	SentryMiddlewareForGin() gin.HandlerFunc
}

type sentryClient struct {
}

func NewSentryClient(cfg SentryConfig) (SentryClient, func(), error) {
	err := initSentry(cfg)
	return &sentryClient{}, func() {
		sentry.Flush(time.Second)
	}, err
}

// 初始化Sentry
func initSentry(cfg SentryConfig) error {
	err := sentry.Init(
		sentry.ClientOptions{
			Dsn:              cfg.Dsn,
			Debug:            !env.IsReleasing(),            // 线上环境为false 其他环境为true
			Transport:        sentry.NewHTTPSyncTransport(), // 同步发送到sentry
			SampleRate:       cfg.SampleRate,
			TracesSampleRate: cfg.TracesSampleRate,
			AttachStacktrace: cfg.AttachStacktrace,
			IgnoreErrors:     cfg.IgnoreErrors,
			ServerName:       cfg.ServerName,
			Environment:      env.Mode(),
		},
	)

	return err
}

type SentryConfig struct {
	Dsn              string   `json:"dsn" mapstructure:"dsn"`
	ServerName       string   `json:"server_name" mapstructure:"server_name"`
	SampleRate       float64  `json:"sample_rate" mapstructure:"sample_rate"`
	AttachStacktrace bool     `json:"attach_stacktrace" mapstructure:"attach_stacktrace"`
	TracesSampleRate float64  `json:"traces_sample_rate" mapstructure:"traces_sample_rate"`
	IgnoreErrors     []string `json:"ignore_errors" mapstructure:"ignore_errors"`
}

// CaptureException 捕获异常错误
func (cli *sentryClient) CaptureException(err error) {
	sentry.CaptureException(err)
}

// CaptureMessage 捕获异常信息
func (cli *sentryClient) CaptureMessage(message string) {
	sentry.CaptureMessage(message)
}

// Flush 刷新sentry缓存
func (cli *sentryClient) Flush(duration time.Duration) {
	sentry.Flush(duration)
}

// CaptureExceptionForGin Gin捕获自定义错误
func (cli *sentryClient) CaptureExceptionForGin(ctx *gin.Context, err error) {
	if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
		hub.CaptureException(err)
	}
}

// CaptureMessageForGin  Gin捕获自定义信息
func (cli *sentryClient) CaptureMessageForGin(ctx *gin.Context, message string) {
	if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
		hub.CaptureMessage(message)
	}
}

// SentryMiddlewareForGin Gin中间件
func (cli *sentryClient) SentryMiddlewareForGin() gin.HandlerFunc {
	return sentrygin.New(sentrygin.Options{Repanic: true})
}

```

### 代码使用

在入口文件 cmd/http.go 文件中，初始化 sentry 配置，并设置为全局变量。

```go
// 初始化sentry
		sentryClient, clean, err := sentryClient.NewSentryClient(meta.GlobalConfig.Sentry)
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		defer clean()
		meta.SentryClient = sentryClient
```

注册中间件

```go
router.Use(meta.SentryClient.SentryMiddlewareForGin())
```

除了可以之间监控 panic 错误外，还能自定义捕获错误。

```go
panic(err)
meta.SentryClient.CaptureMessageForGin(c, message)
meta.SentryClient.CaptureExceptionForGin(c, err)
```

## AES 对称加密

### 加密解密代码封装

以下代码为 AES 加密，使用 CBC 加密模式，PKCS5Padding 填充模式。

```go
package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// AesEncrypt AES加密
func AesEncrypt(contentStr string, key string, iv string) (string, error) {
	// 1.创建AES加密对象
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", nil
	}
	// 2.CBC加密模式
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))

	// 3.对原文分块并设置PKCS5Padding填充模式对数据进行填充
	// a.获取数据块大小为8位
	blockSize := block.BlockSize()
	// b.获取要填充的位数
	padNum := blockSize - len(contentStr)%blockSize
	// c.生成要填充的字节数组
	padByte := bytes.Repeat([]byte{byte(padNum)}, padNum)
	// d.加入到原文中
	contentByte := append([]byte(contentStr), padByte...)
	// 4.加密
	cryptByte := make([]byte, len(contentByte))
	blockMode.CryptBlocks(cryptByte, contentByte)
	// 5.base编码返回
	return base64.StdEncoding.EncodeToString(cryptByte), nil
}

// AesDecrypt AES解密
func AesDecrypt(cryptStr string, key string, iv string) (string, error) {
	// 1.创建AES加密对象
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", nil
	}
	// 2.CBC加密模式
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))

	// 3.Base64解码
	cryptByte, err := base64.StdEncoding.DecodeString(cryptStr)
	if err != nil {
		return "", err
	}
	// 4.解码
	contentByte := make([]byte, len(cryptByte))
	blockMode.CryptBlocks(contentByte, cryptByte)
	// 5.去除填充字节
	length := len(contentByte)
	unPadding := int(contentByte[length-1])
	return string(contentByte[:length-unPadding]), nil
}
```

### 示例代码

在代码中中使用如下代码即可：

```go
hello := "world，golang"
iv := tools.GenerateRandStr(16, 3)
encryptStr, _ := encrypt.AesEncrypt(hello, meta.GlobalConfig.Common.AesKey, iv)
decryptStr, _ := encrypt.AesDecrypt(encryptStr, meta.GlobalConfig.Common.AesKey, iv)
```

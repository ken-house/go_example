# go_example
## 简介
本项目为基础的web开发架构设计，主要采用gin框架，使用cobra生成应用和命令文件的脚手架，使用wire解决依赖注入问题，
最终实现一个高性能、可扩展、多应用的web框架。

## 主要贡献
+ https://github.com/gin-gonic/gin
+ https://github.com/spf13/cobra
+ https://github.com/google/wire
+ https://github.com/spf13/viper
+ https://xorm.io/xorm
+ https://github.com/go-redis/redis

## 版本
+ 版本v1.0.0实现了cobra+gin框架的结合；
+ 版本v1.1.0增加了wire解决依赖注入及项目文件目录整体架构；
+ 版本v1.2.0增加了xorm连接Mysql数据库；
+ 版本v1.3.0实现了redis连接；

## 使用
要求golang版本必须支持Go Modules，建议版本在1.14以上。

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
创建一个项目目录，例如：go_example，执行命令初始化go modules。
```shell
go mod init github.com/go_example
```
### 创建脚手架
安装cobra
```shell
cobra init ../go_example
```
创建web服务入口
```shell
cobra add http
```
### 引入Gin框架做web服务
在cmd/http.go文件中，创建Gin引擎，注册路由，运行服务。
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
在cmd/root.go文件中，通过cobra.OnInitialize(initConfig)对项目的配置进行初始化。
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
此时，运行项目就可以启动web服务，访问127.0.0.1:8080/hello即可访问。
```shell
go run main.go http
```
### wire解决依赖注入
安装wire
```shell
go get -u github.com/google/wire
```
对项目目录结构进行分层，目录结构如下：

![cmd-markdown-logo](./pics/1.png)

在项目中，web服务调用controller控制器，控制器调用service服务类，服务类调用repository数据仓库层，
数据仓库层调用其他包生成的服务引擎客户端（如mysql）。

这里以控制器调用服务类举例：
在controller目录下创建hello_controller.go文件。 
1. 采用面向接口编程，因此会有一个向外提供服务的interface类型的HelloController， 该接口规范了struct类型的接收者里的方法；
2. 在接收者helloController中定义要调用的服务service.HelloService；
3.  创建一个实例化当前控制器的方法，提供给上层调用；
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
在service目录下创建hello_service.go文件。同hello_controller.go文件类似定义。

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
assembly目录下创建wire文件，这里命名为controller.go，定义了控制器的依赖关系。
必须在文件第一行加上// +build wireinject表示不参与编译。
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
同样创建一个服务类的service.go文件，定义服务类的依赖关系。
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
进入assembly目录下，执行wire命令生成wire_gen.go文件。
```shell
cd ./internal/assembly
wire
```
这样就解决了文件互相依赖的问题，每层更加专注实现自己的功能，不用关心依赖方的实现。

## 连接MySQL

## 连接Redis






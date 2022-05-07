# go_example
## 简介
本项目为基础的web开发架构设计，主要采用gin框架，使用cobra生成应用和命令文件的脚手架，使用wire解决依赖注入问题，
最终实现一个高性能、可扩展、多应用的web框架。

## 主要贡献
+ https://github.com/gin-gonic/gin
+ https://github.com/spf13/cobra
+ https://github.com/google/wire
+ https://github.com/spf13/viper

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
go_example
    │  go.mod
    │  go.sum
    │  LICENSE
    │  main.go
    │  README.md
    ├─cmd # 多应用入口文件
    │      http.go # web服务
    │      root.go
    │      
    ├─configs # 配置文件
    │  └─debug
    │          common.yaml
    └─internal # 架构主体
        ├─assembly # wire 文件
        │      controller.go
        │      server.go
        │      service.go
        │      wire_gen.go
        │      
        ├─controller # 控制器
        │      hello_controller.go
        │      
        ├─meta # 常量、全局变量等
        │      meta.go
        │      
        ├─server # web服务
        │      http.go
        │      
        ├─service # 服务类
        │      hello_service.go
        │      
        └─utils # 其他公共方法
            └─negotiate
                negotiate.go
```
## 实现过程
待完成
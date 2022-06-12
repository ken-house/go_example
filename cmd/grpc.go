/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/go_example/internal/meta"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/spf13/viper"

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
		port := viper.GetString("server.grpc.port")
		listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", addr, port))
		if err != nil {
			log.Fatalf("Listen err:%+v\n", err)
		}

		// 2.创建一个grpc服务
		//creds, err := credentials.NewServerTLSFromFile("./assets/certs/grpc_tls/server.pem", "./assets/certs/grpc_tls/server.key")
		//if err != nil {
		//	log.Fatalf("NewServerTLSFromFile err:%+v\n", err)
		//}

		// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
		certificate, err := tls.LoadX509KeyPair("./assets/certs/grpc_tls/server.pem", "./assets/certs/grpc_tls/server.key")
		if err != nil {
			log.Fatalf("tls.LoadX509KeyPair err:%+v\n", err)
		}
		// 创建一个新的、空的 CertPool
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile("./assets/certs/grpc_tls/ca.crt")
		if err != nil {
			log.Fatalf("ioutil.ReadFile err:%+v", err)
		}
		// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			log.Fatalf("certPool.AppendCertsFromPEM err:%+v\n", err)
		}
		// credentials.NewTLS:构建基于 TLS 的 TransportCredentials 选项
		creds := credentials.NewTLS(&tls.Config{ // Config 结构用于配置 TLS 客户端或服务器
			Certificates: []tls.Certificate{certificate}, // 设置证书链，允许包含一个或多个
			ClientAuth:   tls.RequestClientCert,          // 请求客户端证书，握手期间不要求发送证书
			ClientCAs:    certPool,                       // 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		})

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
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}

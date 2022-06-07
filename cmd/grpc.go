/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc/credentials"

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
		//creds, err := credentials.NewServerTLSFromFile("./assets/certs/grpc_tls/server.pem", "./assets/certs/grpc_tls/server.key")
		//if err != nil {
		//	log.Fatalf("NewServerTLSFromFile err:%+v", err)
		//}

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
			ClientAuth: tls.RequireAndVerifyClientCert, // 要求必须校验客户端的证书
			ClientCAs:  certPool,                       // 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		})

		app := grpc.NewServer(grpc.Creds(creds))

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

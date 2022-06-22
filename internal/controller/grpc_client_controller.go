package controller

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"

	"github.com/spf13/viper"
	"google.golang.org/grpc/credentials"

	"github.com/go_example/internal/meta"

	"github.com/go_example/internal/lib/auth"

	"github.com/spf13/cast"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/utils/negotiate"

	pb "github.com/go_example/common/protobuf/hello"
	_ "github.com/mbobakov/grpc-consul-resolver"
)

type GrpcClientController interface {
	HelloGrpc(c *gin.Context) (int, gin.Negotiate)
}

type grpcClientController struct {
}

func NewGrpcClientController() GrpcClientController {
	return &grpcClientController{}
}

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

func (ctr *grpcClientController) HelloGrpc(c *gin.Context) (int, gin.Negotiate) {
	idStr := c.DefaultQuery("id", "1")
	// 连接服务
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//creds, err := credentials.NewClientTLSFromFile("./assets/certs/grpc_tls/server.pem", "www.example.com")
	//if err != nil {
	//	log.Printf("NewClientTLSFromFile err:%+v", err)
	//	return negotiate.JSON(http.StatusOK, gin.H{
	//		"data": gin.H{
	//			"name": "",
	//		},
	//	})
	//}

	certificate, err := tls.LoadX509KeyPair("./assets/certs/grpc_tls/client.pem", "./assets/certs/grpc_tls/client.key")
	if err != nil {
		log.Printf("tls.LoadX509KeyPair err:%+v", err)
		return negotiate.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"name": "",
			},
		})
	}
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./assets/certs/grpc_tls/ca.crt")
	if err != nil {
		log.Printf("ioutil.ReadFile err:%+v", err)
		return negotiate.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"name": "",
			},
		})
	}

	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Printf("certPool.AppendCertsFromPEM err:%+v", err)
		return negotiate.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"name": "",
			},
		})
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "www.example.com",
		RootCAs:      certPool,
	})

	grpcAuth := auth.NewAuthentication("root", "root123")
	consulAddr := viper.GetString("consul.addr")
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(grpcAuth))
	conn, err := grpc.Dial("consul://"+consulAddr+"/hello?wait=10s", grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"HealthCheckConfig": {"ServiceName": "%s"}}`, meta.HEALTHCHECK_SERVICE)), grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(grpcAuth), grpc.WithUnaryInterceptor(UnaryClientInterceptor()))
	if err != nil {
		log.Printf("Dial err:%+v", err)
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
		log.Printf("client.SayHello err:%+v", err)
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

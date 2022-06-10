package server

import (
	"context"
	"fmt"

	"github.com/go_example/common/consulClient"

	"github.com/go_example/internal/meta"
	"github.com/go_example/internal/utils/tools"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/go_example/internal/lib/auth"

	pb "github.com/go_example/common/protobuf/hello"
	"google.golang.org/grpc"
)

// 提供的服务名称及服务对应的方法
var serviceNameArr = map[string]string{
	"hello": "SayHello",
}

type GrpcServer interface {
	SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error)
	Register(server *grpc.Server)
	RegisterConsul() (consulClient.ConsulClient, []string, error)
}
type grpcServer struct {
	pb.UnimplementedHelloServiceServer
	consulClient meta.ConsulClient
}

func NewGrpcServer(consulClient meta.ConsulClient) GrpcServer {
	return &grpcServer{
		consulClient: consulClient,
	}
}

// Register 注册服务到grpc
func (srv *grpcServer) Register(server *grpc.Server) {
	pb.RegisterHelloServiceServer(server, &grpcServer{})
}

// RegisterConsul 注册服务到consul
func (srv *grpcServer) RegisterConsul() (consulClient.ConsulClient, []string, error) {
	// 注册服务到consul
	ip := tools.GetOutBoundIp()
	if ip == "" {
		return nil, nil, errors.New("GetOutBoundIp fail")
	}
	port := viper.GetString("server.grpc.port")
	serviceIdArr := make([]string, 0, len(serviceNameArr))
	for serviceName, _ := range serviceNameArr {
		if err := srv.consulClient.RegisterService(serviceName, ip, cast.ToInt(port)); err != nil {
			return nil, nil, err
		}
		serviceIdArr = append(serviceIdArr, fmt.Sprintf("%s-%s-%s", serviceName, ip, port))
	}
	return srv.consulClient, serviceIdArr, nil
}

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

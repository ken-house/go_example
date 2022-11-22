package server

import (
	"context"
	"fmt"
	"github.com/ken-house/go-contrib/prototype/nacosClient"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/ken-house/go-contrib/prototype/consulClient"

	"github.com/go_example/internal/lib/auth"
	"github.com/go_example/internal/meta"
	"github.com/go_example/internal/utils/tools"
	"github.com/pkg/errors"
	"github.com/spf13/cast"

	pb "github.com/ken-house/go-contrib/prototype/protobuf/hello"
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
	RegisterNacos() (nacosClient.ServiceClient, []vo.RegisterInstanceParam, error)
}
type grpcServer struct {
	pb.UnimplementedHelloServiceServer
	consulClient       meta.ConsulClient
	nacosServiceClient meta.NacosServiceClient
}

func NewGrpcServer(
	consulClient meta.ConsulClient,
	nacosServiceClient meta.NacosServiceClient,
) GrpcServer {
	return &grpcServer{
		consulClient:       consulClient,
		nacosServiceClient: nacosServiceClient,
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
	port := meta.GlobalConfig.Server.Grpc.Port
	serviceIdArr := make([]string, 0, len(serviceNameArr))
	for serviceName, _ := range serviceNameArr {
		if err := srv.consulClient.RegisterService(serviceName, ip, cast.ToInt(port)); err != nil {
			return nil, nil, err
		}
		serviceIdArr = append(serviceIdArr, fmt.Sprintf("%s-%s-%s", serviceName, ip, port))
	}
	return srv.consulClient, serviceIdArr, nil
}

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

// SayHello grpc服务
func (srv *grpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	span := trace.SpanFromContext(ctx)
	defer span.End()
	span.SetAttributes(attribute.Int("id", cast.ToInt(in.Id)))

	grpcAuth := auth.NewAuthentication("root", "root123")
	if err := grpcAuth.Auth(ctx); err != nil {
		return nil, err
	}

	name := "world"
	func(ctx context.Context) {
		_, span = meta.GrpcTracer.Start(ctx, "Test", trace.WithAttributes(attribute.String("name", name)))
		defer span.End()
		if in.Id != 1 {
			name = "gRPC"
		}
	}(ctx)

	fmt.Printf("id:%v\n", in.Id)

	return &pb.HelloResponse{
		Id:   in.Id,
		Name: "hello " + name,
	}, nil
}

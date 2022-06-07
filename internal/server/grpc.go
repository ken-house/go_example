package server

import (
	"context"

	"github.com/go_example/internal/lib/auth"

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

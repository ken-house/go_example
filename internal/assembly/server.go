//go:build wireinject

package assembly

import (
	"github.com/go_example/internal/server"
	"github.com/google/wire"
)

func NewHttpServer() (server.HttpServer, func(), error) {
	panic(wire.Build(
		NewGrpcClientController,
		NewHelloController,
		NewAuthController,
		NewHomeController,
		NewExcelController,
		NewJenkinsController,
		NewSmsController,
		NewKafkaController,
		NewAuthService,
		// Crud Makefile Point1
		server.NewHttpServer,
	))
}

func NewSocketServer() (server.SocketServer, func(), error) {
	panic(wire.Build(
		NewSocketController,
		server.NewSocketServer,
	))
}

func NewGrpcServer() (server.GrpcServer, func(), error) {
	panic(wire.Build(
		NewConsulClient,
		NewNacosServiceClient,
		server.NewGrpcServer,
	))
}

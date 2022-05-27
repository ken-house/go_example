//go:build wireinject

package assembly

import (
	"github.com/go_example/internal/server"
	"github.com/google/wire"
)

func NewHttpServer() (server.HttpServer, func(), error) {
	panic(wire.Build(
		NewHelloController,
		NewAuthController,
		NewHomeController,
		NewExcelController,
		NewAuthService,
		server.NewHttpServer,
	))
}

func NewSocketServer() (server.SocketServer, func(), error) {
	panic(wire.Build(
		NewSocketController,
		server.NewSocketServer,
	))
}

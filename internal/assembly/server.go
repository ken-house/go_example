//go:build wireinject

package assembly

import (
	"github.com/go_example/internal/server"
	"github.com/google/wire"
)

func NewHttpServer() (server.HttpServer, func(), error) {
	panic(wire.Build(
		NewHelloController,
		NewLoginController,
		NewHomeController,
		server.NewHttpServer,
	))
}

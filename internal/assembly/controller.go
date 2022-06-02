//go:build wireinject
// +build wireinject

package assembly

import (
	"github.com/go_example/internal/controller"
	"github.com/google/wire"
)

func NewHelloController() (controller.HelloController, func(), error) {
	panic(wire.Build(
		NewHelloService,
		controller.NewHelloController,
	))
}

func NewAuthController() (controller.AuthController, func(), error) {
	panic(wire.Build(
		NewAuthService,
		controller.NewAuthController,
	))
}

func NewHomeController() (controller.HomeController, func(), error) {
	panic(wire.Build(
		controller.NewHomeController,
	))
}

func NewExcelController() (controller.ExcelController, func(), error) {
	panic(wire.Build(
		NewUserService,
		NewExcelService,
		controller.NewExcelController,
	))
}

func NewSocketController() (controller.SocketController, func(), error) {
	panic(wire.Build(
		controller.NewSocketController,
	))
}

func NewGrpcClientController() (controller.GrpcClientController, func(), error) {
	panic(wire.Build(
		controller.NewGrpcClientController,
	))
}

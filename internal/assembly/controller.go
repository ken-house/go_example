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

func NewLoginController() (controller.LoginController, func(), error) {
	panic(wire.Build(
		NewLoginService,
		controller.NewLoginController,
	))
}

func NewHomeController() (controller.HomeController, func(), error) {
	panic(wire.Build(
		controller.NewHomeController,
	))
}

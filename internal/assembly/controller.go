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
		NewEmailService,
		controller.NewHelloController,
	))
}

func NewJenkinsController() (controller.JenkinsController, func(), error) {
	panic(wire.Build(
		NewJenkinsService,
		controller.NewJenkinsController,
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
		NewConsulClient,
		NewNacosServiceClient,
		controller.NewGrpcClientController,
	))
}

func NewSmsController() (controller.SmsController, func(), error) {
	panic(wire.Build(
		NewSmsService,
		controller.NewSmsController,
	))
}

func NewKafkaController() (controller.KafkaController, func(), error) {
	panic(wire.Build(
		NewProducerSyncClient,
		NewProducerAsyncClient,
		controller.NewKafkaController,
	))
}

func NewRabbitmqController() (controller.RabbitmqController, func(), error) {
	panic(wire.Build(
		NewRabbitmqClient,
		controller.NewRabbitmqController,
	))
}

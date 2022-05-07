//go:build wireinject

package assembly

import (
	"github.com/go_example/internal/service"
	"github.com/google/wire"
)

func NewHelloService() (service.HelloService, func(), error) {
	panic(wire.Build(
		service.NewHelloService,
	))
}

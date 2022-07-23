//go:build wireinject

package assembly

import (
	CacheRepo "github.com/go_example/internal/repository/cache"
	MongoRepo "github.com/go_example/internal/repository/mongodb"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	RedisRepo "github.com/go_example/internal/repository/redis"
	"github.com/go_example/internal/service"
	"github.com/google/wire"
	"github.com/ken-house/go-contrib/prototype/excelHandler"
)

func NewHelloService() (service.HelloService, func(), error) {
	panic(wire.Build(
		NewMongoClient,
		MongoRepo.NewUserRepository,
		NewRedisSingleClient,
		RedisRepo.NewUserRepository,
		NewMysqlGroupClient,
		MysqlRepo.NewUserRepository,
		CacheRepo.NewUserRepository,
		service.NewHelloService,
	))
}

func NewAuthService() (service.AuthService, func(), error) {
	panic(wire.Build(
		NewRedisSingleClient,
		RedisRepo.NewUserRepository,
		NewMysqlGroupClient,
		MysqlRepo.NewUserRepository,
		service.NewAuthService,
	))
}

func NewUserService() (service.UserService, func(), error) {
	panic(wire.Build(
		NewMysqlGroupClient,
		MysqlRepo.NewUserRepository,
		service.NewUserService,
	))
}

func NewExcelService() (service.ExcelUserService, func(), error) {
	panic(wire.Build(
		excelHandler.NewExcelExportHandler,
		excelHandler.NewExcelImportHandler,
		service.NewExcelUserService,
	))
}

//go:build wireinject

package assembly

import (
	cacheRepo "github.com/go_example/internal/repository/cache"
	mongoRepo "github.com/go_example/internal/repository/mongodb"
	mysqlRepo "github.com/go_example/internal/repository/mysql"
	redisRepo "github.com/go_example/internal/repository/redis"
	"github.com/go_example/internal/service"
	"github.com/google/wire"
	"github.com/ken-house/go-contrib/prototype/excelHandler"
)

func NewHelloService() (service.HelloService, func(), error) {
	panic(wire.Build(
		NewMongoClient,
		mongoRepo.NewUserRepository,
		NewRedisSingleClient,
		redisRepo.NewUserRepository,
		NewMysqlGroupClient,
		mysqlRepo.NewUserRepository,
		cacheRepo.NewUserRepository,
		service.NewHelloService,
	))
}

func NewJenkinsService() (service.JenkinsService, func(), error) {
	panic(wire.Build(
		NewJenkinsClient,
		service.NewJenkinsService,
	))
}

func NewAuthService() (service.AuthService, func(), error) {
	panic(wire.Build(
		NewRedisSingleClient,
		redisRepo.NewUserRepository,
		NewMysqlGroupClient,
		mysqlRepo.NewUserRepository,
		service.NewAuthService,
	))
}

func NewUserService() (service.UserService, func(), error) {
	panic(wire.Build(
		NewMysqlGroupClient,
		mysqlRepo.NewUserRepository,
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

func NewSmsService() (service.SmsService, func(), error) {
	panic(wire.Build(
		NewAlibabaSmsClient,
		service.NewSmsService,
	))
}

func NewKafkaService() (service.KafkaService, func(), error) {
	panic(wire.Build(
		NewConsumerClient,
		NewConsumerGroupClient,
		service.NewKafkaService,
	))
}

func NewEmailService() (service.EmailService, func(), error) {
	panic(wire.Build(
		NewEmailClient,
		service.NewEmailService,
	))
}

func NewRabbitmqService() (service.RabbitmqService, func(), error) {
	panic(wire.Build(
		NewRabbitmqClient,
		service.NewRabbitmqService,
	))
}

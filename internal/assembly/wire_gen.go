// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package assembly

import (
	"github.com/go_example/internal/controller"
	"github.com/go_example/internal/repository/cache"
	"github.com/go_example/internal/repository/mongodb"
	"github.com/go_example/internal/repository/mysql"
	"github.com/go_example/internal/repository/redis"
	"github.com/go_example/internal/server"
	"github.com/go_example/internal/service"
	"github.com/ken-house/go-contrib/prototype/excelHandler"
)

// Injectors from controller.go:

func NewHelloController() (controller.HelloController, func(), error) {
	helloService, cleanup, err := NewHelloService()
	if err != nil {
		return nil, nil, err
	}
	helloController := controller.NewHelloController(helloService)
	return helloController, func() {
		cleanup()
	}, nil
}

func NewJenkinsController() (controller.JenkinsController, func(), error) {
	jenkinsService, cleanup, err := NewJenkinsService()
	if err != nil {
		return nil, nil, err
	}
	jenkinsController := controller.NewJenkinsController(jenkinsService)
	return jenkinsController, func() {
		cleanup()
	}, nil
}

func NewAuthController() (controller.AuthController, func(), error) {
	authService, cleanup, err := NewAuthService()
	if err != nil {
		return nil, nil, err
	}
	authController := controller.NewAuthController(authService)
	return authController, func() {
		cleanup()
	}, nil
}

func NewHomeController() (controller.HomeController, func(), error) {
	homeController := controller.NewHomeController()
	return homeController, func() {
	}, nil
}

func NewExcelController() (controller.ExcelController, func(), error) {
	userService, cleanup, err := NewUserService()
	if err != nil {
		return nil, nil, err
	}
	excelUserService, cleanup2, err := NewExcelService()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	excelController := controller.NewExcelController(userService, excelUserService)
	return excelController, func() {
		cleanup2()
		cleanup()
	}, nil
}

func NewSocketController() (controller.SocketController, func(), error) {
	socketController := controller.NewSocketController()
	return socketController, func() {
	}, nil
}

func NewGrpcClientController() (controller.GrpcClientController, func(), error) {
	consulClient, err := NewConsulClient()
	if err != nil {
		return nil, nil, err
	}
	nacosServiceClient, cleanup, err := NewNacosServiceClient()
	if err != nil {
		return nil, nil, err
	}
	grpcClientController := controller.NewGrpcClientController(consulClient, nacosServiceClient)
	return grpcClientController, func() {
		cleanup()
	}, nil
}

func NewSmsController() (controller.SmsController, func(), error) {
	smsService, cleanup, err := NewSmsService()
	if err != nil {
		return nil, nil, err
	}
	smsController := controller.NewSmsController(smsService)
	return smsController, func() {
		cleanup()
	}, nil
}

// Injectors from server.go:

func NewHttpServer() (server.HttpServer, func(), error) {
	grpcClientController, cleanup, err := NewGrpcClientController()
	if err != nil {
		return nil, nil, err
	}
	helloController, cleanup2, err := NewHelloController()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	authController, cleanup3, err := NewAuthController()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	homeController, cleanup4, err := NewHomeController()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	excelController, cleanup5, err := NewExcelController()
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	jenkinsController, cleanup6, err := NewJenkinsController()
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	smsController, cleanup7, err := NewSmsController()
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	authService, cleanup8, err := NewAuthService()
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpServer := server.NewHttpServer(grpcClientController, helloController, authController, homeController, excelController, jenkinsController, smsController, authService)
	return httpServer, func() {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

func NewSocketServer() (server.SocketServer, func(), error) {
	socketController, cleanup, err := NewSocketController()
	if err != nil {
		return nil, nil, err
	}
	socketServer := server.NewSocketServer(socketController)
	return socketServer, func() {
		cleanup()
	}, nil
}

func NewGrpcServer() (server.GrpcServer, func(), error) {
	consulClient, err := NewConsulClient()
	if err != nil {
		return nil, nil, err
	}
	nacosServiceClient, cleanup, err := NewNacosServiceClient()
	if err != nil {
		return nil, nil, err
	}
	grpcServer := server.NewGrpcServer(consulClient, nacosServiceClient)
	return grpcServer, func() {
		cleanup()
	}, nil
}

// Injectors from service.go:

func NewHelloService() (service.HelloService, func(), error) {
	mysqlGroupClient, cleanup, err := NewMysqlGroupClient()
	if err != nil {
		return nil, nil, err
	}
	userRepository := mysql.NewUserRepository(mysqlGroupClient)
	redisSingleClient, cleanup2, err := NewRedisSingleClient()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	redisUserRepository := redis.NewUserRepository(redisSingleClient)
	cacheUserRepository := cache.NewUserRepository()
	mongoClient, cleanup3, err := NewMongoClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	mongodbUserRepository := mongodb.NewUserRepository(mongoClient)
	helloService := service.NewHelloService(userRepository, redisUserRepository, cacheUserRepository, mongodbUserRepository)
	return helloService, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

func NewJenkinsService() (service.JenkinsService, func(), error) {
	jenkinsClient, err := NewJenkinsClient()
	if err != nil {
		return nil, nil, err
	}
	jenkinsService := service.NewJenkinsService(jenkinsClient)
	return jenkinsService, func() {
	}, nil
}

func NewAuthService() (service.AuthService, func(), error) {
	mysqlGroupClient, cleanup, err := NewMysqlGroupClient()
	if err != nil {
		return nil, nil, err
	}
	userRepository := mysql.NewUserRepository(mysqlGroupClient)
	redisSingleClient, cleanup2, err := NewRedisSingleClient()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	redisUserRepository := redis.NewUserRepository(redisSingleClient)
	authService := service.NewAuthService(userRepository, redisUserRepository)
	return authService, func() {
		cleanup2()
		cleanup()
	}, nil
}

func NewUserService() (service.UserService, func(), error) {
	mysqlGroupClient, cleanup, err := NewMysqlGroupClient()
	if err != nil {
		return nil, nil, err
	}
	userRepository := mysql.NewUserRepository(mysqlGroupClient)
	userService := service.NewUserService(userRepository)
	return userService, func() {
		cleanup()
	}, nil
}

func NewExcelService() (service.ExcelUserService, func(), error) {
	excelExportHandler, cleanup, err := excelHandler.NewExcelExportHandler()
	if err != nil {
		return nil, nil, err
	}
	excelImportHandler := excelHandler.NewExcelImportHandler()
	excelUserService := service.NewExcelUserService(excelExportHandler, excelImportHandler)
	return excelUserService, func() {
		cleanup()
	}, nil
}

func NewSmsService() (service.SmsService, func(), error) {
	alibabaSmsClient, err := NewAlibabaSmsClient()
	if err != nil {
		return nil, nil, err
	}
	smsService := service.NewSmsService(alibabaSmsClient)
	return smsService, func() {
	}, nil
}

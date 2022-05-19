package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/controller"
	"github.com/go_example/internal/middleware"
	"github.com/go_example/internal/service"
)

type HttpServer interface {
	Register(router *gin.Engine)
}

type httpServer struct {
	helloCtr    controller.HelloController
	authCtr     controller.AuthController
	homeCtr     controller.HomeController
	authService service.AuthService
}

func NewHttpServer(
	helloCtr controller.HelloController,
	authCtr controller.AuthController,
	homeCtr controller.HomeController,
	authService service.AuthService,
) HttpServer {
	return &httpServer{
		helloCtr:    helloCtr,
		authCtr:     authCtr,
		homeCtr:     homeCtr,
		authService: authService,
	}
}

func (srv *httpServer) Register(router *gin.Engine) {
	router.GET("/hello", srv.Hello())
	// 登录接口
	router.POST("/auth/login", srv.Login())
	// 刷新token接口
	router.GET("/auth/refresh_token", srv.RefreshToken())
	// 用户主页
	router.GET("/home", middleware.JWTAuthMiddleware(srv.authService), srv.Home())
}

func (srv *httpServer) Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Negotiate(srv.helloCtr.Say(c))
	}
}

func (srv *httpServer) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Negotiate(srv.authCtr.Login(c))
	}
}

func (srv *httpServer) Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Negotiate(srv.homeCtr.Index(c))
	}
}

func (srv *httpServer) RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Negotiate(srv.authCtr.RefreshToken(c))
	}
}

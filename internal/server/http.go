package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/controller"
	"github.com/go_example/internal/middleware"
)

type HttpServer interface {
	Register(router *gin.Engine)
}

type httpServer struct {
	helloCtr controller.HelloController
	loginCtr controller.LoginController
	homeCtr  controller.HomeController
}

func NewHttpServer(
	helloCtr controller.HelloController,
	loginCtr controller.LoginController,
	homeCtr controller.HomeController,
) HttpServer {
	return &httpServer{
		helloCtr: helloCtr,
		loginCtr: loginCtr,
		homeCtr:  homeCtr,
	}
}

func (srv *httpServer) Register(router *gin.Engine) {
	router.GET("/hello", srv.Hello())
	// 登录接口
	router.POST("/login", srv.Login())
	// 用户主页
	router.GET("/home", middleware.JWTAuthMiddleware(), srv.Home())
}

func (srv *httpServer) Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Negotiate(srv.helloCtr.Say(c))
	}
}

func (srv *httpServer) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Negotiate(srv.loginCtr.Login(c))
	}
}

func (srv *httpServer) Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Negotiate(srv.homeCtr.Index(c))
	}
}

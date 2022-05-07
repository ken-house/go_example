package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/controller"
)

type HttpServer interface {
	Register(router *gin.Engine)
}

type httpServer struct {
	helloCtr controller.HelloController
}

func NewHttpServer(helloCtr controller.HelloController) HttpServer {
	return &httpServer{
		helloCtr: helloCtr,
	}
}

func (srv *httpServer) Register(router *gin.Engine) {
	router.GET("/hello", srv.Hello())
}

func (srv *httpServer) Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Negotiate(srv.helloCtr.Say(c))
	}
}

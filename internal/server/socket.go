package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/controller"
)

type SocketServer interface {
	Register(router *gin.Engine)
}

type socketServer struct {
	socketCtr controller.SocketController
}

func NewSocketServer(socketCtr controller.SocketController) SocketServer {
	return &socketServer{
		socketCtr: socketCtr,
	}
}

func (srv *socketServer) Register(router *gin.Engine) {
	router.GET("/socket_server", srv.SocketServer())

}

func (srv *socketServer) SocketServer() gin.HandlerFunc {
	return func(c *gin.Context) {
		srv.socketCtr.SocketServer(c)
	}
}

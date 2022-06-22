package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/controller"
	"github.com/go_example/internal/middleware"
	"github.com/go_example/internal/service"
)

type HttpServer interface {
	Register(router *gin.Engine)
}

type httpServer struct {
	grpcClientCtr controller.GrpcClientController
	helloCtr      controller.HelloController
	authCtr       controller.AuthController
	homeCtr       controller.HomeController
	excelCtr      controller.ExcelController
	authService   service.AuthService
}

func NewHttpServer(
	grpcClientCtr controller.GrpcClientController,
	helloCtr controller.HelloController,
	authCtr controller.AuthController,
	homeCtr controller.HomeController,
	excelCtr controller.ExcelController,
	authService service.AuthService,
) HttpServer {
	return &httpServer{
		grpcClientCtr: grpcClientCtr,
		helloCtr:      helloCtr,
		authCtr:       authCtr,
		homeCtr:       homeCtr,
		excelCtr:      excelCtr,
		authService:   authService,
	}
}

func (srv *httpServer) Register(router *gin.Engine) {
	// 告诉gin框架去哪加载讲台⽂件此处可以使⽤正则表达式
	router.LoadHTMLGlob("views/*.html")
	// Excel文件导入
	router.GET("/excel/import_index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// websocket连接
	router.GET("/socket/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "socket.html", nil)
	})
	// websocket连接
	router.GET("/socket/index2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "socket2.html", nil)
	})

	router.GET("/hello", srv.Hello())
	// 登录接口
	router.POST("/auth/login", srv.Login())
	// 刷新token接口
	router.GET("/auth/refresh_token", srv.RefreshToken())
	// 用户主页
	router.GET("/home", middleware.JWTAuthMiddleware(srv.authService), srv.Home())
	// Excel文件导出
	router.GET("/excel/export", srv.Export())
	router.POST("/excel/import", srv.Import())
	router.GET("/grpc/hello", srv.HelloGrpc())
}

func (srv *httpServer) HelloGrpc() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Negotiate(srv.grpcClientCtr.HelloGrpc(c))
	}

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

func (srv *httpServer) Export() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Negotiate(srv.excelCtr.Export(c))
	}
}

func (srv *httpServer) Import() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Negotiate(srv.excelCtr.Import(c))
	}
}

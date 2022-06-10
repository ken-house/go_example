/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go_example/internal/lib/auth"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"

	"github.com/go_example/internal/assembly"
	"github.com/go_example/internal/meta"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long:  `http server`,
	Run: func(cmd *cobra.Command, args []string) {
		// 读取证书内容
		auth.SetCerts()

		// 实例化依赖注入服务
		httpSrv, clean, err := assembly.NewHttpServer()
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		defer clean()

		// 设置gin的运行环境
		gin.SetMode(meta.EnvMode)

		// 初始化engine
		app := gin.Default()

		// 注册路由
		httpSrv.Register(app)

		addr := viper.GetString("server.http.addr")
		port := viper.GetString("server.http.port")

		// 自定义server
		srv := &http.Server{
			Addr:    fmt.Sprintf("%s:%s", addr, port),
			Handler: app,
		}

		// 启动监听服务
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %+v\n", err)
			}
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		log.Println("Server exiting")
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

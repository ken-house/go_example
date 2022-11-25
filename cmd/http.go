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

	"github.com/felixge/fgprof"
	"github.com/ken-house/go-contrib/utils/env"

	"github.com/go_example/internal/lib/auth"

	"github.com/gin-gonic/gin"

	_ "net/http/pprof"

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
		// pprof性能分析
		if !env.IsReleasing() {
			http.DefaultServeMux.Handle("/debug/fgprof", fgprof.Handler())
			pprofAddr := meta.GlobalConfig.Server.HttpPprof.Addr
			pprofPort := meta.GlobalConfig.Server.HttpPprof.Port
			go func() {
				log.Println(http.ListenAndServe(fmt.Sprintf("%s:%s", pprofAddr, pprofPort), nil))
			}()
		}

		// 初始化分布式追踪提供者
		tp, clean2, err := assembly.NewTracerProvider()
		if err != nil {
			log.Printf("%+v\n", err)
		}
		defer clean2()
		meta.HttpTracer = tp.GetTracer("httpServer")

		// 初始化指标监控提供者
		mp, clean3, err := assembly.NewMeterProvider()
		if err != nil {
			log.Printf("%+v\n", err)
		}
		defer clean3()
		meta.HttpMeter = mp.GetMeter("httpServer")

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

		// 接入OpenTelemetry，使用prometheus做指标监控
		mp.MeterPrometheusForGin(app)

		// 注册路由
		httpSrv.Register(app)

		addr := meta.GlobalConfig.Server.Http.Addr
		port := meta.GlobalConfig.Server.Http.Port

		// 自定义server
		srv := &http.Server{
			Addr:    fmt.Sprintf("%s:%s", addr, port),
			Handler: app,
		}
		fmt.Printf("Listen %s:%s\n", addr, port)

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
		fmt.Println("Server exiting")
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

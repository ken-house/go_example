/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/go_example/common/auth"

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

		// 运行应用
		port := viper.GetString("server.http.addr")
		if err := app.Run(port); err != nil {
			log.Fatalf("%+v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/routers"
	"github.com/spf13/cobra"
	"log"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long:  `http server`,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		// 加载路由 todo 路由加载可以优化
		routers.LoadRouter(r)

		// 运行http服务
		if err := r.Run(); err != nil {
			log.Fatalf("%+v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

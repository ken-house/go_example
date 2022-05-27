/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/assembly"
	"github.com/go_example/internal/utils/env"
	"github.com/spf13/cobra"
)

// socketCmd represents the socket command
var socketCmd = &cobra.Command{
	Use:   "socket",
	Short: "A brief description of your command",
	Long:  `socket server`,
	Run: func(cmd *cobra.Command, args []string) {
		socketSrv, cleanup, err := assembly.NewSocketServer()
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		defer cleanup()

		gin.SetMode(env.Mode())

		app := gin.Default()
		socketSrv.Register(app)
		if err := app.Run(":30000"); err != nil {
			log.Fatalf("socket server start failed")
		}
	},
}

func init() {
	rootCmd.AddCommand(socketCmd)
}

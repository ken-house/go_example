/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/go_example/internal/meta"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/assembly"
	"github.com/ken-house/go-contrib/utils/env"
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

		addr := meta.GlobalConfig.Server.Socket.Addr
		port := meta.GlobalConfig.Server.Socket.Port
		srv := http.Server{
			Addr:    fmt.Sprintf("%s:%s", addr, port),
			Handler: app,
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("socket server start failed")
			}
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
		<-quit

		ctx, cleanup := context.WithTimeout(context.Background(), 5*time.Second)
		defer cleanup()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		log.Println("Server exiting")
	},
}

func init() {
	rootCmd.AddCommand(socketCmd)
}

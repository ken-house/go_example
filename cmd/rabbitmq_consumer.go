/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/assembly"
	"github.com/ken-house/go-contrib/utils/env"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// rabbitmqConsumerCmd represents the rabbitmqConsumer command
var rabbitmqConsumerCmd = &cobra.Command{
	Use:   "rabbitmq_consumer",
	Short: "rabbitmq_consumer",
	Long:  `rabbitmq_consumer`,
	Run: func(cmd *cobra.Command, args []string) {
		gin.SetMode(env.Mode())
		client, cleanup, err := assembly.NewRabbitmqClient()
		if err != nil {
			log.Fatalln(err)
		}
		defer cleanup()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// 优雅关闭
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
		<-quit
	},
}

func init() {
	rootCmd.AddCommand(rabbitmqConsumerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rabbitmqConsumerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rabbitmqConsumerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
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
		rabbitmqService, cleanup, err := assembly.NewRabbitmqService()
		if err != nil {
			log.Fatalln(err)
		}
		defer cleanup()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		queueName := "hello_queue"
		// 声明并绑定
		err = rabbitmqService.DeclareAndBind(ctx, "hello_exchange", "direct", queueName, "hello")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("开始消费...")
		// 消费
		rabbitmqService.Process(ctx, queueName, "C1")

		// 优雅关闭
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
		<-quit
	},
}

func init() {
	rootCmd.AddCommand(rabbitmqConsumerCmd)
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

// kafkaConsumerCmd represents the kafkaConsumer command
var kafkaConsumerCmd = &cobra.Command{
	Use:   "kafka_consumer",
	Short: "A brief description of your command",
	Long:  `kafka consumer`,
	Run: func(cmd *cobra.Command, args []string) {
		gin.SetMode(env.Mode())
		kafkaService, cleanup, err := assembly.NewKafkaService()
		if err != nil {
			log.Fatalln(err)
		}
		defer cleanup()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// 简单处理
		//err = kafkaService.Process(ctx)
		//if err != nil {
		//	log.Fatalln(err)
		//}

		// 批量处理
		kafkaService.ProcessBatch(ctx)

		// 优雅关闭
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
		<-quit
	},
}

func init() {
	rootCmd.AddCommand(kafkaConsumerCmd)
}

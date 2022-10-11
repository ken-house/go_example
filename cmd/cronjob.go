/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/go_example/internal/assembly"
	"github.com/spf13/cobra"
	"log"
)

// cronjobCmd represents the cronjob command
var cronjobCmd = &cobra.Command{
	Use:   "cronjob",
	Short: "定时脚本",
	Long:  `每分钟执行一次`,
	Run: func(cmd *cobra.Command, args []string) {
		userService, cleanup, err := assembly.NewUserService()
		if err != nil {
			log.Fatalln(err)
		}
		defer cleanup()
		userService.Cronjob()
	},
}

func init() {
	rootCmd.AddCommand(cronjobCmd)
}

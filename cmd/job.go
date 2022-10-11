/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/go_example/internal/assembly"
	"log"

	"github.com/spf13/cobra"
)

// jobCmd represents the job command
var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "一次性任务",
	Long:  `仅执行一次`,
	Run: func(cmd *cobra.Command, args []string) {
		userService, cleanup, err := assembly.NewUserService()
		if err != nil {
			log.Fatalln(err)
		}
		defer cleanup()
		userService.Job()
	},
}

func init() {
	rootCmd.AddCommand(jobCmd)
}

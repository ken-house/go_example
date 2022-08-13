/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go_example/internal/meta"
	"github.com/ken-house/go-contrib/prototype/zapLogger"
	"github.com/ken-house/go-contrib/utils/env"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go_example",
	Short: "A brief description of your application",
	Long:  `rootCmd`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 初始化配置文件
	cobra.OnInitialize(initConfig, initLog, initValidator)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// 初始化配置文件
func initConfig() {
	// 从系统环境变量中读取运行环境
	meta.EnvMode = env.Mode()

	// viper目前仅支持单文件
	viper.SetConfigFile(meta.CfgFile + "/" + meta.EnvMode + "/common.yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
		log.Fatal(err)
	}
}

// 初始化日志
func initLog() {
	// 普通日志记录
	outputFile := fmt.Sprintf("./logs/log_%s.log", time.Now().Format("20060102"))
	//zapLogger.SimpleLogger([]string{outputFile})

	// 支持日志文件切割
	zapLogger.CustomLogger(&lumberjack.Logger{
		Filename:   outputFile,
		MaxSize:    10,
		MaxAge:     7,
		MaxBackups: 5,
		LocalTime:  false,
		Compress:   false,
	}, "")
}

// 参数验证器
func initValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validatePhone", meta.ValidatePhone)
		v.RegisterValidation("validatePassword", meta.ValidatePassword)
		v.RegisterValidation("validateUsername", meta.ValidateUsername)
		v.RegisterValidation("validateVerifyCode", meta.ValidateVerifyCode)
	}
}

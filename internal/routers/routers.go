package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func LoadRouter(e *gin.Engine) {
	e.GET("/hello", HelloHandler)
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello":    "world",
		"username": viper.GetString("common.username"),
		"pwd":      viper.GetString("common.pwd"),
	})
}

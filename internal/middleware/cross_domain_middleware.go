package middleware

import (
	"github.com/gin-gonic/gin"
)

func CrossDomainMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // 可以替换成域名，例如：https://www.baidu.com
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}

package middleware

import (
	"net/http"
	"strings"

	"github.com/go_example/common/auth"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "request header authorization is empty",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"message": "request header authorization is invalid",
			})
			c.Abort()
			return
		}

		claims, err := auth.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "request header authorization parse token err",
			})
			c.Abort()
			return
		}

		c.Set("userInfo", claims.UserInfo)
		c.Next()
	}
}

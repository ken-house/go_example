package middleware

import (
	"net/http"
	"strings"

	"github.com/go_example/internal/lib/auth"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "令牌为空",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"message": "令牌格式错误",
			})
			c.Abort()
			return
		}

		claims, err := auth.ParseToken(parts[1], "access_token")
		if err != nil {
			if err.Error() == "账号已在其他设备登录" {
				c.JSON(http.StatusOK, gin.H{
					"message": "账号已在其他设备登录",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message":     "当前登录已失效，请尝试请求refresh_token获取新令牌",
					"refresh_url": "/auth/refresh_token",
				})
			}
			c.Abort()
			return
		}

		c.Set("userInfo", claims.UserInfo)
		c.Next()
	}
}

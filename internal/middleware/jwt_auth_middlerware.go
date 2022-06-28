package middleware

import (
	"net/http"
	"strings"

	"github.com/go_example/common/errorAssets"

	"github.com/go_example/internal/service"

	"github.com/go_example/internal/lib/auth"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(authService service.AuthService) func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(http.StatusOK, errorAssets.ERR_PARAM.ToastError())
			c.Abort()
			return
		}

		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, errorAssets.ERR_PARAM.ToastError())
			c.Abort()
			return
		}

		claims, err := auth.ParseToken(authService, parts[1], "access_token")
		if err != nil {
			if err.Error() == "账号已在其他设备登录" {
				c.JSON(http.StatusOK, errorAssets.ERR_LOGIN_REMOTE.ToastError())
			} else {
				c.JSON(http.StatusOK, errorAssets.ERR_LOGIN_FAILURE.ToastError())
			}
			c.Abort()
			return
		}

		c.Set("userInfo", claims.UserInfo)
		c.Next()
	}
}

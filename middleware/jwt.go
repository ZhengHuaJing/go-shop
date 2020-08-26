package middleware

import (
	"github.com/zhenghuajing/fresh_shop/pkg/app"
	"github.com/zhenghuajing/fresh_shop/pkg/e"
	"github.com/zhenghuajing/fresh_shop/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 验证用户权限
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		token := c.GetHeader("Authorization")

		// Token 为空
		if token == "" {
			appG.Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
			c.Abort()
			return
		}

		// Token 错误
		claims, err := util.ParseToken(token)
		if err != nil {
			appG.Response(http.StatusUnauthorized, e.ERROR_CHECK_TOKEN_FAIL, nil)
			c.Abort()
			return
		}

		// Token 超时
		if time.Now().Unix() > claims.ExpiresAt {
			appG.Response(http.StatusUnauthorized, e.ERROR_TOKEN_TIMEOUT, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

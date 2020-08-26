package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
	"github.com/zhenghuajing/fresh_shop/pkg/app"
	"github.com/zhenghuajing/fresh_shop/pkg/e"
	"github.com/zhenghuajing/fresh_shop/pkg/util"
	"github.com/zhenghuajing/fresh_shop/service/user_service"
	"net/http"
)

func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		token := c.GetHeader("Authorization")

		// 解析 Token
		claims, err := util.ParseToken(token)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_PARSE_TOKEN_FAIL, nil)
			c.Abort()
			return
		}

		// 获取用户所属角色名
		names, err := global.Enforcer.GetRolesForUser(claims.UserName)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_GET_ROLE_FAIL, nil)
			c.Abort()
			return
		}

		// 根据用户验证是否有权限
		res := global.Enforcer.Enforce(names[0], c.Request.URL.Path, c.Request.Method)
		if !res {
			appG.Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

// 通过用户ID，判断是否为本人操作
func IsSelfOperate() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}
		token := c.GetHeader("Authorization")

		claims, err := util.ParseToken(token)

		// 获取用户所属角色名
		names, err := global.Enforcer.GetRolesForUser(claims.UserName)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR, nil)
			c.Abort()
			return
		}

		// 如果不是管理员
		if names[0] != global.Config.Casbin.RoleAdmin {
			userModel := model.User{
				Model: gorm.Model{
					ID: uint(claims.ID),
				},
				UserName: claims.UserName,
			}

			user, err := user_service.GetUser(userModel)
			if err != nil {
				appG.Response(http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
				c.Abort()
				return
			}
			if user.UserName != claims.Issuer {
				appG.Response(http.StatusUnauthorized, e.UNAUTHORIZED, nil)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

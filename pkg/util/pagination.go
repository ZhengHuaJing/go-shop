package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/zhenghuajing/fresh_shop/global"
)

// GetPage get page parameters
func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * global.Config.App.PageSize
	}

	return result
}

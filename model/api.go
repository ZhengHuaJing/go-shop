package model

import (
	"github.com/jinzhu/gorm"
)

// 后端api
type Api struct {
	gorm.Model

	Path        string `json:"path" gorm:"comment:'路由path'"`
	Description string `json:"description" gorm:"comment:'路由描述'"`
	ApiTag      string `json:"api_tag" gorm:"comment:'路由分组'"`
	Method      string `json:"method" gorm:"comment:'请求方法'"`
}

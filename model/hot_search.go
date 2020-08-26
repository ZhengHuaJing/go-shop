package model

import "github.com/jinzhu/gorm"

// 热搜词
type HotSearch struct {
	gorm.Model

	Keyword string `json:"keyword" gorm:"comment:'热搜词'"`
	Index   int    `json:"index" gorm:"comment:'显示顺序'"`
}

package model

import "github.com/jinzhu/gorm"

// 商品分类
type Category struct {
	gorm.Model

	ParentCategoryID int       `json:"parent_category_id" gorm:"comment:'父类目级别'"`
	ParentCategory   *Category `json:"parent_category"`

	Name         string `json:"name" gorm:"not null;unique;comment:'类别名'"`
	Code         string `json:"code" gorm:"unique;comment:'类别code'"`
	Desc         string `json:"desc" gorm:"comment:'类别描述'"`
	CategoryType int    `json:"category_type" gorm:"comment:'类目级别（1: 一级类目，2: 二级类目，3: 三级类目）'"`
	IsTab        int    `json:"is_tab" gorm:"default:0;comment:'是否导航'"`
}

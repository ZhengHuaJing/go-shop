package model

import "github.com/jinzhu/gorm"

// 分类商品广告
type Ad struct {
	gorm.Model

	CategoryID int      `json:"category_id" gorm:"comment:'分类ID'"`
	Category   Category `json:"category"`

	ProductID int     `json:"product_id" gorm:"comment:'商品ID'"`
	Product   Product `json:"product"`
}

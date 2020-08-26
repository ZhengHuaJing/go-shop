package model

import "github.com/jinzhu/gorm"

// 分类品牌广告
type Brand struct {
	gorm.Model

	CategoryID int      `json:"category_id" gorm:"comment:'分类ID'"`
	Category   Category `json:"category"`

	Name     string `json:"name" gorm:"comment:'品牌名'"`
	Desc     string `json:"desc" gorm:"comment:'品牌描述'"`
	ImageUrl string `json:"image_url" gorm:"comment:'图片url'"`
}

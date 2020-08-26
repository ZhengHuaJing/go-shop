package model

import "github.com/jinzhu/gorm"

// 首页轮播图
type IndexBanner struct {
	gorm.Model

	ProductID int     `json:"product_id" gorm:"comment:'商品ID'"`
	Product   Product `json:"product"`

	Index    int    `json:"index" gorm:"comment:'轮播顺序'"`
	ImageUrl string `json:"image_url" gorm:"comment:'轮播图'"`
}

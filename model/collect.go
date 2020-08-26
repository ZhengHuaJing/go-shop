package model

import "github.com/jinzhu/gorm"

// 用户收藏
type Collect struct {
	gorm.Model

	UserID int  `json:"user_id" gorm:"comment:'用户ID'"`
	User   User `json:"user"`

	ProductID int     `json:"product_id" gorm:"comment:'商品ID'"`
	Product   Product `json:"product"`
}

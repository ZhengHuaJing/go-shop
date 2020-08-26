package model

import "github.com/jinzhu/gorm"

// 购物车
type ShoppingCart struct {
	gorm.Model

	UserID int  `json:"user_id" gorm:"comment:'用户ID'"`
	User   User `json:"user"`

	ProductID int     `json:"product_id" gorm:"comment:'商品ID'"`
	Product   Product `json:"product"`

	Num int `json:"num" gorm:"comment:'购买数量'"`
}

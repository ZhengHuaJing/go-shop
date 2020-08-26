package model

import "github.com/jinzhu/gorm"

// 订单中的商品
type OrderProduct struct {
	gorm.Model

	OrderInfoID int     `json:"order_info_id" gorm:"comment:'订单信息ID'"`
	ProductID   int     `json:"product_id" gorm:"comment:'商品ID'"`
	Product     Product `json:"product"`

	Num int `json:"num" gorm:"comment:'商品数量'"`
}

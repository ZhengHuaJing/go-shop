package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 订单信息
type OrderInfo struct {
	gorm.Model

	UserID int  `json:"user_id" gorm:"comment:'用户ID'"`
	User   User `json:"user"`

	OrderProducts []OrderProduct `json:"order_products"`

	OrderNo    string     `json:"order_sn" gorm:"comment:'订单编号'"`
	NonceStr   string     `json:"nonce_str" gorm:"comment:'随机加密串'"`
	TradeNo    string     `json:"trade_no" gorm:"comment:'交易号'"`
	PayStatus  string     `json:"pay_status" gorm:"comment:'订单状态（未支付，超时关闭，交易创建，交易结束，待支付）'"`
	PayType    string     `json:"pay_type" gorm:"comment:'支付类型（支付宝，微信）'"`
	PostScript string     `json:"post_script" gorm:"comment:'订单留言'"`
	Money      string     `json:"money" gorm:"comment:'订单金额'"`
	PayTime    *time.Time `json:"pay_time" gorm:"comment:'支付时间'"`
}

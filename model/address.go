package model

import "github.com/jinzhu/gorm"

// 用户收货地址
type Address struct {
	gorm.Model

	UserID int  `json:"user_id" gorm:"comment:'用户ID'"`
	User   User `json:"user"`

	Province      string `json:"province" gorm:"comment:'省份'"`
	City          string `json:"city" gorm:"comment:'城市'"`
	District      string `json:"district" gorm:"comment:'区域'"`
	DetailAddress string `json:"detail_address" gorm:"comment:'详细地址'"`
	SignerName    string `json:"signer_name" gorm:"comment:'签收人'"`
	SignerMobile  string `json:"signer_mobile" gorm:"comment:'手机号'"`
}

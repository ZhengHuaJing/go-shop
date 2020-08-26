package model

import "github.com/jinzhu/gorm"

type VerifyCode struct {
	gorm.Model

	Code  string `json:"code" gorm:"comment:'验证码'"`
	Email string `json:"email" gorm:"comment:'邮箱'"`
}

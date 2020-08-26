package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

// 用户
type User struct {
	gorm.Model

	UserName string     `json:"user_name" gorm:"not null;unique;comment:'用户名'"`
	Password string     `json:"password" gorm:"not null;comment:'密码'"`
	Birthday *time.Time `json:"birthday" gorm:"comment:'出生年月'"`
	Gender   string     `json:"gender" gorm:"comment:'性别'"`
	Mobile   string     `json:"mobile" gorm:"comment:'手机号'"`
	RoleName string     `json:"role_name" gorm:"-"`
	RealName string     `json:"real_name" gorm:"comment:'联系人姓名'"`
	UUID     uuid.UUID  `json:"uuid" gorm:"comment:'UUID'"`
	IP       string     `json:"ip" gorm:"comment:'最后登录IP'"`
}

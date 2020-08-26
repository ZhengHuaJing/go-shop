package model

import (
	"github.com/jinzhu/gorm"
)

// 用户留言
type LeavingMessage struct {
	gorm.Model

	UserID int  `json:"user_id" gorm:"comment:'用户ID'"`
	User   User `json:"user"`

	MessageType int    `json:"message_type" gorm:"comment:'留言类型'"`
	Title       string `json:"title" gorm:"comment:'标题'"`
	Content     string `json:"content" gorm:"comment:'内容'"`
}

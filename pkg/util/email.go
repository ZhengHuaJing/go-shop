package util

import (
	"github.com/zhenghuajing/fresh_shop/global"
	"gopkg.in/gomail.v2"
)

// 发送邮件
func SendEmail(toEmail, msg string) {
	emailCfg := global.Config.Email
	m := gomail.NewMessage()
	m.SetHeader("From", emailCfg.HostName)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "生鲜超市")
	m.SetBody("text/html", msg)

	d := gomail.NewDialer(emailCfg.Host, emailCfg.Port, emailCfg.HostName, emailCfg.Password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

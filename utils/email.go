package utils

import (
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/request"
	"gopkg.in/gomail.v2"
	"strconv"
)

// SendEmail 发送邮件
func SendEmail(data request.Email) (err error) {
	// 创建新消息
	m := gomail.NewMessage()

	// 设置消息
	// SetHeader From 的邮箱地址必须和 NewDialer 中的邮箱地址一样。
	// SMTP服务器要求发件人地址（From）与用于认证的邮箱地址（NewDialer 中的邮箱地址）匹配，
	// 以防止未经授权的发件人发送邮件，这是一种防止垃圾邮件的安全措施。
	m.SetHeader("From", global.GGB_CONFIG.Email.From)
	m.SetHeader("To", data.To)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", data.Body)

	// 创建一个拨号器（本质就是通过邮箱的 SMTP 服务器，发送/接收电子邮件）
	port, err := strconv.Atoi(global.GGB_CONFIG.Email.Port)
	if err != nil {
		return
	}
	dialer := gomail.NewDialer(global.GGB_CONFIG.Email.Host, port, global.GGB_CONFIG.Email.Username, global.GGB_CONFIG.Email.Password)

	// 发送邮件
	err = dialer.DialAndSend(m)
	return err
}

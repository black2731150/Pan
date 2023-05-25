package email

import (
	"fmt"
	"pan/global"
	"time"

	"github.com/go-gomail/gomail"
)

func SendEmail(recipient, subjec, body string) error {
	smtpHost := global.Panserver.Config.SMTP.Host
	smtpPort := global.Panserver.Config.SMTP.Port
	sender := global.Panserver.Config.SMTP.Username
	senderPassword := global.Panserver.Config.SMTP.Password

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subjec)
	m.SetBody("text/plain", fmt.Sprintf("你的云盘系统注册验证码为：%s", body))

	d := gomail.NewDialer(smtpHost, smtpPort, sender, senderPassword)

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	global.AddToMap(recipient, body, 60*time.Second)
	return nil
}

func TestTheEmailCode(email, code string) bool {
	readlcode, ok := global.GetEmailCodeFromMap(email)
	if !ok {
		return false
	}

	if readlcode != code {
		return false
	}
	return true
}

package mail

import (
	"blog/core/setting"
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*setting.Smtp
}


func NewEmail(smtp *setting.Smtp) *Email  {
	return &Email{smtp}
}

func (e *Email) SendMail (to []string, subject, body string) error{
	message := gomail.NewMessage()
	message.SetHeader("From", e.From)
	message.SetHeader("To", to...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)
	dialer := gomail.NewDialer(e.Host, e.Port, e.Username, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(message)
}
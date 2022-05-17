package mail

import (
	"blog/internal/config"
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*config.Smtp
}


func New(smtp *config.App) *Email  {
	return &Email{smtp.Smtp}
}

func (e *Email) SendMail (to []string, subject, body string) error{
	message := gomail.NewMessage()
	message.SetHeader("From", e.FromName)
	message.SetHeader("To", to...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)
	dialer := gomail.NewDialer(e.Host, e.Port, e.Username, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(message)
}
package initialize

import (
	"blog/app/mail"
	"blog/core/global"
	"log"
)

/**
初始化邮件服务器
*/
func SetupSMTP() {
	log.Println("初始化SMTP服务器... ")
	global.Email = mail.NewEmail(global.Setting.Smtp)
	log.Print("初始化SMTP服务器完成.")
}


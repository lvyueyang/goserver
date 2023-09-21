package utils

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"server/config"
)

type SendEmail struct {
	Title   string
	Content string
	To      []string // 收件人
}

func (e *SendEmail) Send() error {
	conf := config.Config.Email

	d := gomail.NewDialer(conf.Host, conf.Port, conf.User, conf.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetAddressHeader("From", conf.From, conf.User)
	m.SetHeader("To", e.To...)
	m.SetHeader("Subject", e.Title)
	m.SetBody("text/html", e.Content)

	return d.DialAndSend(m)

}

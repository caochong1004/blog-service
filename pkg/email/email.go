package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTInfo
}

type SMTInfo struct {
	Host string
	Port int
	IsSSL bool
	UserName string
	Password string
	From string
}

func NewEmail(info *SMTInfo) *Email  {
	return &Email{SMTInfo:info}
}

func (e *Email) SendMail(to []string, subject, body string) error  {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetHeader("text/html", body)

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}
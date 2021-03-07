package mail

import (
	"gin-api/pkg/mail"
	"testing"
)

func TestSend(t *testing.T) {
	options := &mail.Options{
		MailHost: "smtp.163.com",
		MailPort: 465,
		MailUser: "xxx@163.com",
		MailPass: "xxx", // 密码或授权码
		MailTo:   []string{
			"xxx@qq.com",
		},
		Subject:  "测试邮件发送",
		Body:     "我是测试内容 hello go.",
	}

	err := mail.Send(options)
	if err != nil {
		t.Error("Mail Send error", err)
		return
	}

	t.Log("Mail Send success")
}

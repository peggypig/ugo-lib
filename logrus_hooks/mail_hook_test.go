package logrus_hooks

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

/**
* @author: codezhang
*
* @create: 2018-07-16 15:14
**/

func Test_MailHook(t *testing.T) {
	log := logrus.New()
	log.Hooks.Add(&MailHook{
		Receivers:  "*****@qq.com;*************@qq.com", //设置接受邮件，多个邮件使用;间隔
		Sender:     "**********@qq.com",                 //设置发送邮件
		SenderPass: "****************",                  //设置授权码
		SmtpServer: "smtp.qq.com",                       //smtp服务器地址
		SmtpPort:   587,                                 //smtp服务器端口
		Title:      "日志报错",                              //设置邮件主题
		MailType:   MAIL_TYPE_TEXT,                      //设置邮件内容类型
	})
	log.Error("aaa")
	time.Sleep(time.Second * 10)
}

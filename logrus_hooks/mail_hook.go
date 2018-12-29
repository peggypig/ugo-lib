package logrus_hooks

import (
	"github.com/sirupsen/logrus"
	"net/smtp"
	"strconv"
	"strings"
	"time"
	"encoding/json"
)

/**
* @author: codezhang
*
* @create: 2018-07-16 14:06
**/

type MailHook struct {
	Sender     string         //发送者
	SenderPass string         //授权码
	SmtpServer string         //邮件服务器地址
	SmtpPort   int            //邮件服务器端口
	Receivers  string         //接受者 ,以';'间隔
	Title      string         //邮件主题
	Content    interface{}    //邮件内容
	HookLevels []logrus.Level //发送邮件的日志等级
	MailType   MailType       //邮件类型
}

func (hook *MailHook) Levels() []logrus.Level {
	if len(hook.HookLevels) <= 0 {
		hook.HookLevels = append(hook.HookLevels, logrus.PanicLevel)
		hook.HookLevels = append(hook.HookLevels, logrus.ErrorLevel)
		hook.HookLevels = append(hook.HookLevels, logrus.FatalLevel)
	}
	return hook.HookLevels
}

func (hook *MailHook) Fire(entry *logrus.Entry) (err error) {
	go func(h *MailHook, en *logrus.Entry) {
		//异步发送邮件
		auth := smtp.PlainAuth("", h.Sender, h.SenderPass, h.SmtpServer)
		contentType := ""
		if h.MailType == MAIL_TYPE_TEXT {
			contentType = CONTENT_TYPE_TEXT
		} else if h.MailType == MAIL_TYPE_HTML {
			contentType = CONTENT_TYPE_HTML
		}
		if h.Content == nil {
			// 设置了邮件内容则直接发送设置的内容
			// 没有设置邮件内容
			var temp = make(map[string]interface{})
			temp[KEY_MSG] = en.Message
			temp[KEY_DATA] = en.Data
			temp[KEY_TIME] , _ = time.Parse(TIME_FORMAT,time.Now().Local().Format(TIME_FORMAT))
			temp[KEY_LEVEL] = en.Level.String()
			h.Content = temp
		}
		content, errJson := json.Marshal(h.Content)
		if errJson != nil {
			entry.Logger.Info(errJson.Error())
		}
		msg := []byte("To: " + h.Receivers + "\r\nFrom: " + h.Sender +
			"<" + h.Sender + ">\r\nSubject: " + h.Title +
			"\r\n" + contentType + "\r\n\r\n" + string(content))
		err := smtp.SendMail(hook.SmtpServer+SEMICOLON+strconv.Itoa(hook.SmtpPort), auth, hook.Sender, strings.Split(hook.Receivers, MAIL_RECEIVER_SPLIT), msg)
		if err != nil {
			entry.Logger.Info(err.Error())
		}
	}(hook, entry)
	return nil
}

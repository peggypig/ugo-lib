package logrus_hooks

/**
* @author: codezhang
*
* @create: 2018-07-16 10:12
**/

const (
	TIME_FORMAT         = "2006_01_02_15_04_05"
	POINT               = "."
	FILE_MODAL          = 0666
	NOTICE              = 1
	MAIL_RECEIVER_SPLIT = ";"
	SEMICOLON           = ":"
	CONTENT_TYPE_HTML   = "Content-Type: text/html; charset=UTF-8"
	CONTENT_TYPE_TEXT   = "Content-Type: text/plain; charset=UTF-8"
	KEY_MSG             = "msg"
	KEY_DATA            = "data"
	KEY_TIME            = "time"
	KEY_LEVEL           = "level"
)

type MailType int

const (
	MAIL_TYPE_TEXT = iota
	MAIL_TYPE_HTML
)

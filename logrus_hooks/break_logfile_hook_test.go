package logrus_hooks

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

/**
* @author: codezhang
*
* @create: 2018-07-16 14:01
**/

/**
关于 CronExpr 的编写详见https://github.com/gorhill/cronexpr
*/

func Test_BreakLogFileHook(t *testing.T) {
	log := logrus.New()
	log.Hooks.Add(&BreakLogfileHook{
		BaseFilename: "./test.log",      // 设置日志文件基础名称  包括路径  手动建立文件夹路径
		MaxAge:       25 * time.Second,  // 设置日志文件最大存在时长
		CronExpr:     "*/5 * * * * * *", // cron表达式
		DeleteFile:   true,              // 设置是否删除日志文件 默认false不删除，不删除的情况下，MaxAge的值不生效
	})
	log.Info("aaa")
	go func() {
		for {
			log.Info("aaa")
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 2)
}

package logrus_hooks

import (
	"github.com/sirupsen/logrus"
	"testing"
)

/**
* @author: codezhang
*
* @create: 2018-07-18 17:30
**/

func Test_FieldHook(t *testing.T) {
	log := logrus.New()
	for i := 0; i < 10; i++ {
		log.Info("aaa")
	}
	log.Hooks.Add(&FieldHook{
		HookLevels: logrus.AllLevels,   //设置hook的生效日志等级
		Fields: map[string]interface{}{ //添加Key:value
			"UserName": "codezhang",
			"Action":   "test",
		},
	})
	for i := 0; i < 10; i++ {
		log.Info("aaa")
	}
}

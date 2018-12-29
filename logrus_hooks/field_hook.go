package logrus_hooks

import "github.com/sirupsen/logrus"

/**
* @author: codezhang
*
* @create: 2018-07-18 17:24
**/

type FieldHook struct {
	HookLevels []logrus.Level //添加字段的日志等级，默认为全部
	Fields     logrus.Fields
}

func (hook *FieldHook) Fire(entry *logrus.Entry) error {
	if hook.Fields != nil {
		for key , value := range hook.Fields{
			entry.Data[key] = value
		}
	}
	return nil
}

func (hook *FieldHook) Levels() []logrus.Level {
	if len(hook.HookLevels) <= 0 {
		return logrus.AllLevels
	} else {
		return hook.HookLevels
	}
}

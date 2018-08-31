package validation

import (
	"reflect"
	"strings"
)

/**
*
* @description:
*
* @author: codezhang
*
* @create: 2018-08-30 14:59
**/

/**
校验器
*/
func Validator(model interface{}) (errs []error) {
	reflectType := reflect.TypeOf(model)
	reflectValue := reflect.ValueOf(model)
	if reflectType.Kind() == reflect.Struct {
		// 只针对结构体校验
		fieldNum := reflectType.NumField()
		for i := 0; i < fieldNum; i++ {
			field := reflectType.Field(i)
			// tag的key为“valid”
			tag := field.Tag.Get("valid")
			if len(tag) <= 0 {
				continue
			} else {
				// 使用英文分号将多个校验规则分开
				rules := strings.Split(tag, ";")
				for _, rule := range rules {
					// 根据每个rule进行对参数进行校验
					// 使用英文冒号将校验规则的规则和错误消息提取出来
					regularAndMsg := strings.Split(rule, ":")
					var regular string
					var msg string
					var funcName string
					var funcParam string
					value := reflectValue.Field(i)
					if len(regularAndMsg) == 2 {
						// 设置了错误信息
						regular = regularAndMsg[0] // 校验规则
						msg = regularAndMsg[1]     // 错误消息
					} else if len(regularAndMsg) == 1 {
						// 没有设置错误信息
						regular = regularAndMsg[0] // 校验规则
					}else{
						panic("error valid format(validatorPluginString(rule)[:errInfo])")
					}
					funcName = string([]rune(regular)[:strings.Index(regular, "(")])
					funcParam = string([]rune(regular)[strings.Index(regular, "(")+1 : strings.Index(regular, ")")])
					if fun := validatorFuncPlugin[funcName]; fun != nil {
						err := fun(funcParam, msg, field.Name, value, value.Kind())
						if err != nil {
							errs = append(errs, err)
						}
					}
				}
			}
		}
	}
	return
}

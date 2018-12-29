package validation

import "reflect"

/**
*
* @description: 
*
* @author: codezhang
*
* @create: 2018-08-30 15:38
**/

// 针对某个校验规则进行校验的方法  返回错误
type validatorFunc func(rule string ,msg string,fieldName string, value reflect.Value,kind reflect.Kind)(error)

// 字段属性校验插件
var validatorFuncPlugin map[string]validatorFunc

func initValidatorFuncPlugin()  {
	validatorFuncPlugin = make(map[string]validatorFunc)
}

func RegisterValidatorPlugin(rule string , plugin validatorFunc)  {
	validatorFuncPlugin[rule] = plugin
}


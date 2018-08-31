package validation

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
	"strconv"
)

/**
*
* @description:
*
* @author: codezhang
*
* @create: 2018-08-30 16:10
**/

/**
邮箱校验
*/
func mailValidatorPlugin(rule string, msg string, fieldName string, value reflect.Value, kind reflect.Kind) error {
	var err error = nil
	if len(msg) <= 0 {
		// 使用默认的错误信息
		msg = fieldName + " is invalid mail format"
	}
	if kind == reflect.String {
		// 只针对string类型的数据做正则校验
		if len(rule) <= 0 {
			// 如果没有设置邮箱的验证规则，则使用默认的
			rule = "^[A-Za-z0-9]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
		}
		matched, errReg := regexp.MatchString(rule, value.Interface().(string))
		if errReg != nil {
			err = errReg
		} else if !matched {
			err = errors.New(msg)
		}
	}else {
		panic("valid:Mail([*]) tag for field with string type,but filed named "+fieldName+" is "+kind.String())
	}
	return err
}
/**
 自定义正则校验
 */
func regexpValidatorPlugin(rule string, msg string, fieldName string, value reflect.Value, kind reflect.Kind) error {
	var err error = nil
	if len(msg) <= 0 {
		// 使用默认的错误信息
		msg = fieldName + " is invalid format"
	}
	if kind == reflect.String {
		// 只针对string类型的数据做正则校验
		if len(rule) <= 0 {
			panic("valid:Regexp(regexpString[!=null]) is expected at "+fieldName)
		}
		matched, errReg := regexp.MatchString(rule, value.Interface().(string))
		if errReg != nil {
			panic(" regexpString is invalid at "+fieldName+" with tag valid:Regexp(*)")
		} else if !matched {
			err = errors.New(msg)
		}
	}else {
		panic("valid:Mail([*]) tag for field with string type,but filed named "+fieldName+" is "+kind.String())
	}
	return err
}


/**
非零值校验
*/
func requireValidatorPlugin(rule string, msg string, fieldName string, value reflect.Value, kind reflect.Kind) error {
	var err error = nil
	if len(msg) <= 0 {
		// 使用默认的错误信息
		msg = fieldName + " is required , can not be null value "
	}
	switch kind {
	case reflect.Chan, reflect.Func, reflect.Map,
		reflect.Ptr, reflect.Interface, reflect.Slice:
		// chan func map ptr interface slice
		if value.IsNil() {
			err = errors.New(msg)
		}
	case reflect.String:
		// 字符串
		if len(value.String()) <= 0 {
			err = errors.New(msg)
		}
	default:
		panic("valid:Require() tag for field with Chan/Func/Map/Ptr/Interface/Slice/String type," +
			"but filed named "+fieldName+" is "+kind.String())
	}
	return err
}

/**
 数字的取值范围校验 两边取等号
 */
func rangeValidatorPlugin(rule string, msg string, fieldName string, value reflect.Value, kind reflect.Kind) error {
	var err error = nil
	if !strings.Contains(rule,",") {
		panic("error valid:Range tag format at "+fieldName+", valid:Range(min,max) tag is expected")
	}
	minStr := strings.Split(rule, ",")[0]
	maxStr := strings.Split(rule, ",")[1]
	if len(msg) <= 0 {
		// 使用默认的错误信息
		msg = fieldName + " must between min:" + minStr + " and max:" + maxStr
	}
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		min, errMin := strconv.ParseInt(minStr, 10, 64)
		max, errMax := strconv.ParseInt(maxStr, 10, 64)
		if errMin != nil || errMax != nil || max < min{
			panic("error valid:Range tag format at "+fieldName+", valid:Range(min[int],max[int,>min]) tag is expected")
		}
		if value.Int() <= int64(min) || value.Int() >= int64(max) {
			err = errors.New(msg)
		}
	case reflect.Float32, reflect.Float64:
		min, errMin := strconv.ParseFloat(minStr, 64)
		max, errMax := strconv.ParseFloat(maxStr, 64)
		if errMin != nil || errMax != nil {
			panic("error valid:Range tag format at "+fieldName+", valid:Range(min[float],max[float,>min]) tag is expected")
		}
		if value.Float() <= float64(min) || value.Float() >= float64(max) {
			err = errors.New(msg)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		min, errMin := strconv.ParseUint(minStr, 10, 64)
		max, errMax := strconv.ParseUint(maxStr, 10, 64)
		if errMin != nil || errMax != nil {
			panic("error valid:Range tag format at "+fieldName+", valid:Range(min[uint],max[uint,>min]) tag is expected")
		}
		if value.Uint() <= uint64(min) || value.Uint() >= uint64(max) {
			err = errors.New(msg)
		}
	default:
		panic("valid:Range(min[number],max[number,>min]) tag for field with number type,but filed named "+fieldName+" is "+kind.String())
	}
	return err
}

/**
 字符串的最大长度 取等号
 */
func maxLengthValidatorPlugin(rule string, msg string, fieldName string, value reflect.Value, kind reflect.Kind) error {
	var err error = nil
	maxLen ,errMaxLen := strconv.Atoi(rule)
	if errMaxLen != nil || maxLen < 0 {
		panic("error valid:MaxLength tag format at "+fieldName+", valid:MaxLength(max[int,>=0]) tag is expected")
	}else {
		if len(msg) <= 0 {
			// 使用默认的消息
			msg = "max length of "+fieldName+" require <= "+rule
		}
		if kind == reflect.String {
			// 只针对字符串
			if len(value.String()) >= maxLen {
				err = errors.New(msg)
			}
		}else {
			panic("valid:MaxLength(max[int,>=0]) tag for field with string type,but filed named "+fieldName+" is "+kind.String())
		}
	}
	return err
}

/**
 字符串的最小长度 不取等号
 */
func minLengthValidatorPlugin(rule string, msg string, fieldName string, value reflect.Value, kind reflect.Kind) error {
	var err error = nil
	minLen ,errMinLen := strconv.Atoi(rule)
	if errMinLen != nil || minLen < 0 {
		panic("error valid:MinLength tag format at "+fieldName+", valid:MinLength(min[int,>=0]) tag is expected")
	}else {
		if len(msg) <= 0 {
			// 使用默认的消息
			msg = "min length of "+fieldName+" require > "+rule
		}
		if kind == reflect.String {
			// 只针对字符串
			if len(value.String()) < minLen {
				err = errors.New(msg)
			}
		}else {
			panic("valid:MinLength(min[int,>=0]) tag for field with string type,but filed named "+fieldName+" is "+kind.String())
		}
	}
	return err
}


/**
 字符串的固定长度
 */
func lengthValidatorPlugin(rule string, msg string, fieldName string, value reflect.Value, kind reflect.Kind) error {
	var err error = nil
	minLen ,errMinLen := strconv.Atoi(rule)
	if errMinLen != nil || minLen < 0 {
		panic("error valid:Length tag format at "+fieldName+", valid:Length(int[>=0]) tag is expected")
	}else {
		if len(msg) <= 0 {
			// 使用默认的消息
			msg = "length of "+fieldName+" require = "+rule
		}
		if kind == reflect.String {
			// 只针对字符串
			if len(value.String()) < minLen {
				err = errors.New(msg)
			}
		}else {
			panic("valid:Length(len[int,>=0]) tag for field with string type,but filed named "+fieldName+" is "+kind.String())
		}
	}
	return err
}

/**
 字符串的长度的取值范围
 */
func lengthRangeValidatorPlugin(rule string, msg string, fieldName string, value reflect.Value, kind reflect.Kind) error {
	var err error = nil
	if !strings.Contains(rule,",") {
		panic("error valid:LengthRange tag format at "+fieldName+", valid:LengthRange(min[int,>=0],max[int,>min]) tag is expected")
	}
	minStr := strings.Split(rule, ",")[0]
	maxStr := strings.Split(rule, ",")[1]
	minLen , errMinLen := strconv.Atoi(minStr)
	maxLen , errMaxLen := strconv.Atoi(maxStr)
	if errMinLen != nil || minLen < 0 || errMaxLen != nil || maxLen <= minLen {
		panic("error valid:LengthRange tag format at "+fieldName+", valid:LengthRange(min[int,>=0],max[int,>min])  tag is expected")
	}else {
		if len(msg) <= 0 {
			// 使用默认的消息
			msg = "length of "+fieldName+" require > "+minStr+" and < "+maxStr
		}
		if kind == reflect.String {
			// 只针对字符串
			if len(value.String()) < minLen || len(value.String()) > maxLen{
				err = errors.New(msg)
			}
		}else {
			panic("valid:LengthRange(min[int,>=0],max[int,>min]) tag for field with string type,but filed named "+fieldName+" is "+kind.String())
		}
	}
	return err
}

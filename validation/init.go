package validation

/**
*
* @description:
*
* @author: codezhang
*
* @create: 2018-08-30 16:22
**/

func init() {
	initValidatorFuncPlugin()

	// 增加默认的检测插件
	validatorFuncPlugin["Mail"] = mailValidatorPlugin
	validatorFuncPlugin["Require"] = requireValidatorPlugin
	validatorFuncPlugin["Range"] = rangeValidatorPlugin
	validatorFuncPlugin["Regexp"] = regexpValidatorPlugin
	validatorFuncPlugin["MaxLength"] = maxLengthValidatorPlugin
	validatorFuncPlugin["MinLength"] = minLengthValidatorPlugin
	validatorFuncPlugin["Length"] = lengthValidatorPlugin
	validatorFuncPlugin["LengthRange"] = lengthRangeValidatorPlugin
}

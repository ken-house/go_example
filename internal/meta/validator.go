package meta

import (
	"github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
)

// 正则匹配字符串验证器公共方法
func regexpStringValidatorFunc(pattern string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		re := regexp2.MustCompile(pattern, 0)
		isMatch, _ := re.MatchString(value)
		return isMatch
	}
}

// ValidatePhone 验证手机号
var ValidatePhone = regexpStringValidatorFunc(`^1[3-9][0-9]{9}$`)

// ValidatePassword 验证密码必须为6-16位字母+数字组合
var ValidatePassword = regexpStringValidatorFunc(`^(?=.*[0-9])(?=.*[a-zA-Z])([0-9a-zA-Z]{6,16})$`)

// ValidateUsername 验证注册用户名必须为6-30位字母或数字
var ValidateUsername = regexpStringValidatorFunc(`^[0-9A-Za-z]{6,30}$`)

// ValidateVerifyCode 验证码必须为6位数字
var ValidateVerifyCode = regexpStringValidatorFunc(`^[0-9]{6}$`)

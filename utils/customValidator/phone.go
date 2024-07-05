package customValidator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ValidateMobile 校验手机号
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	//使用正则表达式验证手机号格式
	re := regexp.MustCompile(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`)
	return re.MatchString(mobile)
}

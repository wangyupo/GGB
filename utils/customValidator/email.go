package customValidator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ValidateEmail 校验邮箱
func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	// 使用正则表达式验证邮箱格式
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

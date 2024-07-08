package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/utils/customValidator"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

// Validator 自定义 customValidator
func Validator() {
	// 修改gin框架中的 customValidator 引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取结构体中 json 的 tag 的方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		var locale = global.GGB_CONFIG.System.Language // 从配置中读取语言环境（默认中文）
		uni := ut.New(enT, zhT, enT)                   // 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		global.GGB_Trans, ok = uni.GetTranslator(locale)
		if !ok {
			global.GGB_LOG.Error(fmt.Sprintf("uni.GetTranslator(%s)", locale))
			return
		}

		// 判断语言环境
		var err error
		switch locale {
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, global.GGB_Trans)
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, global.GGB_Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, global.GGB_Trans)
		}
		if err != nil {
			global.GGB_LOG.Error("注册翻译器失败", zap.Error(err))
			panic(err)
		}

		// 注册自定义校验函数
		RegisterValidatorFunc(v, "mobile", "手机号格式不正确", customValidator.ValidateMobile)
		RegisterValidatorFunc(v, "email", "邮箱格式不正确", customValidator.ValidateEmail)
	}
}

// Func customValidator.ValidateMobile
type Func func(fl validator.FieldLevel) bool

// RegisterValidatorFunc 注册自定义校验tag
func RegisterValidatorFunc(v *validator.Validate, tag string, msgStr string, fn Func) {
	// 注册tag自定义校验
	_ = v.RegisterValidation(tag, validator.Func(fn))
	//自定义错误内容
	_ = v.RegisterTranslation(tag, global.GGB_Trans, func(ut ut.Translator) error {
		return ut.Add(tag, "{0}"+msgStr, true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})

	return
}

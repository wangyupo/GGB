package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"strings"
)

// HandleValidatorError 处理字段校验异常
func HandleValidatorError(err error, c *gin.Context) {
	//如何返回错误信息
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.FailWithValidate(removeTopStruct(errs.Translate(global.GGB_Trans)), c)
	return
}

// removeTopStruct 定义一个去掉结构体名称前缀的自定义方法
func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		// 从文本的逗号开始切分（处理前: "PasswordLoginForm.mobile": "mobile为必填字段"；处理后"mobile": "mobile为必填字段"）
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

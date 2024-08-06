package common

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/request"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
	"path/filepath"
	"text/template"
)

type EmailApi struct{}

// SendEmail 发送邮件
//
//	@Tags		CommonEmail
//	@Summary	发送邮件
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Success	200	{object}	response.MsgResponse	"返回邮件发送成功提示"
//	@Router		/common/email [GET]
func (e *EmailApi) SendEmail(c *gin.Context) {
	// 使用 html 模板作为邮件内容的载体
	t, err := template.ParseFiles(filepath.Join("resource/template", "email.html"))
	if err != nil {
		global.GGB_LOG.Error("Email模板读取失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 模板填充数据
	var body bytes.Buffer
	err = t.Execute(&body, map[string]interface{}{
		"content": "这是来自 GGB 的一封测试邮件",
	})
	if err != nil {
		global.GGB_LOG.Error("Email模板数据填充失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 发送邮件
	err = utils.SendEmail(request.Email{
		To:      "xxx@qq.com",
		Subject: "欢迎使用GGB后端服务架构",
		Body:    body.String(),
	})
	if err != nil {
		global.GGB_LOG.Error("发送邮件失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithMessage("邮件发送成功", c)
}

package common

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wangyupo/GGB/api/v1"
)

type EmailRouter struct{}

func (e *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	emailRouterWithoutRecord := Router.Group("/common/email")
	emailApi := v1.ApiGroupApp.CommonApiGroup.EmailApi
	{
		emailRouterWithoutRecord.GET("", emailApi.SendEmail) // 发送邮件
	}
}

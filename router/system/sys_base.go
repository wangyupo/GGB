package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wangyupo/GGB/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouterWithoutRecord := Router
	sysBaseApi := v1.ApiGroupApp.SysApiGroup.SysBaseApi
	{
		baseRouterWithoutRecord.POST("/login", sysBaseApi.Login)                  // 登录
		baseRouterWithoutRecord.POST("/captcha", sysBaseApi.GetCaptcha)           // 获取图形验证码
		baseRouterWithoutRecord.POST("/captcha/verify", sysBaseApi.VerifyCaptcha) // 校验图形验证码
	}
}

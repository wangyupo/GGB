package log

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/v1"
)

type LoginLogRouter struct{}

func (s *LoginLogRouter) InitLoginLogRouter(Router *gin.RouterGroup) {
	logLoginRouter := Router.Group("/log/login")
	sysLoginLogApi := v1.ApiGroupApp.LogApiGroup.SysLogLoginApi
	{
		logLoginRouter.GET("", sysLoginLogApi.GetSysLogLoginList) // 获取登录日志列表
	}
}

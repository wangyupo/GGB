package log

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api"
)

type LoginLogRouter struct{}

func (s *LoginLogRouter) InitLoginLogRouter(Router *gin.RouterGroup) {
	logLoginRouter := Router.Group("/log/login")
	sysLoginLogApi := api.ApiGroupApp.SysLogApiGroup.SysLoginLogApi
	{
		logLoginRouter.GET("", sysLoginLogApi.GetLoginLogList)
	}
}

package log

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/log"
)

type LoginLogRouter struct{}

func (s *LoginLogRouter) InitLoginLogRouter(Router *gin.RouterGroup) {
	logLoginRouter := Router.Group("/log/login")
	{
		logLoginRouter.GET("", log.GetLoginLogList)
	}
}

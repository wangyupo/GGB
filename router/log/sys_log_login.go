package log

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/log"
)

func LoginRouter(r *gin.RouterGroup) {
	logLoginRouter := r.Group("/log/login")
	{
		logLoginRouter.GET("", log.GetLoginLogList)
	}
}

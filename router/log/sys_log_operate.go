package log

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api"
)

type OperateLogRouter struct{}

func (o *OperateLogRouter) InitSysLogOperateRouter(r *gin.RouterGroup) {
	logOperateRouter := r.Group("/log/operate")
	logOperateApi := api.ApiGroupApp.LogApiGroup.SysLogOperateApi
	{
		logOperateRouter.GET("", logOperateApi.GetSysLogOperateList)
		logOperateRouter.POST("", logOperateApi.CreateSysLogOperate)
		logOperateRouter.DELETE("/:id", logOperateApi.DeleteSysLogOperate)
	}
}

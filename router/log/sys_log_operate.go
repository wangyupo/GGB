package log

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/v1"
)

type OperateLogRouter struct{}

func (o *OperateLogRouter) InitSysLogOperateRouter(r *gin.RouterGroup) {
	logOperateRouter := r.Group("/log/operate")
	logOperateApi := v1.ApiGroupApp.LogApiGroup.SysLogOperateApi
	{
		logOperateRouter.GET("", logOperateApi.GetSysLogOperateList)       // 获取操作日志列表
		logOperateRouter.GET("/:id", logOperateApi.GetSysLogOperate)       // 获取操作日志详情
		logOperateRouter.POST("", logOperateApi.CreateSysLogOperate)       // 创建操作日志
		logOperateRouter.DELETE("/:id", logOperateApi.DeleteSysLogOperate) // 删除操作日志
	}
}

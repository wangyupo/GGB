package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/system"
)

func SysDictDataRouter(r *gin.RouterGroup) {
	sysDictDataRouter := r.Group("/system/dictData")
	{
		sysDictDataRouter.GET("", system.GetSysDictDataList)
		sysDictDataRouter.POST("", system.CreateSysDictData)
		sysDictDataRouter.PUT("/:id", system.UpdateSysDictData)
		sysDictDataRouter.DELETE("/:id", system.DeleteSysDictData)
	}
}

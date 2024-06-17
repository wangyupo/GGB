package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/system"
)

type DictDataRouter struct{}

func (s *DictDataRouter) InitDictDataRouter(Router *gin.RouterGroup) {
	sysDictDataRouter := Router.Group("/system/dictData")
	{
		sysDictDataRouter.GET("", system.GetSysDictDataList)
		sysDictDataRouter.POST("", system.CreateSysDictData)
		sysDictDataRouter.PUT("/:id", system.UpdateSysDictData)
		sysDictDataRouter.DELETE("/:id", system.DeleteSysDictData)
	}
}

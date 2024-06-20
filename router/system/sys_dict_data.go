package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api"
)

type DictDataRouter struct{}

func (s *DictDataRouter) InitDictDataRouter(Router *gin.RouterGroup) {
	sysDictDataRouter := Router.Group("/system/dictData")
	SysDictDataApi := api.ApiGroupApp.SysApiGroup.SysDictDataApi
	{
		sysDictDataRouter.GET("", SysDictDataApi.GetSysDictDataList)
		sysDictDataRouter.POST("", SysDictDataApi.CreateSysDictData)
		sysDictDataRouter.PUT("/:id", SysDictDataApi.UpdateSysDictData)
		sysDictDataRouter.DELETE("/:id", SysDictDataApi.DeleteSysDictData)
	}
}

package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/v1"
	"github.com/wangyupo/GGB/middleware"
)

type DictDataRouter struct{}

func (s *DictDataRouter) InitDictDataRouter(Router *gin.RouterGroup) {
	dictDataRouter := Router.Group("/system/dictData").Use(middleware.OperationRecord())
	dictDataRouterWithoutRecord := Router.Group("/system/dictData")
	dictDataApi := v1.ApiGroupApp.SysApiGroup.SysDictDataApi
	{
		dictDataRouter.POST("", dictDataApi.CreateSysDictData)       // 新增字典数据
		dictDataRouter.PUT("/:id", dictDataApi.UpdateSysDictData)    // 编辑字典数据
		dictDataRouter.DELETE("/:id", dictDataApi.DeleteSysDictData) // 删除字典数据
	}
	{
		dictDataRouterWithoutRecord.GET("", dictDataApi.GetSysDictDataList) // 获取字典数据列表
	}
}

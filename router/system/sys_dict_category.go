package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/v1"
	"github.com/wangyupo/GGB/middleware"
)

type DictCategoryRouter struct{}

func (s *DictCategoryRouter) InitDictCategoryRouter(Router *gin.RouterGroup) {
	dictCategoryRouter := Router.Group("/system/dictCategory").Use(middleware.OperationRecord())
	dictCategoryRouterWithoutRecord := Router.Group("/system/dictCategory")
	dictCategoryApi := v1.ApiGroupApp.SysApiGroup.SysDictCategoryApi
	{
		dictCategoryRouter.POST("", dictCategoryApi.CreateSysDictCategory)       // 创建字典类型
		dictCategoryRouter.PUT("/:id", dictCategoryApi.UpdateSysDictCategory)    // 更新字典类型
		dictCategoryRouter.DELETE("/:id", dictCategoryApi.DeleteSysDictCategory) // 删除字典类型
	}
	{
		dictCategoryRouterWithoutRecord.GET("", dictCategoryApi.GetSysDictCategoryList) // 获取字典类型列表
		dictCategoryRouterWithoutRecord.GET("/:id", dictCategoryApi.GetSysDictCategory) // 获取字典类型详情
	}
}

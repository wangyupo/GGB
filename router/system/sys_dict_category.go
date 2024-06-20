package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api"
)

type DictCategoryRouter struct{}

func (s *DictCategoryRouter) InitDictCategoryRouter(Router *gin.RouterGroup) {
	sysDictCategoryRouter := Router.Group("/system/dictCategory")
	sysDictCategoryApi := api.ApiGroupApp.SysApiGroup.SysDictCategoryApi
	{
		sysDictCategoryRouter.GET("", sysDictCategoryApi.GetSysDictCategoryList)
		sysDictCategoryRouter.POST("", sysDictCategoryApi.CreateSysDictCategory)
		sysDictCategoryRouter.PUT("/:id", sysDictCategoryApi.UpdateSysDictCategory)
		sysDictCategoryRouter.DELETE("/:id", sysDictCategoryApi.DeleteSysDictCategory)
	}
}

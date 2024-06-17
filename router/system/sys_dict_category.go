package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/system"
)

type DictCategoryRouter struct{}

func (s *DictCategoryRouter) InitDictCategoryRouter(Router *gin.RouterGroup) {
	sysDictCategoryRouter := Router.Group("/system/dictCategory")
	{
		sysDictCategoryRouter.GET("", system.GetSysDictCategoryList)
		sysDictCategoryRouter.POST("", system.CreateSysDictCategory)
		sysDictCategoryRouter.PUT("/:id", system.UpdateSysDictCategory)
		sysDictCategoryRouter.DELETE("/:id", system.DeleteSysDictCategory)
	}
}

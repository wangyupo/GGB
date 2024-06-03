package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/system"
)

func SysDictCategoryRouter(r *gin.RouterGroup) {
	sysDictCategoryRouter := r.Group("/system/dictCategory")
	{
		sysDictCategoryRouter.GET("", system.GetSysDictCategoryList)
		sysDictCategoryRouter.POST("", system.CreateSysDictCategory)
		sysDictCategoryRouter.PUT("/:id", system.UpdateSysDictCategory)
		sysDictCategoryRouter.DELETE("/:id", system.DeleteSysDictCategory)
	}
}

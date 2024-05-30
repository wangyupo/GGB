package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/system"
)

func SysUserRouter(r *gin.RouterGroup) {
	sysUserRouter := r.Group("/sysUser")
	{
		sysUserRouter.GET("", system.GetSystemUserList)
		sysUserRouter.GET("/:id", system.GetSystemUser)
		sysUserRouter.POST("", system.CreateSystemUser)
		sysUserRouter.PUT("/:id", system.UpdateSystemUser)
		sysUserRouter.DELETE("/:id", system.DeleteSystemUser)
	}
}

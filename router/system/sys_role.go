package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/system"
)

func SysRoleRouter(r *gin.RouterGroup) {
	sysRoleRouter := r.Group("/system/role")
	{
		sysRoleRouter.GET("", system.GetSysRoleList)
		sysRoleRouter.POST("", system.CreateSysRole)
		sysRoleRouter.PUT("/:id", system.UpdateSysRole)
		sysRoleRouter.DELETE("/:id", system.DeleteSysRole)
	}
}

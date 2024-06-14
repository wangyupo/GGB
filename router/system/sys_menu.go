package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/system"
)

func SysMenuRouter(r *gin.RouterGroup) {
	sysMenuRouter := r.Group("/system/menu")
	{
		sysMenuRouter.GET("", system.GetSysMenuList)
		sysMenuRouter.POST("", system.CreateSysMenu)
		sysMenuRouter.PUT("/:id", system.UpdateSysMenu)
		sysMenuRouter.DELETE("/:id", system.DeleteSysMenu)

		sysMenuRouter.PUT("/move", system.MoveSysMenu)
	}
}

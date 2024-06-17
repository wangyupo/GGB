package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/system"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	sysMenuRouter := Router.Group("/system/menu")
	{
		sysMenuRouter.GET("", system.GetSysMenuList)
		sysMenuRouter.POST("", system.CreateSysMenu)
		sysMenuRouter.PUT("/:id", system.UpdateSysMenu)
		sysMenuRouter.DELETE("/:id", system.DeleteSysMenu)

		sysMenuRouter.PUT("/move", system.MoveSysMenu)
	}
}

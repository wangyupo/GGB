package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	sysMenuRouter := Router.Group("/system/menu")
	sysMenuApi := api.ApiGroupApp.SysApiGroup.SysMenuApi
	{
		sysMenuRouter.GET("", sysMenuApi.GetSysMenuList)
		sysMenuRouter.POST("", sysMenuApi.CreateSysMenu)
		sysMenuRouter.PUT("/:id", sysMenuApi.UpdateSysMenu)
		sysMenuRouter.DELETE("/:id", sysMenuApi.DeleteSysMenu)

		sysMenuRouter.PUT("/move", sysMenuApi.MoveSysMenu)
	}
}

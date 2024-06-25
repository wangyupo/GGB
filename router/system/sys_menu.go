package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api"
	"github.com/wangyupo/GGB/middleware"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("/system/menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("/system/menu")
	menuApi := api.ApiGroupApp.SysApiGroup.SysMenuApi
	{
		menuRouter.POST("", menuApi.CreateSysMenu)       // 新建菜单
		menuRouter.PUT("/:id", menuApi.UpdateSysMenu)    // 编辑菜单
		menuRouter.DELETE("/:id", menuApi.DeleteSysMenu) // 删除菜单
		menuRouter.PUT("/move", menuApi.MoveSysMenu)     // 移动菜单
	}
	{
		menuRouterWithoutRecord.GET("", menuApi.GetSysMenuList) // 获取菜单列表
	}
}

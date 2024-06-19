package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api"
)

type RoleRouter struct{}

func (s *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	sysRoleRouter := Router.Group("/system/role")
	sysRoleApi := api.ApiGroupApp.SysApiGroup.SysRoleApi
	{
		sysRoleRouter.GET("", sysRoleApi.GetSysRoleList)
		sysRoleRouter.POST("", sysRoleApi.CreateSysRole)
		sysRoleRouter.PUT("/:id", sysRoleApi.UpdateSysRole)
		sysRoleRouter.DELETE("/:id", sysRoleApi.DeleteSysRole)

		sysRoleRouter.PATCH("/:id/status", sysRoleApi.ChangeRoleStatus)
		sysRoleRouter.POST("/menu", sysRoleApi.RoleAssignMenu)
		sysRoleRouter.POST("/user", sysRoleApi.RoleAssignUser)
		sysRoleRouter.DELETE("/user", sysRoleApi.RoleUnAssignUser)
		sysRoleRouter.GET("/user", sysRoleApi.GetUserByRole)
		sysRoleRouter.GET("/:id/menu", sysRoleApi.GetMenuByRole)
	}
}

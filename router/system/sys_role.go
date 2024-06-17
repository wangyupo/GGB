package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/system"
)

type RoleRouter struct{}

func (s *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	sysRoleRouter := Router.Group("/system/role")
	{
		sysRoleRouter.GET("", system.GetSysRoleList)
		sysRoleRouter.POST("", system.CreateSysRole)
		sysRoleRouter.PUT("/:id", system.UpdateSysRole)
		sysRoleRouter.DELETE("/:id", system.DeleteSysRole)

		sysRoleRouter.PATCH("/:id/status", system.ChangeRoleStatus)
		sysRoleRouter.POST("/menu", system.RoleAssignMenu)
		sysRoleRouter.POST("/user", system.RoleAssignUser)
		sysRoleRouter.DELETE("/user", system.RoleUnAssignUser)
		sysRoleRouter.GET("/user", system.GetUserByRole)
		sysRoleRouter.GET("/:id/menu", system.GetMenuByRole)
	}
}

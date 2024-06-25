package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/v1"
	"github.com/wangyupo/GGB/middleware"
)

type RoleRouter struct{}

func (s *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleRouter := Router.Group("/system/role").Use(middleware.OperationRecord())
	roleRouterWithoutRecord := Router.Group("/system/role")
	roleApi := v1.ApiGroupApp.SysApiGroup.SysRoleApi
	{
		roleRouter.POST("", roleApi.CreateSysRole)                // 新建角色
		roleRouter.PUT("/:id", roleApi.UpdateSysRole)             // 编辑角色
		roleRouter.DELETE("/:id", roleApi.DeleteSysRole)          // 删除角色
		roleRouter.PATCH("/:id/status", roleApi.ChangeRoleStatus) // 修改角色状态
		roleRouter.POST("/menu", roleApi.RoleAssignMenu)          // 角色授权菜单
		roleRouter.POST("/user", roleApi.RoleAssignUser)          // 角色授权用户
		roleRouter.DELETE("/user", roleApi.RoleUnAssignUser)      // 角色取消授权用户
	}
	{
		roleRouterWithoutRecord.GET("", roleApi.GetSysRoleList)         // 获取角色列表
		roleRouterWithoutRecord.GET("/user", roleApi.GetUserByRole)     // 获取角色用户
		roleRouterWithoutRecord.GET("/:id/menu", roleApi.GetMenuByRole) // 获取角色权限菜单
	}
}

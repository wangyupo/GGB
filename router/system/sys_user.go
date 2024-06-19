package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	sysUserRouter := Router.Group("/system/user")
	sysUserApi := api.ApiGroupApp.SysApiGroup.SysUserApi
	{
		sysUserRouter.GET("", sysUserApi.GetSystemUserList)
		sysUserRouter.GET("/:id", sysUserApi.GetSystemUser)
		sysUserRouter.POST("", sysUserApi.CreateSystemUser)
		sysUserRouter.PUT("/:id", sysUserApi.UpdateSystemUser)
		sysUserRouter.DELETE("/:id", sysUserApi.DeleteSystemUser)

		sysUserRouter.PATCH("/password", sysUserApi.ChangePassword)
		sysUserRouter.PATCH("/:id/reset-password", sysUserApi.ResetPassword)
		sysUserRouter.PATCH("/:id/status", sysUserApi.ChangeSystemUserStatus)
	}
}

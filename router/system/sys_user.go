package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/system"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	sysUserRouter := Router.Group("/system/user")
	{
		sysUserRouter.GET("", system.GetSystemUserList)
		sysUserRouter.GET("/:id", system.GetSystemUser)
		sysUserRouter.POST("", system.CreateSystemUser)
		sysUserRouter.PUT("/:id", system.UpdateSystemUser)
		sysUserRouter.DELETE("/:id", system.DeleteSystemUser)

		sysUserRouter.PATCH("/password", system.ChangePassword)
		sysUserRouter.PATCH("/:id/reset-password", system.ResetPassword)
		sysUserRouter.PATCH("/:id/status", system.ChangeSystemUserStatus)
	}
}

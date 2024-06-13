package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/system"
)

func SysUserRouter(r *gin.RouterGroup) {
	sysUserRouter := r.Group("/system/user/info")
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

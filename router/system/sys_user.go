package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/v1"
	"github.com/wangyupo/GGB/middleware"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("/system/user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("/system/user")
	userApi := v1.ApiGroupApp.SysApiGroup.SysUserApi
	{
		userRouter.POST("", userApi.CreateSystemUser)                   // 新建用户
		userRouter.PUT("/:id", userApi.UpdateSystemUser)                // 编辑用户
		userRouter.DELETE("/:id", userApi.DeleteSystemUser)             // 删除用户
		userRouter.PATCH("/password", userApi.ChangePassword)           // 修改用户自身密码
		userRouter.PATCH("/:id/reset-password", userApi.ResetPassword)  // 重置用户密码
		userRouter.PATCH("/:id/status", userApi.ChangeSystemUserStatus) // 修改用户状态
	}
	{
		userRouterWithoutRecord.GET("", userApi.GetSystemUserList) // 获取用户列表
		userRouterWithoutRecord.GET("/:id", userApi.GetSystemUser) // 获取用户详情
	}
}

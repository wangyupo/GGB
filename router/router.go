package router

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/wangyupo/GGB/api/system"
	"github.com/wangyupo/GGB/middleware"
	middlewareLog "github.com/wangyupo/GGB/middleware/log"
	"github.com/wangyupo/GGB/router/log"
	"github.com/wangyupo/GGB/router/system"
	"os"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	ApiGroup := r.Group(os.Getenv("ROUTE_PREFIX"))
	{
		ApiGroup.POST("/login", apiSystem.Login) // 登录
	}

	ApiGroup.Use(middleware.Jwt())
	{
		ApiGroup.POST("/logout", middlewareLog.LoginLog(2), apiSystem.Logout) // 登出

		system.SysUserRouter(ApiGroup)
		system.SysRoleRouter(ApiGroup)
		system.SysMenuRouter(ApiGroup)
		system.SysDictCategoryRouter(ApiGroup)
		system.SysDictDataRouter(ApiGroup)

		log.LoginRouter(ApiGroup)
	}

	return r
}

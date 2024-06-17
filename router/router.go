package router

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/wangyupo/GGB/api/system"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/middleware"
	middlewareLog "github.com/wangyupo/GGB/middleware/log"
	"github.com/wangyupo/GGB/router/log"
	"github.com/wangyupo/GGB/router/system"
	"go.uber.org/zap"
)

func InitRouter() *gin.Engine {
	// 初始化zap日志
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	Router := gin.Default()

	// 路由-不做鉴权
	PublicGroup := Router.Group(global.GGB_CONFIG.System.RouterPrefix)
	{
		PublicGroup.POST("/login", apiSystem.Login) // 登录
	}

	// 路由-需要鉴权
	PrivateGroup := Router.Group(global.GGB_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.Logger(logger)).Use(middleware.Jwt())
	{
		PrivateGroup.POST("/logout", middlewareLog.LoginLog(2), apiSystem.Logout) // 登出

		system.SysUserRouter(PrivateGroup)         // 用户管理
		system.SysRoleRouter(PrivateGroup)         // 角色管理
		system.SysMenuRouter(PrivateGroup)         // 菜单管理
		system.SysDictCategoryRouter(PrivateGroup) // 字典类型管理
		system.SysDictDataRouter(PrivateGroup)     // 字典数据管理

		log.LoginRouter(PrivateGroup) // 登录日志
	}

	return Router
}

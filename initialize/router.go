package initialize

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/wangyupo/GGB/api/system"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/middleware"
	middlewareLog "github.com/wangyupo/GGB/middleware/log"
	"github.com/wangyupo/GGB/router"
	"go.uber.org/zap"
)

// Routers 注册路由
func Routers() *gin.Engine {
	// 初始化zap日志
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	Router := gin.Default()

	systemRouter := router.RouterGroupApp.System
	logRouter := router.RouterGroupApp.Log

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

		systemRouter.InitUserRouter(PrivateGroup)         // 用户管理
		systemRouter.InitRoleRouter(PrivateGroup)         // 角色管理
		systemRouter.InitMenuRouter(PrivateGroup)         // 菜单管理
		systemRouter.InitDictCategoryRouter(PrivateGroup) // 字典类型管理
		systemRouter.InitDictDataRouter(PrivateGroup)     // 字典数据管理

		logRouter.InitLoginLogRouter(PrivateGroup) // 登录日志
	}

	return Router
}

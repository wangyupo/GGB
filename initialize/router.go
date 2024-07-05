package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/api/v1"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/middleware"
	"github.com/wangyupo/GGB/router"
	"go.uber.org/zap"
)

// Routers 注册路由
func Routers() *gin.Engine {
	// 初始化zap日志
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	Router := gin.Default()

	commonRouter := router.RouterGroupApp.Common
	systemRouter := router.RouterGroupApp.System
	logRouter := router.RouterGroupApp.Log
	sysUserApi := v1.ApiGroupApp.SysApiGroup.SysUserApi

	// 路由-不做鉴权
	PublicGroup := Router.Group(global.GGB_CONFIG.System.RouterPrefix)
	{
		PublicGroup.POST("/login", sysUserApi.Login) // 登录

		PublicGroup.GET("/captcha", sysUserApi.GetCaptcha)            // 获取图形验证码
		PublicGroup.POST("/captcha/verify", sysUserApi.VerifyCaptcha) // 校验图形验证码
	}

	// 路由-需要鉴权
	PrivateGroup := Router.Group(global.GGB_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.Logger(logger)).Use(middleware.Jwt())
	{
		PrivateGroup.POST("/logout", sysUserApi.Logout) // 登出

		commonRouter.InitUploadFileRouter(PrivateGroup) // 上传文件
		commonRouter.InitTranscriptRouter(PrivateGroup) // Excel导入/导出

		systemRouter.InitUserRouter(PrivateGroup)         // 用户管理
		systemRouter.InitRoleRouter(PrivateGroup)         // 角色管理
		systemRouter.InitMenuRouter(PrivateGroup)         // 菜单管理
		systemRouter.InitDictCategoryRouter(PrivateGroup) // 字典类型管理
		systemRouter.InitDictDataRouter(PrivateGroup)     // 字典数据管理

		logRouter.InitLoginLogRouter(PrivateGroup)      // 登录日志
		logRouter.InitSysLogOperateRouter(PrivateGroup) // 操作日志
	}

	return Router
}

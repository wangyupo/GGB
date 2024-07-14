package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "github.com/wangyupo/GGB/api/v1"
	_ "github.com/wangyupo/GGB/docs"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/middleware"
	"github.com/wangyupo/GGB/router"
)

// Routers 注册路由
func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	commonRouter := router.RouterGroupApp.Common
	systemRouter := router.RouterGroupApp.System
	logRouter := router.RouterGroupApp.Log
	sysUserApi := v1.ApiGroupApp.SysApiGroup.SysUserApi

	// 匹配 swagger 路由（启动后端服务后，访问地址：服务地址:端口/swagger/index.html）
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 路由-不做鉴权
	PublicGroup := Router.Group(global.GGB_CONFIG.System.RouterPrefix)
	{
		PublicGroup.POST("/login", sysUserApi.Login) // 登录

		PublicGroup.POST("/captcha", sysUserApi.GetCaptcha)           // 获取图形验证码
		PublicGroup.POST("/captcha/verify", sysUserApi.VerifyCaptcha) // 校验图形验证码
	}

	// 路由-需要鉴权
	PrivateGroup := Router.Group(global.GGB_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.Jwt())
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

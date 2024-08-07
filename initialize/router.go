package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "github.com/wangyupo/GGB/api/v1"
	_ "github.com/wangyupo/GGB/docs"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/middleware"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/router"
)

// Routers 注册路由
func Routers() *gin.Engine {
	Router := gin.New(func(engine *gin.Engine) {
		engine.HandleMethodNotAllowed = true // 开启方法不匹配规则校验
	})
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	commonRouter := router.RouterGroupApp.Common
	systemRouter := router.RouterGroupApp.System
	logRouter := router.RouterGroupApp.Log
	sysBaseApi := v1.ApiGroupApp.SysApiGroup.SysBaseApi

	// 匹配 swagger 路由（启动后端服务后，访问地址：服务地址:端口/swagger/index.html）
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 路由-不做鉴权
	PublicGroup := Router.Group(global.GGB_CONFIG.System.RouterPrefix)
	{
		systemRouter.InitBaseRouter(PublicGroup)
	}

	// 路由-需要鉴权
	PrivateGroup := Router.Group(global.GGB_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.Jwt())
	{
		PrivateGroup.POST("/logout", sysBaseApi.Logout) // 登出

		commonRouter.InitEmailRouter(PrivateGroup)      // 发送邮件
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

	// 自定义404响应
	Router.NoRoute(func(c *gin.Context) {
		response.NotFound(c)
	})

	// 自定义方法不匹配响应
	Router.NoMethod(func(c *gin.Context) {
		response.MethodNotAllowed(c)
	})

	return Router
}

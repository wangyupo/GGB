package router

import (
	"github.com/gin-gonic/gin"
	apiSystem "github.com/wangyupo/GGB/api/system"
	"github.com/wangyupo/GGB/middleware"
	"github.com/wangyupo/GGB/router/system"
	"os"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	ApiGroup := r.Group(os.Getenv("ROUTE_PREFIX"))
	{
		ApiGroup.GET("", apiSystem.Login)
	}

	ApiGroup.Use(middleware.Jwt())
	{
		system.SysUserRouter(ApiGroup)
	}

	return r
}

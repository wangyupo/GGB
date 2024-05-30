package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/router/system"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	system.SysUserRouter(r)

	return r
}

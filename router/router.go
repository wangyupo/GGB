package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/gin-cli/router/system"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	system.SysUserRouter(r)

	return r
}

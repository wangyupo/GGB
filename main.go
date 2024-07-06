package main

import (
	"github.com/wangyupo/GGB/core"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title 						GGB Swagger API接口文档
// @version 					v0.0.1
// @description 				基于 Gin 框架构建的高效、可扩展的后端服务架构，专为现代 Web 应用而设计。
// @BasePath 					/api
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
func main() {
	global.GGB_VP = core.Viper()       // 初始化Viper
	global.GGB_LOG = core.Zap()        // 初始化日志库
	zap.ReplaceGlobals(global.GGB_LOG) // 替换全局的日志记录器，可以在程序的任何地方通过 zap.L() 函数来获取这个全局日志记录器，并进行日志记录
	global.GGB_DB = initialize.Gorm()  // 初始化gorm，连接数据库
	initialize.Timer()                 // 初始化定时任务
	initialize.Validator()             // 初始化校验器
	if global.GGB_DB != nil {
		initialize.RegisterTables() // 初始化表
		db, _ := global.GGB_DB.DB() // 程序结束前关闭数据库链接
		defer db.Close()
	}

	core.RunWindowsServer()
}

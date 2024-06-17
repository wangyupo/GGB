package main

import (
	"github.com/wangyupo/GGB/core"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	global.GGB_VP = core.Viper()      // 初始化Viper
	global.GGB_DB = initialize.Gorm() // gorm连接数据库
	if global.GGB_DB != nil {
		initialize.RegisterTables()
		// 程序结束前关闭数据库链接
		db, _ := global.GGB_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}

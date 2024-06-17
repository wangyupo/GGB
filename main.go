package main

import (
	"fmt"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/initialize"
	"github.com/wangyupo/GGB/migration"
	"github.com/wangyupo/GGB/router"
	"os"
)

func init() {
	initialize.LoadEnvVariable()
	initialize.ConnectDB()
}

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	r := router.InitRouter()

	if global.GGB_DB != nil {
		migration.Migrate()
		fmt.Print(123)
		// 程序结束前关闭数据库链接
		db, _ := global.GGB_DB.DB()
		defer db.Close()
	}

	r.Run(os.Getenv("SERVER_PORT"))
}

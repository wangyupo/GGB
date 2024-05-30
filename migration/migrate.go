package migration

import (
	"fmt"
	"github.com/wangyupo/gin-cli/global"
	"github.com/wangyupo/gin-cli/model/system"
)

// Migrate 表和数据迁移
func Migrate() {
	err := global.DB.AutoMigrate(system.SysUser{})
	if err != nil {
		fmt.Print(err)
	}
}

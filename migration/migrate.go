package migration

import (
	"fmt"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
)

// Migrate 表和数据迁移
func Migrate() {
	err := global.DB.AutoMigrate(system.SysUser{})
	if err != nil {
		fmt.Print(err)
	}
}

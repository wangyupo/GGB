package initialize

import (
	"github.com/wangyupo/GGB/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func GormMysql() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       global.GGB_CONFIG.Mysql.Dsn(),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 日志级别
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名（启用前：sys_users；启用后：sys_user）
		},
	})
	if err != nil {
		global.GGB_LOG.Error("数据库链接错误", zap.Error(err))
		panic(err)
	}

	return db
}

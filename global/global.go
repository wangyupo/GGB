package global

import "gorm.io/gorm"

var (
	GGB_DB               *gorm.DB
	DefaultLoginPassword = "123456"
)

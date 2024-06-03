package global

import "gorm.io/gorm"

var (
	DB                   *gorm.DB
	DefaultLoginPassword = "123456"
)

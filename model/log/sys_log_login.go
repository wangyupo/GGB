package log

import (
	"github.com/wangyupo/GGB/enums"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
	"gorm.io/gorm"
)

type SysLogLogin struct {
	global.BaseModel
	UserId    uint            `json:"userId" gorm:"comment:用户ID"`               // 用户ID
	Type      enums.LoginType `json:"type" gorm:"size:64;comment:操作类型 0登出 1登入"` // 操作类型 0登出 1登入
	TypeText  string          `json:"typeText" gorm:"_"`
	IP        string          `json:"ip" gorm:"size:128;comment:请求ip"`   // 请求ip
	UserAgent string          `json:"userAgent" gorm:"comment:用户设备和浏览器"` // 用户设备和浏览器
	User      system.SysUser  `json:"user" gorm:"foreignKey:UserId;references:ID"`
	UserName  string          `json:"userName" gorm:"-"` //  gorm:"-" 用于告诉 GORM 字段只在应用层面使用，不需要持久化到数据库。这意味着这个字段不会被映射到数据库表中的任何列上。
}

func (s *SysLogLogin) AfterFind(tx *gorm.DB) (err error) {
	s.TypeText = s.Type.Text()
	s.UserName = s.User.UserName
	return
}

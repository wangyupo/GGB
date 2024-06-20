package log

import (
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
)

type SysLogLogin struct {
	global.BaseModel
	UserId    uint           `json:"userId" gorm:"comment:用户ID"`
	Type      uint           `json:"type" gorm:"size:64;comment:操作类型 1登入 2登出"`
	IP        string         `json:"ip" gorm:"size:128;comment:请求ip"`
	UserAgent string         `json:"userAgent" gorm:"comment:用户设备和浏览器"`
	User      system.SysUser `json:"user" gorm:"foreignKey:UserId;references:ID"`
}

package system

import "github.com/wangyupo/GGB/global"

// SysUser 系统用户
type SysUser struct {
	global.BaseModel
	UserName string `json:"userName" gorm:"size:64;index;unique;not null;comment:用户登录名"`
	NickName string `json:"nickName" gorm:"size:64;default:系统用户;comment:用户昵称"`
	Email    string `json:"email" gorm:"size:128;comment:用户邮箱"`
	Password string `json:"-" gorm:"comment:密码"`
	Status   int    `json:"status" gorm:"default:1;comment:用户状态 1正常 2冻结"`
}

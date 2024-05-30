package system

import "github.com/wangyupo/gin-cli/global"

type SysUser struct {
	global.BaseModel
	UserName string `json:"userName" gorm:"index;comment:用户登录名"`
	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	Email    string `json:"email" gorm:"comment:用户邮箱"`
	Status   int    `json:"status" gorm:"default:1;comment:用户状态 1正常 2冻结"`
}

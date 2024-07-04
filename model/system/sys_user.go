package system

import "github.com/wangyupo/GGB/global"

// SysUser 系统用户
type SysUser struct {
	global.BaseModel
	UserName string     `json:"userName" form:"userName" gorm:"type:varchar(64);index;unique;not null;comment:用户登录名"`
	NickName string     `json:"nickName" gorm:"type:varchar(64);default:系统用户;comment:用户昵称"`
	Email    string     `json:"email" gorm:"type:varchar(128);comment:用户邮箱"`
	Password string     `json:"-" gorm:"comment:密码"`
	Status   uint       `json:"status" gorm:"type:tinyint(1);default:0;comment:用户状态 0禁用 1启用"`
	Roles    []*SysRole `json:"roles" gorm:"many2many:sys_role_user;"`
}

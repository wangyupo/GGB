package system

import "github.com/wangyupo/GGB/global"

// SysRoleUser 系统角色和系统用户关联
type SysRoleUser struct {
	global.BaseModel
	UserID uint `json:"userId" gorm:"comment:用户ID"`
	RoleID uint `json:"roleId" gorm:"comment:角色ID"`
}

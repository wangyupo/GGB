package system

import "github.com/wangyupo/GGB/global"

// SysRoleUser 系统角色和系统用户关联
type SysRoleUser struct {
	global.BaseModel
	SysUserID uint `json:"sysUserId" gorm:"comment:用户ID"`
	SysRoleID uint `json:"sysRoleId" gorm:"comment:角色ID"`
}

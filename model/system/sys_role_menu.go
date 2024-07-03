package system

import "github.com/wangyupo/GGB/global"

// SysRoleMenu 系统角色和系统菜单关联
type SysRoleMenu struct {
	global.BaseModel
	SysRoleID uint `json:"sysRoleId" gorm:"comment:角色ID"`
	SysMenuID uint `json:"sysMenuId" gorm:"comment:菜单ID"`
}

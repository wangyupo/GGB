package system

import "github.com/wangyupo/GGB/global"

// SysRoleMenu 系统角色和系统菜单关联
type SysRoleMenu struct {
	global.BaseModel
	RoleID uint `json:"roleId" gorm:"comment:角色ID"`
	MenuID uint `json:"menuId" gorm:"comment:菜单ID"`
}

package request

import "github.com/wangyupo/GGB/model/system"

type SysRoleQuery struct {
	system.SysRole
}

type ChangeRoleStatus struct {
	Status int `json:"status" binding:"required"`
}

type RoleAssignMenu struct {
	SysRoleID  uint   `json:"sysRoleId" binding:"required"`
	SysMenuIds []uint `json:"sysMenuIds"`
}

type RoleAssignUser struct {
	SysRoleID  uint   `json:"sysRoleId" binding:"required"`
	SysUserIds []uint `json:"sysUserIds"`
}

package request

import "github.com/wangyupo/GGB/model/system"

type SysRoleQuery struct {
	system.SysRole
}

type ChangeRoleStatus struct {
	Status int `json:"status"`
}

type RoleAssignMenu struct {
	SysRoleID  uint   `json:"sysRoleId"`
	SysMenuIds []uint `json:"sysMenuIds"`
}

type RoleAssignUser struct {
	SysRoleID  uint   `json:"sysRoleId"`
	SysUserIds []uint `json:"sysUserIds"`
}

package request

import "github.com/wangyupo/GGB/model/system"

type SysRoleQuery struct {
	system.SysRole
}

type ChangeRoleStatus struct {
	Status int `json:"status"`
}

type RoleAssignMenu struct {
	RoleID  uint   `json:"roleId"`
	MenuIds []uint `json:"menuIds"`
}

type RoleAssignUser struct {
	RoleID  uint   `json:"roleId"`
	UserIds []uint `json:"userIds"`
}

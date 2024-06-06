package request

type ChangeRoleStatus struct {
	ID     uint `json:"id"`
	Status int  `json:"status"`
}

type RoleAssignMenu struct {
	RoleID  uint   `json:"roleId"`
	MenuIds []uint `json:"menuIds"`
}

type RoleAssignUser struct {
	RoleID  uint   `json:"roleId"`
	UserIds []uint `json:"userIds"`
}

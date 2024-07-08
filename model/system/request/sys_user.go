package request

import "github.com/wangyupo/GGB/model/system"

type ChangePassword struct {
	Password    string `json:"password" binding:"required"`    // 密码
	NewPassword string `json:"newPassword" binding:"required"` // 新密码
}

type ChangeSystemUserStatus struct {
	Status int `json:"status" binding:"required,oneof=0 1"`
}

type SystemUserList struct {
	system.SysUser
}

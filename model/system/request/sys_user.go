package request

import "github.com/wangyupo/GGB/model/system"

type Login struct {
	UserName string `json:"userName" binding:"required,min=2,max=10"` // 用户名
	Password string `json:"password" binding:"required,min=2,max=18"` // 密码
}

type ChangePassword struct {
	Password    string `json:"password" binding:"required"`    // 密码
	NewPassword string `json:"newPassword" binding:"required"` // 新密码
}

type ChangeSystemUserStatus struct {
	Status int `json:"status" binding:"required"`
}

type SystemUserList struct {
	system.SysUser
}

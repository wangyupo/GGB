package request

import "github.com/wangyupo/GGB/model/system"

type Login struct {
	UserName string `json:"userName"` // 用户名
	Password string `json:"password"` // 密码
}

type ChangePassword struct {
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type ChangeSystemUserStatus struct {
	Status int `json:"status"`
}

type SystemUserList struct {
	system.SysUser
}

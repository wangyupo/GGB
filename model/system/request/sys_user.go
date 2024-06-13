package request

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

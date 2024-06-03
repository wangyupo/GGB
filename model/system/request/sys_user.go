package request

type Login struct {
	UserName string `json:"userName"` // 用户名
	Password string `json:"password"` // 密码
}

type ChangePassword struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

package request

// Login 登录
type Login struct {
	UserName string `json:"userName" binding:"required,min=2,max=10"` // 用户名
	Password string `json:"password" binding:"required,min=2,max=18"` // 密码
}

// CaptchaRequest 获取验证码
type CaptchaRequest struct {
	CaptchaType string `json:"captchaType" binding:"required"` // 验证码类型
}

// Captcha 校验验证码
type Captcha struct {
	CaptchaId string `json:"captchaId" binding:"required"`           // 验证码ID
	Captcha   string `json:"captcha" binding:"required,min=5,max=5"` // 验证码
}

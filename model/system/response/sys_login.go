package response

import "github.com/wangyupo/GGB/model/system"

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type CaptchaResponse struct {
	ID     string `json:"id"`
	Base64 string `json:"base_64"`
}

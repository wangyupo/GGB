package response

import "github.com/wangyupo/GGB/model/system"

type LoginLogResponse struct {
	system.SysLogLogin
	UserName string `json:"userName"`
}

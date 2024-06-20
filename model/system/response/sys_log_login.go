package response

import (
	"github.com/wangyupo/GGB/model/log"
)

type LoginLogResponse struct {
	log.SysLogLogin
	UserName string `json:"userName"`
}

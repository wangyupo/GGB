package v1

import (
	"github.com/wangyupo/GGB/api/v1/log"
	"github.com/wangyupo/GGB/api/v1/system"
)

type ApiGroup struct {
	SysApiGroup system.ApiGroup
	LogApiGroup log.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

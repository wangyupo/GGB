package api

import (
	"github.com/wangyupo/GGB/api/log"
	"github.com/wangyupo/GGB/api/system"
)

type ApiGroup struct {
	SysApiGroup    system.ApiGroup
	SysLogApiGroup log.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

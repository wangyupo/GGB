package api

import "github.com/wangyupo/GGB/api/system"

type ApiGroup struct {
	SysApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

package service

import (
	"github.com/wangyupo/GGB/service/common"
	"github.com/wangyupo/GGB/service/log"
	"github.com/wangyupo/GGB/service/system"
)

type ServiceGroup struct {
	CommonService      common.ServiceGroup
	SystemServiceGroup system.ServiceGroup
	LogServiceGroup    log.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

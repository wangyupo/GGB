package service

import (
	"github.com/wangyupo/GGB/service/log"
	"github.com/wangyupo/GGB/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	LogServiceGroup    log.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

package log

import "github.com/wangyupo/GGB/service"

type ApiGroup struct {
	SysLogLoginApi
	SysLogOperateApi
}

var (
	sysLoginLogService   = service.ServiceGroupApp.LogServiceGroup.SysLogLoginService
	sysLogOperateService = service.ServiceGroupApp.LogServiceGroup.SysLogOperateService
)

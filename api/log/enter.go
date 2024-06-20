package log

import "github.com/wangyupo/GGB/service"

type ApiGroup struct {
	SysLoginLogApi
}

var (
	sysLoginLogService = service.ServiceGroupApp.SystemLogGroup.SysLoginLogService
)

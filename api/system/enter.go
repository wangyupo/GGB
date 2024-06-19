package system

import "github.com/wangyupo/GGB/service"

type ApiGroup struct {
	SysUserApi
	SysRoleApi
}

var (
	sysUserService = service.ServiceGroupApp.SystemServiceGroup.SysUserService
	sysRoleService = service.ServiceGroupApp.SystemServiceGroup.SysRoleService
)

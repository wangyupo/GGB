package system

import "github.com/wangyupo/GGB/service"

type ApiGroup struct {
	SysUserApi
	SysRoleApi
	SysMenuApi
	SysDictCategoryApi
	SysDictDataApi
}

var (
	sysUserService         = service.ServiceGroupApp.SystemServiceGroup.SysUserService
	sysRoleService         = service.ServiceGroupApp.SystemServiceGroup.SysRoleService
	sysMenuService         = service.ServiceGroupApp.SystemServiceGroup.SysMenuService
	sysDictCategoryService = service.ServiceGroupApp.SystemServiceGroup.SysDictCategoryService
	sysDictDataService     = service.ServiceGroupApp.SystemServiceGroup.SysDictDataService
)

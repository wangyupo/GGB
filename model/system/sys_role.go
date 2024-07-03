package system

import "github.com/wangyupo/GGB/global"

// SysRole 系统角色
type SysRole struct {
	global.BaseModel
	RoleName    string     `json:"roleName" form:"roleName" gorm:"type:varchar(64);unique;index;size:128;comment:角色名，用于显示"`
	RoleCode    string     `json:"roleCode" gorm:"type:varchar(64);unique;comment:角色的唯一标识码"`
	Description string     `json:"description" gorm:"type:varchar(128);comment:角色描述等信息"`
	Status      uint       `json:"status" gorm:"type:tinyint(1);default:0;comment:角色状态 0禁用 1启用"`
	Users       []*SysUser `gorm:"many2many:sys_role_user;"`
}

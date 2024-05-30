package system

import "github.com/wangyupo/GGB/global"

// SysRole 系统角色
type SysRole struct {
	global.BaseModel
	RoleName    string `json:"roleName" gorm:"unique;index;size:128;comment:角色名，用于显示"`
	RoleCode    string `json:"roleCode" gorm:"unique;comment:角色的唯一标识码"`
	Description string `json:"description" gorm:"comment:角色描述等信息"`
	Status      uint   `json:"status" gorm:"default:1;comment:角色状态 1启用 2禁用"`
}

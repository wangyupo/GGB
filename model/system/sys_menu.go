package system

import "github.com/wangyupo/GGB/global"

// SysMenu 系统菜单
type SysMenu struct {
	global.BaseModel
	Label    string `json:"label" gorm:"size:128;not null;comment:菜单名称"`
	Path     string `json:"path" gorm:"comment:菜单路径"`
	Icon     string `json:"icon" gorm:"comment:菜单图标"`
	ParentId uint   `json:"parentId" gorm:"default:0;comment:父级菜单ID"`
	Sort     int    `json:"sort" gorm:"default:0;comment:排序号（越小越靠前）"`
	Type     uint   `json:"type" gorm:"default:1;comment:菜单类型 1菜单 2页面"`
}

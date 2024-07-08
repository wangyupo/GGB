package system

import "github.com/wangyupo/GGB/global"

// SysMenu 系统菜单
type SysMenu struct {
	global.BaseModel
	Label    string     `json:"label" gorm:"type:varchar(64);not null;comment:菜单名称"`        // 菜单名称
	Path     string     `json:"path" gorm:"type:varchar(128);comment:菜单路径"`                 // 菜单路径
	Icon     string     `json:"icon" gorm:"type:varchar(64);comment:菜单图标"`                  // 菜单图标
	ParentId uint       `json:"parentId" gorm:"default:0;comment:父级菜单ID"`                   // 父级菜单ID
	Sort     int        `json:"sort" gorm:"type:int(3);default:0;comment:排序号（越小越靠前）"`       // 排序号（越小越靠前）
	Type     uint       `json:"type" gorm:"type:tinyint(1);default:0;comment:菜单类型 0菜单 1页面"` // 菜单类型 0菜单 1页面
	Roles    []*SysRole `json:"-" gorm:"many2many:sys_role_menu"`
}

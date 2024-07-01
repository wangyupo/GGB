package system

import "github.com/wangyupo/GGB/global"

// SysDictCategory 系统字典
type SysDictCategory struct {
	global.BaseModel
	Label       string `json:"label" form:"label" gorm:"type:varchar(64);comment:分类名"`
	LabelCode   string `json:"labelCode" gorm:"type:varchar(64);comment:分类编码"`
	Description string `json:"description" gorm:"type:varchar(255);comment:分类描述"`
}

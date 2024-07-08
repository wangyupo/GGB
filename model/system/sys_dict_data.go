package system

import "github.com/wangyupo/GGB/global"

// SysDictData 系统字典数据
type SysDictData struct {
	global.BaseModel
	CategoryID  uint   `json:"categoryId" gorm:"index"`
	Label       string `json:"label" form:"label" gorm:"type:varchar(64);comment:字典键"` // 字典键
	Value       string `json:"value" gorm:"type:varchar(64);comment:字典值"`              // 字典值
	Description string `json:"description" gorm:"type:varchar(255);comment:字典项描述"`     // 字典项描述
}

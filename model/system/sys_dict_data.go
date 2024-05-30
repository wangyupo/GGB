package system

import "github.com/wangyupo/GGB/global"

// SysDictData 系统字典数据
type SysDictData struct {
	global.BaseModel
	CategoryID  uint   `json:"categoryID" gorm:"index"`
	Label       string `json:"label" gorm:"size:128;comment:字典键"`
	Value       string `json:"value" gorm:"size:255;comment:字典值"`
	Description string `json:"description" gorm:"size:255;comment:字典项描述"`
}

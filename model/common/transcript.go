package common

import "github.com/wangyupo/GGB/global"

type Transcript struct {
	global.BaseModel
	Name      string `json:"name" form:"name" gorm:"type:varchar(64);comment:学生姓名"`
	Language  uint   `json:"language" gorm:"type:int(3);comment:语文"`
	Math      uint   `json:"math" gorm:"type:int(3);comment:数学"`
	English   uint   `json:"english" gorm:"type:int(3);comment:英语"`
	Geography uint   `json:"geography" gorm:"type:int(3);comment:地理"`
	Politics  uint   `json:"politics" gorm:"type:int(3);comment:政治"`
	UserId    uint   `json:"userId" gorm:"comment:上传Excel的用户id"`
}

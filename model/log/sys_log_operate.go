package log

import (
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
	"gorm.io/gorm"
	"time"
)

type SysLogOperate struct {
	global.BaseModel
	Ip           string         `json:"ip" form:"ip" gorm:"comment:请求ip"`
	Method       string         `json:"method" form:"method" gorm:"comment:请求方法"`
	Path         string         `json:"path" form:"path" gorm:"comment:请求路径"`
	Status       int            `json:"status" form:"status" gorm:"comment:请求状态"`
	Latency      time.Duration  `json:"latency" form:"latency" gorm:"comment:延迟" swaggertype:"integer"`
	Agent        string         `json:"agent" form:"agent" gorm:"comment:代理"`
	ErrorMessage string         `json:"errorMessage" form:"errorMessage" gorm:"comment:错误信息"`
	Body         string         `json:"body" form:"body" gorm:"comment:请求Body"`
	Response     string         `json:"response" form:"response" gorm:"响应Body"`
	UserID       uint           `json:"userId" form:"userId" gorm:"comment:用户id"`
	User         system.SysUser `json:"user"`
	UserName     string         `json:"userName" gorm:"-"`
}

func (s *SysLogOperate) AfterFind(tx *gorm.DB) (err error) {
	s.UserName = s.User.UserName
	return
}

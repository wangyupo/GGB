package request

import (
	"github.com/wangyupo/GGB/model/log"
)

type SysLogOperateQuery struct {
	log.SysLogOperate
	StartDate string `json:"startDate" form:"startDate"`
	EndDate   string `json:"endDate" form:"endDate"`
}

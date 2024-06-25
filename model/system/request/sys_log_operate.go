package request

import "github.com/wangyupo/GGB/model/system"

type SysLogOperateQuery struct {
	system.SysLogOperate
	StartDate string `json:"startDate" form:"startDate"`
	EndDate   string `json:"endDate" form:"endDate"`
}

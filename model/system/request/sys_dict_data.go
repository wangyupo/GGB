package request

import "github.com/wangyupo/GGB/model/system"

type SysDictDataQuery struct {
	system.SysDictData
	CategoryId uint `json:"categoryId" form:"categoryId"`
}

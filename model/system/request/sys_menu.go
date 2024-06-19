package request

import "github.com/wangyupo/GGB/model/system"

type SysMenuQuery struct {
	system.SysMenu
}

type MoveMenu struct {
	OriginID uint   `json:"originId"`
	TargetID uint   `json:"targetId"`
	DropType string `json:"dropType"`
}

package request

import "github.com/wangyupo/GGB/model/system"

type SysMenuQuery struct {
	system.SysMenu
}

type MoveMenu struct {
	OriginID uint   `json:"originId" binding:"required"`
	TargetID uint   `json:"targetId" binding:"required"`
	DropType string `json:"dropType" binding:"required"`
}

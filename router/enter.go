package router

import (
	"github.com/wangyupo/GGB/router/common"
	"github.com/wangyupo/GGB/router/log"
	"github.com/wangyupo/GGB/router/system"
)

type RouterGroup struct {
	Common common.RouterGroup
	System system.RouterGroup
	Log    log.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

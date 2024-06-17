package router

import (
	"github.com/wangyupo/GGB/router/log"
	"github.com/wangyupo/GGB/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	Log    log.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

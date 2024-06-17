package core

import (
	"fmt"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/router"
	"net/http"
	"time"
)

func RunWindowsServer() {
	Router := router.InitRouter()
	address := fmt.Sprintf(":%s", global.GGB_CONFIG.System.Addr)

	server := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    20 * time.Second, // 读取请求头的超时时间
		WriteTimeout:   20 * time.Second, // 写入响应的超时时间
		MaxHeaderBytes: 1 << 20,          // 请求头的最大字节数，1 MB
	}

	// 使用http.Server 监听和服务
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

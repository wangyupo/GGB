package core

import (
	"fmt"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/initialize"
	"net/http"
	"time"
)

// RunWindowsServer 运行http服务
func RunWindowsServer() {
	if global.GGB_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	Router := initialize.Routers()
	address := fmt.Sprintf(":%s", global.GGB_CONFIG.System.Addr)

	server := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    20 * time.Second, // 读取请求头的超时时间
		WriteTimeout:   20 * time.Second, // 写入响应的超时时间
		MaxHeaderBytes: 1 << 20,          // 请求头的最大字节数，1 MB
	}

	fmt.Printf(`
ggb_server 服务启动成功
项目地址：		https://github.com/wangyupo/GGB
OpenAPI 文档地址：http://127.0.0.1%s/swagger/index.html
	`, address)

	// 使用http.Server 监听和服务
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/utils"
	"io"
)

// RequestLog 记录请求日志中间件
func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 读取请求数据
		path := c.Request.URL.Path
		method := c.Request.Method
		queryParams := c.Request.URL.RawQuery
		userId, _ := utils.GetUserID(c)

		bodyBytes, _ := io.ReadAll(c.Request.Body) // Body 的内容被消耗掉，读取到了内存

		requestBody := string(bodyBytes)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 重新设置 c.Request.Body，以便请求体可以在后续的处理中再次读取

		// 记录响应数据
		responseWriter := gin.ResponseWriter()
	}
}

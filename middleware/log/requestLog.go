package log

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// RequestLog 记录请求日志中间件
func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求数据
		//path := c.Request.URL.Path            // 获取请求路径
		//method := c.Request.Method            // 获取请求方式
		//queryParams := c.Request.URL.RawQuery // 获取请求参数
		//clientIP := c.ClientIP()              // 获取客户端IP
		//userId, _ := utils.GetUserID(c)       // 从token获取用户id
		//
		//// 读取请求体
		//bodyBytes, _ := io.ReadAll(c.Request.Body)                // Body 的内容被消耗掉，读取到了内存
		//c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 重新设置 c.Request.Body，以便请求体可以在后续的处理中再次读取
		//requestBody := string(bodyBytes)
		//
		//// 设置自定义响应写入器
		//writer := &CustomResponseWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		//c.Writer = writer
		//
		//// 处理请求
		//c.Next()
		//
		//// 获取响应数据
		//responseBody := writer.Body.String()
		//statusCode := c.Writer.Status()
	}
}

// 清理一个月前的日志（硬删除）
func cleanupOldLogs() {
	// 计算一个月前的时间
}

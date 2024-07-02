package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/log"
	"github.com/wangyupo/GGB/service"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var sysLogOperateService = service.ServiceGroupApp.LogServiceGroup.SysLogOperateService

// 定义全局变量用于响应池和缓冲区大小
var respPool sync.Pool
var bufferSize = 1024

// 初始化响应池
func init() {
	respPool.New = func() interface{} {
		return make([]byte, bufferSize)
	}
}

// OperationRecord 记录操作日志的中间件
func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 定义变量以保存请求体
		var body []byte
		// 读取请求体数据
		err := getRequestBody(c, &body)
		if err != nil {
			global.GGB_LOG.Error("read body from request error:", zap.Error(err))
		}

		// 获取用户ID
		userId, _ := utils.GetUserID(c)

		// 构造操作记录
		record := log.SysLogOperate{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.RequestURI,
			Agent:  c.Request.UserAgent(),
			Body:   getBodyContent(body, c.GetHeader("Content-Type")),
			UserID: userId,
		}

		// 创建响应写入器对象，并替换原有的响应写入器
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		// 记录当前时间，用于计算请求处理时间
		now := time.Now()

		// 处理请求
		c.Next()

		// 设置操作记录的延迟时间
		record.Latency = time.Since(now)
		// 设置操作记录的错误信息
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		// 设置操作记录的返回状态码
		record.Status = c.Writer.Status()
		// 设置操作记录的响应内容
		record.Response = writer.body.String()

		// 判断是否需要截断响应内容
		if shouldTruncateResponse(c.Writer.Header()) {
			if len(record.Response) > bufferSize {
				record.Response = "[超出记录长度]"
			}
		}

		// 将操作记录保存到数据库中
		if err := sysLogOperateService.CreateSysLogOperate(record); err != nil {
			global.GGB_LOG.Error("create operation record error:", zap.Error(err))
		}
	}
}

// 从请求上下文中读取请求体数据
func getRequestBody(c *gin.Context, body *[]byte) error {
	if c.Request.Method != http.MethodGet {
		var err error
		*body, err = io.ReadAll(c.Request.Body)
		if err != nil {
			return err
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(*body))
	} else {
		query := c.Request.URL.RawQuery
		query, _ = url.QueryUnescape(query)
		split := strings.Split(query, "&")
		m := make(map[string]string)
		for _, v := range split {
			kv := strings.Split(v, "=")
			if len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}
		*body, _ = json.Marshal(&m)
	}
	return nil
}

// 根据请求体和内容类型返回请求体内容,必要时会截断请求体内容
func getBodyContent(body []byte, contentType string) string {
	if strings.Contains(contentType, "multipart/form-data") {
		return "[文件]"
	}
	if len(body) > bufferSize {
		return "[超出记录长度]"
	}
	return string(body)
}

// 判断是否应该截断响应内容，通常用于文件下载等内容较大的响应情况
func shouldTruncateResponse(header http.Header) bool {
	disposition := header.Get("Content-Disposition")
	contentType := header.Get("Content-Type")
	return strings.Contains(header.Get("Pragma"), "public") ||
		strings.Contains(header.Get("Expires"), "0") ||
		strings.Contains(header.Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") ||
		strings.Contains(contentType, "application/force-download") ||
		strings.Contains(contentType, "application/octet-stream") ||
		strings.Contains(contentType, "application/vnd.ms-excel") ||
		strings.Contains(contentType, "application/download") ||
		strings.Contains(disposition, "attachment") ||
		strings.Contains(header.Get("Content-Transfer-Encoding"), "binary")
}

// responseBodyWriter 是自定义的响应写入器，用于捕获响应内容
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write 实现自定义的写入方法，捕获并写入响应内容
func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

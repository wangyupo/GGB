package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
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

var respPool sync.Pool
var bufferSize = 1024

func init() {
	respPool.New = func() interface{} {
		return make([]byte, bufferSize)
	}
}

// OperationRecord 记录操作日志的中间件
func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		err := getRequestBody(c, &body)
		if err != nil {
			global.GGB_LOG.Error("read body from request error:", zap.Error(err))
		}

		userId, _ := utils.GetUserID(c)

		record := system.SysLogOperate{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   getBodyContent(body, c.GetHeader("Content-Type")),
			UserID: userId,
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		record.Latency = time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Response = writer.body.String()

		if shouldTruncateResponse(c.Writer.Header()) {
			if len(record.Response) > bufferSize {
				record.Response = "[超出记录长度]"
			}
		}

		if err := sysLogOperateService.CreateSysLogOperate(record); err != nil {
			global.GGB_LOG.Error("create operation record error:", zap.Error(err))
		}
	}
}

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

func getBodyContent(body []byte, contentType string) string {
	if strings.Contains(contentType, "multipart/form-data") {
		return "[文件]"
	}
	if len(body) > bufferSize {
		return "[超出记录长度]"
	}
	return string(body)
}

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

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

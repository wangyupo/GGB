package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Logger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 记录请求结束时间和处理时长
		duration := time.Since(start)

		// 记录请求的详细信息
		logger.Info("HTTP Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", duration),
			zap.String("clientIP", c.ClientIP()),
			zap.String("userAgent", c.Request.UserAgent()),
			zap.String("referer", c.Request.Referer()),
		)

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				logger.Error("HTTP Request Error",
					zap.String("error", e),
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.Int("status", c.Writer.Status()),
					zap.Duration("duration", duration),
					zap.String("clientIP", c.ClientIP()),
					zap.String("userAgent", c.Request.UserAgent()),
					zap.String("referer", c.Request.Referer()),
				)
			}
		}
	}
}

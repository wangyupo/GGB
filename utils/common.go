package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetQueryInt 提取和验证整数参数
func GetQueryInt(c *gin.Context, key string, defaultValue int) int {
	valueStr := c.DefaultQuery(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func GetToken(c *gin.Context) string {
	token, _ := c.Cookie("Authorization")
	if token == "" {
		token = c.Request.Header.Get("Authorization")
	}
	return token
}

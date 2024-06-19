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

// Str2uint 字符串转uint
func Str2uint(str string) (uint, error) {
	param, err := strconv.ParseUint(str, 10, 64)
	return uint(param), err
}

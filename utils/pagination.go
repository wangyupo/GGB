package utils

import "github.com/gin-gonic/gin"

// GetPaginationParams 获取分页参数
func GetPaginationParams(c *gin.Context) (int, int) {
	// 你可以在这里修改下参数名，例如page和limit
	pageNumber := GetQueryInt(c, "pageNumber", 1)
	pageSize := GetQueryInt(c, "pageSize", 10)
	return pageNumber, pageSize
}

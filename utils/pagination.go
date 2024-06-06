package utils

import "github.com/gin-gonic/gin"

// GetPaginationParams 获取分页参数
func GetPaginationParams(c *gin.Context) (int, int) {
	// 你可以在这里修改下参数名，例如page和limit
	offset := GetQueryInt(c, "pageNumber", 1)
	limit := GetQueryInt(c, "pageSize", 10)
	offset = (offset - 1) * limit
	return offset, limit
}

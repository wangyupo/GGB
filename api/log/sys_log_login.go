package log

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/utils"
)

type SysLoginLogApi struct{}

// GetLoginLogList 列表
func (s *SysLoginLogApi) GetLoginLogList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	userId, _ := utils.Str2uint(c.Query("userId"))

	list, total, err := sysLoginLogService.GetLoginLogList(userId, offset, limit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

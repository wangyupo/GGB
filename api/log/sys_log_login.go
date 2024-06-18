package log

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/utils"
)

// GetLoginLogList 列表
func GetLoginLogList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	userId := c.Query("userId")

	// 声明 log.SysLogLogin 类型的变量以存储查询结果
	loginLogList := make([]system.SysLogLogin, 0)
	var total int64

	// 准备数据库查询
	db := global.GGB_DB.Model(&system.SysLogLogin{})
	if userId != "" {
		db = db.Where("user_id = ?", "%"+userId+"%")
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取分页数据
	err := db.Offset(offset).Limit(limit).Preload("User").Find(&loginLogList).Error
	if err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 结果集增加 userName 字段
	results := make([]map[string]interface{}, len(loginLogList))
	for i, log := range loginLogList {
		logMap, _ := utils.ExcludeNestedFields(log, []string{"User"})
		// 添加用户名
		logMap["userName"] = log.User.UserName
		results[i] = logMap
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  results,
		Total: total,
	}, c)
}

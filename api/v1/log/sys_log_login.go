package log

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
	"reflect"
	"time"
)

type SysLogLoginApi struct{}

// GetSysLogLoginList 获取登录日志列表
func (s *SysLogLoginApi) GetSysLogLoginList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	userId, _ := utils.Str2uint(c.Query("userId"))

	list, total, err := sysLoginLogService.GetSysLogLoginList(userId, offset, limit)
	if err != nil {
		global.GGB_LOG.Error("获取登录日志列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// ExportExcel 导出Excel
func (s *SysLogLoginApi) ExportExcel(c *gin.Context) {
	userId, _ := utils.Str2uint(c.Query("userId"))

	list, _, err := sysLoginLogService.GetSysLogLoginList(userId, 1, 999)
	if err != nil {
		global.GGB_LOG.Error("获取登录日志列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	var ExcelList [][]interface{}
	ExcelList = append(ExcelList, []interface{}{"用户名称", "ip地址", "操作类型", "记录时间"})

	ColumnKey := []string{"userName", "ip", "type", "createdAt"}
	for _, listItem := range list.([]map[string]interface{}) {
		var listValueItem []interface{}
		for _, key := range ColumnKey {
			value := listItem[key]
			if reflect.TypeOf(listItem[key]) == reflect.TypeOf(time.Time{}) {
				value = value.(time.Time).Format(time.DateTime)
			}
			listValueItem = append(listValueItem, value)
		}
		ExcelList = append(ExcelList, listValueItem)
	}

	filePath, _ := utils.ExportExcelByTemplate(ExcelList)

	// 使用Gin的c.File方法直接提供该文件进行下载
	c.File(filePath)
}

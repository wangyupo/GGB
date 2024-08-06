package log

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/log"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
	"time"
)

type SysLogLoginApi struct{}

// GetSysLogLoginList
//
//	@Tags		SysLogLogin
//	@Summary	获取登录日志列表
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		pageNumber	query		int																	true	"分页"
//	@Param		pageSize	query		int																	true	"每页条数"
//	@Success	200			{object}	response.Response{data=response.PageResult{list=[]log.SysLogLogin}}	"返回列表，总数"
//	@Router		/system/log/login [GET]
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

// ExportExcel
//
//	@Tags		SysLogLogin
//	@Summary	导出Excel
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	octet-stream
//	@Success	200	{file}	file	"返回文件流（blob）"
//	@Router		/system/log/login/excel [GET]
func (s *SysLogLoginApi) ExportExcel(c *gin.Context) {
	userId, _ := utils.Str2uint(c.Query("userId"))

	list, _, err := sysLoginLogService.GetSysLogLoginList(userId, 0, 999)
	if err != nil {
		global.GGB_LOG.Error("获取登录日志列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	var ExcelList [][]interface{}
	ExcelList = append(ExcelList, []interface{}{"用户名称", "ip地址", "操作类型", "记录时间"})
	for _, row := range list.([]log.SysLogLogin) {
		var listValueItem []interface{}
		listValueItem = append(listValueItem, row.UserName, row.IP, row.TypeText, row.CreatedAt.Format(time.DateTime))
		ExcelList = append(ExcelList, listValueItem)
	}

	filePath, err := utils.CreateExcelByList(ExcelList, "Sheet1")
	if err != nil {
		global.GGB_LOG.Error("导出登录日志Excel失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 使用Gin的c.File方法直接提供该文件进行下载
	c.File(filePath)
}

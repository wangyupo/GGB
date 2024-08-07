package log

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/log"
	"github.com/wangyupo/GGB/model/system/request"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
)

type SysLogOperateApi struct{}

// GetSysLogOperateList
//
//	@Tags		SysLogOperate
//	@Summary	获取系统操作日志列表
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		pageNumber	query		int																		true	"分页"
//	@Param		pageSize	query		int																		true	"每页条数"
//	@Success	200			{object}	response.Response{data=response.PageResult{list=[]log.SysLogOperate}}	"返回列表，总数"
//	@Router		/log/operate [GET]
func (s *SysLogOperateApi) GetSysLogOperateList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	var query request.SysLogOperateQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := sysLogOperateService.GetSysLogOperateList(query, offset, limit)
	if err != nil {
		global.GGB_LOG.Error("查询系统操作日志列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// CreateSysLogOperate 新建系统操作日志
func (s *SysLogOperateApi) CreateSysLogOperate(c *gin.Context) {
	// 声明 system.SysLogOperate 类型的变量以存储 JSON 数据
	var req log.SysLogOperate

	// 绑定 JSON 请求体中的数据到 SysLogOperate 结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		// 错误处理
		utils.HandleValidatorError(err, c)
		return
	}

	err := sysLogOperateService.CreateSysLogOperate(req)
	if err != nil {
		global.GGB_LOG.Error("新建系统操作日志失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// GetSysLogOperate
//
//	@Tags		SysLogOperate
//	@Summary	获取系统操作日志详情
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id	path		int											true	"操作日志id（logOperateId）"
//	@Success	200	{object}	response.Response{data=log.SysLogOperate}	"返回操作日志详情"
//	@Router		/log/operate/:id [GET]
func (s *SysLogOperateApi) GetSysLogOperate(c *gin.Context) {
	// 获取路径参数
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	id, _ := utils.Str2uint(c.Param("id"))

	sysLogOperate, err := sysLogOperateService.GetSysLogOperate(id)
	if err != nil {
		global.GGB_LOG.Error("获取系统操作日志详情失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(sysLogOperate, c)
}

// DeleteSysLogOperate
//
//	@Tags		SysLogOperate
//	@Summary	删除系统操作日志
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id	path		int						true	"操作日志id（logOperateId）"
//	@Success	200	{object}	response.MsgResponse	"返回操作成功提示"
//	@Router		/log/operate/:id [DELETE]
func (s *SysLogOperateApi) DeleteSysLogOperate(c *gin.Context) {
	// 获取路径参数
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	id, _ := utils.Str2uint(c.Param("id"))

	err := sysLogOperateService.DeleteSysLogOperate(id)
	if err != nil {
		global.GGB_LOG.Error("删除系统操作日志失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

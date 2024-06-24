package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
)

type SysDictDataApi struct{}

// GetSysDictDataList 列表
func (s *SysDictDataApi) GetSysDictDataList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	var query request.SysDictDataQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := sysDictDataService.GetSysDictDataList(query, offset, limit)
	if err != nil {
		global.GGB_LOG.Error("查询字典列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// CreateSysDictData 新建
func (s *SysDictDataApi) CreateSysDictData(c *gin.Context) {
	// 声明 system.SysDictData 类型的变量以存储 JSON 数据
	var req system.SysDictData

	// 绑定 JSON 请求体中的数据到 sysDictData 结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := sysDictDataService.CreateSysDictData(req)
	if err != nil {
		global.GGB_LOG.Error("新建字典失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// GetSysDictData 详情
func (s *SysDictDataApi) GetSysDictData(c *gin.Context) {
	// 获取路径参数
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	id, _ := utils.Str2uint(c.Param("id"))

	sysDictData, err := sysDictDataService.GetSysDictData(id)
	if err != nil {
		global.GGB_LOG.Error("获取字典详情失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(sysDictData, c)
}

// UpdateSysDictData 编辑
func (s *SysDictDataApi) UpdateSysDictData(c *gin.Context) {
	// 获取路径参数
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	id, _ := utils.Str2uint(c.Param("id"))

	// 声明 system.SysDictData 类型的变量以存储查询结果
	var req system.SysDictData
	// 绑定请求参数到数据对象
	if err := c.ShouldBindJSON(&req); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := sysDictDataService.UpdateSysDictData(req, id)
	if err != nil {
		global.GGB_LOG.Error("编辑字典失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// DeleteSysDictData 删除
func (s *SysDictDataApi) DeleteSysDictData(c *gin.Context) {
	// 获取路径参数
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	id, _ := utils.Str2uint(c.Param("id"))

	err := sysDictDataService.DeleteSysDictData(id)
	if err != nil {
		global.GGB_LOG.Error("删除字典失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

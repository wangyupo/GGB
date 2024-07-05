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

type SysDictCategoryApi struct{}

// GetSysDictCategoryList 查询字典类型列表
func (s *SysDictCategoryApi) GetSysDictCategoryList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	var query request.SysDictCategoryQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := sysDictCategoryService.GetSysDictCategoryList(query, offset, limit)
	if err != nil {
		global.GGB_LOG.Error("查询字典类型列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// CreateSysDictCategory 新建字典类型
func (s *SysDictCategoryApi) CreateSysDictCategory(c *gin.Context) {
	// 声明 system.SysDictCategory 类型的变量以存储 JSON 数据
	var req system.SysDictCategory

	// 绑定 JSON 请求体中的数据到 SysDictCategory 结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		// 错误处理
		utils.HandleValidatorError(err, c)
		return
	}

	err := sysDictCategoryService.CreateSysDictCategory(req)
	if err != nil {
		global.GGB_LOG.Error("新建字典类型失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// GetSysDictCategory 获取字典类型详情
func (s *SysDictCategoryApi) GetSysDictCategory(c *gin.Context) {
	// 获取路径参数
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	id, _ := utils.Str2uint(c.Param("id"))

	sysDictCategory, err := sysDictCategoryService.GetSysDictCategory(id)
	if err != nil {
		global.GGB_LOG.Error("获取字典类型详情失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(sysDictCategory, c)
}

// UpdateSysDictCategory 编辑字典类型
func (s *SysDictCategoryApi) UpdateSysDictCategory(c *gin.Context) {
	// 获取路径参数
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	id, _ := utils.Str2uint(c.Param("id"))

	// 声明 system.SysDictCategory 类型的变量以存储查询结果
	var req system.SysDictCategory
	// 绑定请求参数到数据对象
	if err := c.ShouldBindJSON(&req); err != nil {
		// 错误处理
		utils.HandleValidatorError(err, c)
		return
	}

	err := sysDictCategoryService.UpdateSysDictCategory(req, id)
	if err != nil {
		global.GGB_LOG.Error("编辑字典类型失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// DeleteSysDictCategory 删除字典类型
func (s *SysDictCategoryApi) DeleteSysDictCategory(c *gin.Context) {
	// 获取路径参数
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	id, _ := utils.Str2uint(c.Param("id"))

	err := sysDictCategoryService.DeleteSysDictCategory(id)
	if err != nil {
		global.GGB_LOG.Error("删除字典类型失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

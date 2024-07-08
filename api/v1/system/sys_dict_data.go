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

// GetSysDictDataList
// @Tags      SysDictData
// @Summary   获取字典数据列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param	  pageNumber 		query 	int 	true 	"分页"
// @Param	  pageSize  		query 	int 	true 	"每页条数"
// @Success   200   {object}  	response.Response{data=response.PageResult}  "返回列表，总数"
// @Router    /system/dictData [GET]
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

// CreateSysDictData
// @Tags      SysDictData
// @Summary   新建字典数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data	body		system.SysDictData	true  	"SysDictData模型"
// @Success   200   {object}  	response.MsgResponse  		"返回操作成功提示"
// @Router    /system/dictData [POST]
func (s *SysDictDataApi) CreateSysDictData(c *gin.Context) {
	// 声明 system.SysDictData 类型的变量以存储 JSON 数据
	var req system.SysDictData

	// 绑定 JSON 请求体中的数据到 SysDictData 结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		// 错误处理
		utils.HandleValidatorError(err, c)
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

// GetSysDictData
// @Tags      SysDictData
// @Summary   获取字典数据详情
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param	  id  	path 		int 			true 					"字典数据id（dictDataId）"
// @Success   200   {object}  	response.Response{data=system.SysRole}  "返回字典数据详情"
// @Router    /system/dictData/:id [GET]
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

// UpdateSysDictData
// @Tags      SysDictData
// @Summary   编辑字典数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param	  id  	path 		int 				true 	"字典数据id（dictDataId）"
// @Param	  data  body 		system.SysDictData 	true 	"SysDictData模型"
// @Success   200   {object}  	response.MsgResponse  		"返回操作成功提示"
// @Router    /system/dictData/:id [PUT]
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
		utils.HandleValidatorError(err, c)
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

// DeleteSysDictData
// @Tags      SysDictData
// @Summary   删除字典数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param	  id  	path 		int		true 			"字典数据id（dictDataId）"
// @Success   200   {object}  	response.MsgResponse  	"返回操作成功提示"
// @Router    /system/dictData/:id [DELETE]
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

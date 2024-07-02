package common

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common"
	"github.com/wangyupo/GGB/model/common/request"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
)

type TranscriptApi struct{}

// DownloadExcelTemplate 下载Excel模板
func (t *TranscriptApi) DownloadExcelTemplate(c *gin.Context) {
	filePath := global.GGB_CONFIG.Excel.TemplateDir + "Excel导入模板.xlsx"
	c.File(filePath)
}

// ImportByExcel 通过Excel导入数据
func (t *TranscriptApi) ImportByExcel(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		global.GGB_LOG.Error("接收文件失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId, _ := utils.GetUserID(c)
	err = transcriptService.ImportByExcel(file, userId)
	if err != nil {
		global.GGB_LOG.Error("文件上传失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithMessage("数据导入成功", c)
}

// ExportExcel 导出Excel
func (t *TranscriptApi) ExportExcel(c *gin.Context) {
	// 获取其它查询参数
	var query request.TranscriptQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, _, err := transcriptService.GetTranscriptList(query, 0, 999)
	if err != nil {
		global.GGB_LOG.Error("查询成绩列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	var ExcelList [][]interface{}
	// 组装数据
	titleRow := []interface{}{"姓名", "语文", "数学", "英语", "地理", "政治"}
	ExcelList = append(ExcelList, titleRow)
	for _, row := range list.([]common.Transcript) {
		var rowValues []interface{}
		rowValues = append(rowValues, row.Name, row.Language, row.Math, row.English, row.Geography, row.Politics)
		ExcelList = append(ExcelList, rowValues)
	}

	// 生成Excel并返回路径
	filePath, _ := utils.CreateExcelByList(ExcelList)
	c.File(filePath)
}

// GetTranscriptList 查询成绩列表
func (t *TranscriptApi) GetTranscriptList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	var query request.TranscriptQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := transcriptService.GetTranscriptList(query, offset, limit)
	if err != nil {
		global.GGB_LOG.Error("查询成绩列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// CreateTranscript 新建成绩
func (t *TranscriptApi) CreateTranscript(c *gin.Context) {
	// 声明 common.Transcript 类型的变量以存储 JSON 数据
	var req common.Transcript

	// 绑定 JSON 请求体中的数据到 Transcript 结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := transcriptService.CreateTranscript(req)
	if err != nil {
		global.GGB_LOG.Error("新建成绩失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// DeleteTranscript 删除成绩
func (t *TranscriptApi) DeleteTranscript(c *gin.Context) {
	// 获取路径参数
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	id, _ := utils.Str2uint(c.Param("id"))

	err := transcriptService.DeleteTranscript(id)
	if err != nil {
		global.GGB_LOG.Error("删除成绩失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

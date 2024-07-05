package common

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common"
	"github.com/wangyupo/GGB/model/common/request"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/utils"
	"github.com/xuri/excelize/v2"
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
	// 1-获取其它查询参数
	var query request.TranscriptQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 2-获取列表
	list, _, err := transcriptService.GetTranscriptList(query, 0, 999)
	if err != nil {
		global.GGB_LOG.Error("查询成绩列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 3-组装数据
	var ExcelList [][]interface{}
	titleRow := []interface{}{"姓名", "语文", "数学", "英语", "地理", "政治", "总分"}
	ExcelList = append(ExcelList, titleRow)
	for _, row := range list.([]common.Transcript) {
		var rowValues []interface{}
		rowValues = append(rowValues, row.Name, row.Language, row.Math, row.English, row.Geography, row.Politics)
		ExcelList = append(ExcelList, rowValues)
	}

	// 4-新建excel并填充数据
	f, err := utils.ExtraExcelAfterList(ExcelList, "Sheet1")

	// 5-插入图表
	_ = f.AddChart("Sheet1", "I5", &excelize.Chart{
		Type: excelize.Col, // 柱状图
		Title: []excelize.RichTextRun{
			{
				Text: "成绩单", // 图表标题
			},
		},
		Series: []excelize.ChartSeries{
			{
				Name:       "Sheet1!$B$1",      // 柱的名称
				Categories: "Sheet1!$A$2:$A$4", // 柱代表分类
				Values:     "Sheet1!$B$2:$B$4", // 柱对应的值
			},
			{
				Name:       "Sheet1!$C$1",
				Categories: "Sheet1!$A$2:$A$4",
				Values:     "Sheet1!$C$2:$C$4",
			},
			{
				Name:       "Sheet1!$D$1",
				Categories: "Sheet1!$A$2:$A$4",
				Values:     "Sheet1!$D$2:$D$4",
			},
		},
		PlotArea: excelize.ChartPlotArea{
			ShowVal: true, // 显示数据标签
		},
	})

	// 6-保存文件
	filePath, err := utils.SaveExcelByExcelize(f)
	if err != nil {
		global.GGB_LOG.Error("导出成绩列表Excel失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 7-返回文件流
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
		utils.HandleValidatorError(err, c)
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

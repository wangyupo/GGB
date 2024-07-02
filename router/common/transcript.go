package common

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wangyupo/GGB/api/v1"
	"github.com/wangyupo/GGB/middleware"
)

type TranscriptRouter struct{}

func (t *TranscriptRouter) InitTranscriptRouter(Router *gin.RouterGroup) {
	transcriptRouter := Router.Group("/common/excel").Use(middleware.OperationRecord())
	transcriptRouterWithoutRecord := Router.Group("/common/excel")
	transcriptApi := v1.ApiGroupApp.CommonApiGroup.TranscriptApi
	{
		transcriptRouter.GET("/template", transcriptApi.DownloadExcelTemplate) // 下载Excel模板
		transcriptRouter.POST("", transcriptApi.ImportByExcel)                 // 通过Excel导入数据
	}
	{
		transcriptRouterWithoutRecord.GET("", transcriptApi.GetTranscriptList)  // 获取文件列表
		transcriptRouterWithoutRecord.GET("/export", transcriptApi.ExportExcel) // 导出Excel
	}
}

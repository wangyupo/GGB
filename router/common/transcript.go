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
		transcriptRouter.POST("/import", transcriptApi.ImportByExcel)          // 通过Excel导入数据
		transcriptRouter.POST("", transcriptApi.CreateTranscript)              // 新建成绩
		transcriptRouter.DELETE("/:id", transcriptApi.DeleteTranscript)        // 删除成绩
	}
	{
		transcriptRouterWithoutRecord.GET("", transcriptApi.GetTranscriptList)  // 获取成绩列表
		transcriptRouterWithoutRecord.GET("/export", transcriptApi.ExportExcel) // 导出Excel
	}
}

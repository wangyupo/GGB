package common

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wangyupo/GGB/api/v1"
	"github.com/wangyupo/GGB/middleware"
)

type UploadFileRouter struct{}

func (u *UploadFileRouter) InitUploadFileRouter(Router *gin.RouterGroup) {
	uploadFileRouter := Router.Group("/common/upload").Use(middleware.OperationRecord())
	uploadFileRouterWithoutRecord := Router.Group("/common/upload")
	uploadFileApi := v1.ApiGroupApp.CommonApiGroup.UploadFileApi
	{
		uploadFileRouter.POST("", uploadFileApi.UploadFile)       // 上传文件
		uploadFileRouter.DELETE("/:id", uploadFileApi.DeleteFile) // 删除文件
	}
	{
		uploadFileRouterWithoutRecord.GET("", uploadFileApi.GetUploadFileList) // 获取文件列表
	}
}

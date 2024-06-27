package common

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
)

type UploadFileApi struct{}

// UploadFile 上传文件
func (f *UploadFileApi) UploadFile(c *gin.Context) {
	_, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.GGB_LOG.Error("接收文件失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId, _ := utils.GetUserID(c)
	err = uploadFileService.UploadFile(fileHeader, userId)
	if err != nil {
		global.GGB_LOG.Error("文件上传失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithMessage("文件上传成功！", c)
}

// DeleteFile 删除文件
func (f *UploadFileApi) DeleteFile(c *gin.Context) {
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	fileId, _ := utils.Str2uint(c.Param("id"))

	err := uploadFileService.DeleteFile(fileId)
	if err != nil {
		global.GGB_LOG.Error("删除文件失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithMessage("删除成功！", c)
}

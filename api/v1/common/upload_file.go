package common

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/request"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
)

type UploadFileApi struct{}

// UploadFile 上传文件
// @Tags      CommonUploadFile
// @Summary   上传文件
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file   		true  			"上传文件"
// @Success   200   {object}  response.UploadFileResponse  	"返回包括文件路径，文件名称"
// @Router    /common/upload [POST]
func (u *UploadFileApi) UploadFile(c *gin.Context) {
	_, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.GGB_LOG.Error("接收文件失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	userId, _ := utils.GetUserID(c)
	filePath, fileName, err := uploadFileService.UploadFile(fileHeader, userId)
	if err != nil {
		global.GGB_LOG.Error("文件上传失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithDetailed(response.UploadFileResponse{
		FilePath: filePath,
		FileName: fileName,
	}, "文件上传成功", c)
}

// DeleteFile 删除文件
func (u *UploadFileApi) DeleteFile(c *gin.Context) {
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	fileId, _ := utils.Str2uint(c.Param("id"))

	filePath, fileName, err := uploadFileService.DeleteFile(fileId)
	if err != nil {
		global.GGB_LOG.Error("删除文件失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithDetailed(response.UploadFileResponse{
		FilePath: filePath,
		FileName: fileName,
	}, "文件删除成功", c)
}

// GetUploadFileList
// @Tags      CommonUploadFile
// @Summary   获取上传文件列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param	  pageNumber 		query 	int 	true 	"分页"
// @Param	  pageSize  		query 	int 	true 	"每页条数"
// @Success   200   {object}  	response.Response{data=response.PageResult{list=[]common.UploadFile}}  "返回列表，总数"
// @Router    /common/upload [GET]
func (u *UploadFileApi) GetUploadFileList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	var query request.UploadFileQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := uploadFileService.GetUploadFileList(query, offset, limit)
	if err != nil {
		global.GGB_LOG.Error("查询列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

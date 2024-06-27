package common

import (
	"errors"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common"
	"github.com/wangyupo/GGB/model/common/request"
	"github.com/wangyupo/GGB/utils/upload"
	"mime/multipart"
)

type UploadFileService struct{}

// GetUploadFileList 获取列表
func (u *UploadFileService) GetUploadFileList(query request.UploadFileQuery, offset int, limit int) (list interface{}, total int64, err error) {
	// 声明 system.UploadFile 类型的变量以存储查询结果
	uploadFileList := make([]common.UploadFile, 0)

	// 准备数据库查询
	db := global.GGB_DB.Model(&common.UploadFile{})
	if query.FileName != "" {
		db = db.Where("file_name LIKE ?", "%"+query.FileName+"%")
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 获取分页数据
	err = db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&uploadFileList).Error

	return uploadFileList, total, err
}

// UploadFile 上传文件并将文件信息添加到数据库
func (u *UploadFileService) UploadFile(fileHeader *multipart.FileHeader, userId uint) (filePath string, fileName string, err error) {
	oss := upload.NewOss()
	filePath, fileName, err = oss.UploadFile(fileHeader)
	if err != nil {
		return
	}

	fileType := fileHeader.Header.Get("Content-Type") // 获取 MIME 类型
	fileSize := fileHeader.Size                       // 获取文件大小，单位字节

	file := common.UploadFile{
		FileName:         fileHeader.Filename,
		UploadedFileName: fileName,
		FilePath:         filePath,
		FileType:         fileType,
		FileSize:         fileSize,
		UserID:           userId,
	}

	err = global.GGB_DB.Create(&file).Error
	return filePath, fileHeader.Filename, err
}

// FindFile 查询文件
func (u *UploadFileService) FindFile(id uint) (file common.UploadFile, err error) {
	err = global.GGB_DB.Where("id = ?", id).First(&file).Error
	return file, err
}

// DeleteFile 将文件信息从数据库中删除并删除文件
func (u *UploadFileService) DeleteFile(fileId uint) (filePath string, fileName string, err error) {
	// 1-根据id从数据库找出文件
	var file common.UploadFile
	file, err = u.FindFile(fileId)
	if err != nil {
		return
	}

	// 2- 删除已上传文件实体
	oss := upload.NewOss()
	err = oss.DeleteFile(file.UploadedFileName)
	if err != nil {
		return file.FilePath, file.FileName, errors.New("文件删除失败")
	}

	// 3-从数据库删除文件信息记录（永久删除）
	err = global.GGB_DB.Where("id = ?", fileId).Unscoped().Delete(&file).Error

	return file.FilePath, file.FileName, err
}

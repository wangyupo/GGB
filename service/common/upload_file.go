package common

import (
	"errors"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common"
	"github.com/wangyupo/GGB/utils/upload"
	"mime/multipart"
)

type UploadFileService struct{}

// UploadFile 上传文件并将文件信息添加到数据库
func (u *UploadFileService) UploadFile(fileHeader *multipart.FileHeader, userId uint) (err error) {
	oss := upload.NewOss()
	filePath, fileName, err := oss.UploadFile(fileHeader)
	if err != nil {
		return err
	}

	fileType := fileHeader.Header.Get("Content-Type") // 获取 MIME 类型
	fileSize := fileHeader.Size                       // 获取文件大小，单位字节

	file := common.UploadFile{
		FileName: fileName,
		FilePath: filePath,
		FileType: fileType,
		FileSize: fileSize,
		UserID:   userId,
	}

	err = global.GGB_DB.Create(&file).Error
	return err
}

// FindFile 查询文件
func (u *UploadFileService) FindFile(id uint) (file common.UploadFile, err error) {
	err = global.GGB_DB.Where("id = ?", id).First(&file).Error
	return file, err
}

// DeleteFile 将文件信息从数据库中删除并删除文件
func (u *UploadFileService) DeleteFile(fileId uint) (err error) {
	// 1-根据id从数据库找出文件
	var file common.UploadFile
	file, err = u.FindFile(fileId)
	if err != nil {
		return
	}

	// 2- 删除已上传文件实体
	oss := upload.NewOss()
	err = oss.DeleteFile(file.FileName)
	if err != nil {
		return errors.New("文件删除失败")
	}

	// 3-从数据库删除文件信息记录（永久删除）
	err = global.GGB_DB.Where("id = ?", fileId).Unscoped().Delete(&file).Error

	return err
}

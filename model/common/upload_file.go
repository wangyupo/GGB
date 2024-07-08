package common

import "github.com/wangyupo/GGB/global"

type UploadFile struct {
	global.BaseModel
	FileName         string `json:"fileName" form:"fileName" gorm:"comment:文件名称"`                   // 文件名称
	UploadedFileName string `json:"uploadedFileName" form:"uploadedFileName" gorm:"comment:文件存储名称"` // 文件存储名称
	FilePath         string `json:"filePath" form:"filePath" gorm:"comment:文件存储路径"`                 // 文件存储路径
	FileSize         int64  `json:"fileSize" gorm:"comment:文件大小（字节）"`                               // 文件大小（字节）
	FileType         string `json:"fileType" gorm:"comment:文件MIME类型，如：image/png"`                   // 文件MIME类型，如：image/png
	UserID           uint   `json:"userId" form:"userId" gorm:"comment:上传用户id"`                     // 上传用户id
}

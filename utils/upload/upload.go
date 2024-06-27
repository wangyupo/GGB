package upload

import (
	"github.com/wangyupo/GGB/global"
	"mime/multipart"
)

// OSS 对象存储接口
type OSS interface {
	// UploadFile 上传文件
	UploadFile(file *multipart.FileHeader) (filePath string, fileName string, err error)
	// DeleteFile 删除文件
	DeleteFile(key string) (err error)
}

// NewOss OSS的实例化方法
func NewOss() OSS {
	switch global.GGB_CONFIG.System.OssType {
	case "local":
		return &Local{}
	default:
		return &Local{}
	}
}

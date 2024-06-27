package common

import "github.com/wangyupo/GGB/service"

type ApiGroup struct {
	UploadFileApi
}

var (
	uploadFileService = service.ServiceGroupApp.CommonService.UploadFileService
)

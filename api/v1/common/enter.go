package common

import "github.com/wangyupo/GGB/service"

type ApiGroup struct {
	UploadFileApi
	TranscriptApi
	EmailApi
}

var (
	uploadFileService = service.ServiceGroupApp.CommonService.UploadFileService
	transcriptService = service.ServiceGroupApp.CommonService.TranscriptService
)

package common

import "github.com/wangyupo/GGB/service"

type ApiGroup struct {
	UploadFileApi
	TranscriptApi
}

var (
	uploadFileService = service.ServiceGroupApp.CommonService.UploadFileService
	transcriptService = service.ServiceGroupApp.CommonService.TranscriptService
)

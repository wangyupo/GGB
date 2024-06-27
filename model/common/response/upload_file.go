package response

type UploadFileResponse struct {
	FilePath string `json:"filePath"`
	FileName string `json:"fileName"`
}

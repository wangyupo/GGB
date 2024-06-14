package system

import "github.com/wangyupo/GGB/global"

type SysRequestLog struct {
	global.BaseModel
	UserId       uint   `json:"userId"`
	Path         string `json:"path" gorm:"comment:请求路径"`
	Method       string `json:"method" gorm:"size:64;comment:请求方式"`
	QueryParams  string `json:"queryParams" gorm:"comment:请求参数"`
	RequestBody  string `json:"requestBody" gorm:"comment:请求结构体"`
	ResponseBody string `json:"responseBody" gorm:"comment:响应结构体"`
	StatusCode   uint   `json:"statusCode" gorm:"HTTP 响应状态码"`
}

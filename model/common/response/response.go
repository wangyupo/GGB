package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type MsgResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

const (
	ErrorToken = -1
	Error      = 7
	Success    = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func MsgResult(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, MsgResponse{
		code,
		msg,
	})
}

// SuccessWithMessage 返回成功并携带自定义消息
func SuccessWithMessage(message string, c *gin.Context) {
	MsgResult(Success, message, c)
}

// SuccessWithData 返回成功并携带自定义数据和默认消息
func SuccessWithData(data interface{}, c *gin.Context) {
	Result(Success, data, "查询成功", c)
}

// SuccessWithDetailed 返回成功并携带自定义数据和自定义消息
func SuccessWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(Success, data, message, c)
}

// FailWithMessage 返回失败并携带自定义消息
func FailWithMessage(message string, c *gin.Context) {
	MsgResult(Error, message, c)
}

// NoAuth 返回身份校验不通过并携带自定义消息
func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, MsgResponse{
		ErrorToken,
		message,
	})
}

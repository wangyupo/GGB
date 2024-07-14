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
	// Success 成功
	Success = 0
	// Error 通用错误
	Error = 1000 + iota
	// ErrorAuth 身份校验错误
	ErrorAuth
	// ErrorValidate 数据验证错误
	ErrorValidate
	// ErrorNotFound 接口不存在
	ErrorNotFound
	// ErrorMethodNotAllowed 请求方法不正确
	ErrorMethodNotAllowed
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

// SuccessWithDefaultMessage 返回成功并携带自定义消息
func SuccessWithDefaultMessage(c *gin.Context) {
	MsgResult(Success, "操作成功", c)
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

// FailWithValidate 返回失败并验证消息
func FailWithValidate(data interface{}, c *gin.Context) {
	Result(ErrorValidate, data, "数据校验未通过", c)
}

// NoAuth 返回身份校验不通过并携带自定义消息
func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, MsgResponse{
		ErrorAuth,
		message,
	})
}

// NotFound 接口不存在
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, MsgResponse{
		ErrorNotFound,
		"该接口不存在",
	})
}

// MethodNotAllowed 请求方法不正确
func MethodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, MsgResponse{
		ErrorMethodNotAllowed,
		"该方法不适用于所请求的 URL",
	})
}

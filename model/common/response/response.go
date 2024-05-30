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
	ERROR   = 7
	SUCCESS = 0
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

func SuccessWithMessage(message string, c *gin.Context) {
	MsgResult(SUCCESS, message, c)
}

func SuccessWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func FailWithMessage(message string, c *gin.Context) {
	MsgResult(ERROR, message, c)
}

func NoAuth(message string, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		7,
		nil,
		message,
	})
}

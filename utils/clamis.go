package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/router/system/request"
	"net"
)

// ClearToken 删除浏览器的cookie中的x-token
func ClearToken(c *gin.Context) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, false, false)
	}
}

// SetToken 向浏览器的cookie中设置x-token
func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		// host 是 IP 地址
		c.SetCookie("x-token", token, maxAge, "/", "", false, false)
	} else {
		// host 不是 IP 地址
		c.SetCookie("x-token", token, maxAge, "/", host, false, false)
	}
}

// GetToken 获取浏览器cookie中的x-token
func GetToken(c *gin.Context) string {
	token, _ := c.Cookie("x-token")
	if token == "" {
		token = c.Request.Header.Get("x-token")
	}
	return token
}

// GetClaims 解析token的声明（claims）内容
func GetClaims(c *gin.Context) (*request.CustomClaims, error) {
	token := GetToken(c)
	claims, err := ParseToken(token)
	if err != nil {
		fmt.Print(err.Error())
	}
	return claims, err
}

// GetUserID 从token的声明（claims）中获取user id
func GetUserID(c *gin.Context) (uint, error) {
	cl, err := GetClaims(c)
	return cl.BaseClaims.ID, err
}

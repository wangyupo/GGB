package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/router/system/request"
	"net"
)

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

func GetToken(c *gin.Context) string {
	token, _ := c.Cookie("x-token")
	if token == "" {
		token = c.Request.Header.Get("x-token")
	}
	return token
}

func GetClaims(c *gin.Context) *request.CustomClaims {
	token := GetToken(c)
	claims, _ := ParseToken(token)
	return claims
}

func GetUserID(c *gin.Context) uint {
	cl := GetClaims(c)
	return cl.BaseClaims.ID
}

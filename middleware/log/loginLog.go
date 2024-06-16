package log

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/utils"
)

func LoginLog(ActionType uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取参数
		userId, _ := utils.GetUserID(c)        // 从token获取用户id
		clientIP := c.ClientIP()               // 获取客户端IP
		userAgent := c.GetHeader("User-Agent") // 获取浏览器信息

		// 写入数据表
		loginLog := system.SysLogLogin{
			UserId:    userId,
			Type:      ActionType,
			IP:        clientIP,
			UserAgent: userAgent,
		}
		err := global.DB.Create(&loginLog).Error
		if err != nil {
		}

		c.Next()
	}
}

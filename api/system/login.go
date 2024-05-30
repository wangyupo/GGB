package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/system"
)

// 声明登录传参结构体
type loginForm struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	// 声明 loginForm 类型的变量以存储 JSON 数据
	var loginForm loginForm
	if err := c.BindJSON(&loginForm); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 声明 system.SysUser 类型的变量以存储找出来的用户
	var systemUser system.SysUser
	// 根据 userName 找用户（userName 是唯一的）
	global.DB.Where("user_name=?", loginForm.UserName).First(&systemUser)
	if systemUser.ID == 0 {
		response.FailWithMessage("用户未注册", c)
		return
	}
}

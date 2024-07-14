package system

import (
	"errors"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
	"github.com/wangyupo/GGB/utils"
)

type SysBaseService struct{}

// Login 登录
func (s *SysBaseService) Login(loginForm request.Login) (systemUser system.SysUser, err error) {
	// 根据 userName 找用户（userName 是唯一的）
	err = global.GGB_DB.Preload("Roles").
		Preload("Roles.Menus").Where("user_name = ?", loginForm.UserName).First(&systemUser).Error
	if err != nil {
		return
	}
	if systemUser.ID == 0 {
		return systemUser, errors.New("用户不存在")
	}
	if systemUser.Status == 0 {
		return systemUser, errors.New("用户已禁用")
	}

	// 核对密码
	ok := utils.BcryptCheck(loginForm.Password, systemUser.Password)
	if !ok {
		return systemUser, errors.New("登录密码错误")
	}

	return
}

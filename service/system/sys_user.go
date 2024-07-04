package system

import (
	"errors"
	"fmt"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
	"github.com/wangyupo/GGB/utils"
	"gorm.io/gorm"
)

type SysUserService struct{}

// Login 登录
func (s *SysUserService) Login(loginForm request.Login) (systemUser system.SysUser, err error) {
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

// ChangePassword 用户修改密码
func (s *SysUserService) ChangePassword(userId uint, req request.ChangePassword) (err error) {
	// 根据id查找用户
	var systemUser system.SysUser
	err = global.GGB_DB.Where("id = ?", userId).First(&systemUser).Error
	if err != nil {
		return err
	}
	// 校验原密码
	if ok := utils.BcryptCheck(req.Password, systemUser.Password); !ok {
		return errors.New("密码修改失败，原密码与当前账户不符")
	}
	// 原密码正确，hash新密码，更新sys_user
	systemUser.Password = utils.BcryptHash(req.NewPassword)
	err = global.GGB_DB.Save(&systemUser).Error
	return err
}

// ResetPassword 重置用户密码
func (s *SysUserService) ResetPassword(userId string, DefaultPassword string) (err error) {
	err = global.GGB_DB.Model(&system.SysUser{}).
		Where("id = ?", userId).
		Update("password", utils.BcryptHash(DefaultPassword)).Error
	return err
}

// GetSystemUserInfo 获取用户信息
func (s *SysUserService) GetSystemUserInfo(userId uint) (systemUser system.SysUser, err error) {
	err = global.GGB_DB.Model(&system.SysUser{}).Where("id = ?", userId).First(&systemUser).Error
	return systemUser, err
}

// GetSystemUserList 获取用户列表
func (s *SysUserService) GetSystemUserList(info request.SystemUserList, offset int, limit int) (list interface{}, total int64, err error) {
	// 声明 system.SysUser 类型的变量以存储查询结果
	systemUserList := make([]system.SysUser, 0)

	// 准备数据库查询
	db := global.GGB_DB.Model(&system.SysUser{})

	if info.UserName != "" {
		db = db.Where("user_name LIKE ?", "%"+info.UserName+"%")
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		// 错误处理
		return
	}

	// 获取分页数据
	err = db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&systemUserList).Error
	return systemUserList, total, err
}

// CreateSystemUser 新建用户
func (s *SysUserService) CreateSystemUser(systemUser system.SysUser, DefaultPassword string) (err error) {
	// 检查 UserName 是否重复
	err = global.GGB_DB.Where("user_name = ?", systemUser.UserName).First(&system.SysUser{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("用户名 %s 已存在", systemUser.UserName))
	}

	// 设置默认密码
	systemUser.Password = utils.BcryptHash(DefaultPassword)

	// 创建 systemUser 记录
	err = global.GGB_DB.Create(&systemUser).Error
	return err
}

// UpdateSystemUser 更新用户信息
func (s *SysUserService) UpdateSystemUser(systemUser system.SysUser, userId uint) (err error) {
	var oldSystemUser system.SysUser
	// 从数据库中查找具有指定 ID 的数据
	err = global.GGB_DB.Where("id = ?", userId).First(&oldSystemUser).Error
	if err != nil {
		// 错误处理
		return err
	}

	err = global.GGB_DB.Where("id != ? AND user_name = ?", userId, systemUser.UserName).
		First(&system.SysUser{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("用户名 %s 已存在", systemUser.UserName))
	}

	systemUserMap := map[string]interface{}{
		"UserName": systemUser.UserName,
		"NickName": systemUser.NickName,
		"Email":    systemUser.Email,
	}

	// 更新用户记录
	err = global.GGB_DB.Model(&oldSystemUser).Updates(systemUserMap).Error
	return err
}

// DeleteSystemUser 删除用户
func (s *SysUserService) DeleteSystemUser(userId uint) (err error) {
	err = global.GGB_DB.Where("id = ?", userId).Delete(&system.SysUser{}).Error
	return err
}

// ChangeSystemUserStatus 修改用户状态
func (s *SysUserService) ChangeSystemUserStatus(userId uint, status int) (err error) {
	err = global.GGB_DB.Model(&system.SysUser{}).
		Where("id = ?", userId).
		Update("status", status).Error
	return err
}

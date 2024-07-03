package system

import (
	"errors"
	"fmt"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
	"gorm.io/gorm"
)

type SysRoleService struct{}

func (s *SysRoleService) GetSysRoleList(query request.SysRoleQuery, offset int, limit int) (list interface{}, total int64, err error) {
	// 声明 system.SysRole 类型的变量以存储查询结果
	sysRoleList := make([]system.SysRole, 0)

	// 准备数据库查询
	db := global.GGB_DB.Model(&system.SysRole{})
	if query.RoleName != "" {
		db = db.Where("role_name LIKE ?", "%"+query.RoleName+"%")
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 获取分页数据
	err = db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&sysRoleList).Error
	return sysRoleList, total, err
}

// CreateSysRole 创建角色
func (s *SysRoleService) CreateSysRole(sysRole system.SysRole) (err error) {
	err = global.GGB_DB.Where("role_name = ?", sysRole.RoleName).First(&system.SysRole{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("角色 %s 已存在", sysRole.RoleName))
	}

	err = global.GGB_DB.Where("role_code = ?", sysRole.RoleCode).First(&system.SysRole{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("角色标识码 %s 已存在", sysRole.RoleCode))
	}

	// 创建 sysRole 记录
	err = global.GGB_DB.Create(&sysRole).Error
	return err
}

// GetSysRole 获取角色
func (s *SysRoleService) GetSysRole(userId uint) (sysRole system.SysRole, err error) {
	err = global.GGB_DB.First(&sysRole, userId).Error
	return sysRole, err
}

// UpdateSysRole 更新角色信息
func (s *SysRoleService) UpdateSysRole(sysRole system.SysRole, sysRoleId uint) (err error) {
	var oldSysRole system.SysRole
	err = global.GGB_DB.Where("id = ?", sysRoleId).First(&oldSysRole).Error
	if err != nil {
		return err
	}

	err = global.GGB_DB.Where("id != ? AND role_name = ?", sysRoleId, sysRole.RoleName).
		First(&system.SysRole{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("角色 %s 已存在", sysRole.RoleName))
	}

	err = global.GGB_DB.Where("id != ? AND role_code = ?", sysRoleId, sysRole.RoleCode).
		First(&system.SysRole{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("角色标识码 %s 已存在", sysRole.RoleCode))
	}

	sysRoleMap := map[string]interface{}{
		"RoleName":    sysRole.RoleName,
		"RoleCode":    sysRole.RoleCode,
		"Description": sysRole.Description,
	}

	err = global.GGB_DB.Model(&oldSysRole).Updates(sysRoleMap).Error
	return err
}

// DeleteSysRole 删除角色
func (s *SysRoleService) DeleteSysRole(sysRoleId uint) (err error) {
	err = global.GGB_DB.Where("id = ?", sysRoleId).Delete(&system.SysRole{}).Error
	return err
}

// ChangeRoleStatus 修改角色状态
func (s *SysRoleService) ChangeRoleStatus(sysRoleId uint, status int) (err error) {
	err = global.GGB_DB.Model(&system.SysRole{}).
		Where("id = ?", sysRoleId).
		Update("status", status).Error
	return err
}

// RoleAssignMenu 角色分配菜单
func (s *SysRoleService) RoleAssignMenu(req request.RoleAssignMenu) (err error) {
	// 删掉之前的菜单
	err = global.GGB_DB.Where("sys_role_id = ?", req.SysRoleID).Delete(&system.SysRoleMenu{}).Error
	if err != nil {
		return err
	}

	if len(req.SysMenuIds) == 0 {
		return nil
	}

	// 添加新菜单
	var roleMenu []system.SysRoleMenu
	for _, menuId := range req.SysMenuIds {
		roleMenu = append(roleMenu, system.SysRoleMenu{
			SysRoleID: req.SysRoleID,
			SysMenuID: menuId,
		})
	}

	err = global.GGB_DB.Create(&roleMenu).Error
	return err
}

// RoleAssignUser 角色分配给用户
func (s *SysRoleService) RoleAssignUser(req request.RoleAssignUser) (err error) {
	if len(req.SysUserIds) == 0 {
		return nil
	}

	// 绑定用户与角色
	var roleUser []system.SysRoleUser
	for _, userId := range req.SysUserIds {
		roleUser = append(roleUser, system.SysRoleUser{
			SysRoleID: req.SysRoleID,
			SysUserID: userId,
		})
	}

	err = global.GGB_DB.Create(roleUser).Error
	return err
}

// GetUserByRole 获取角色绑定的用户
func (s *SysRoleService) GetUserByRole(sysRoleId uint, offset int, limit int) (list interface{}, total int64, err error) {
	var users []system.SysUser

	err = global.GGB_DB.Model(&system.SysRoleUser{}).Where("sys_role_id = ?", sysRoleId).Count(&total).Error
	if err != nil {
		return
	}

	err = global.GGB_DB.Model(&system.SysUser{}).
		Joins("JOIN sys_role_user ON sys_role_user.sys_user_id = sys_user.id").
		Where("sys_role_user.sys_role_id = ? AND sys_role_user.deleted_at IS NULL", sysRoleId).
		Order("sys_role_user.created_at DESC").
		Offset(offset).Limit(limit).
		Find(&users).Error
	return users, total, err
}

// RoleUnAssignUser 角色取消绑定用户
func (s *SysRoleService) RoleUnAssignUser(req request.RoleAssignUser) (err error) {
	err = global.GGB_DB.Where("sys_role_id = ? AND sys_user_id in (?)", req.SysRoleID, req.SysUserIds).
		Delete(&system.SysRoleUser{}).Error
	return err
}

// GetMenuByRole 根据角色id查对应菜单
func (s *SysRoleService) GetMenuByRole(sysRoleId uint) (menus []system.SysMenu, err error) {
	// 根据roleId找到对应的菜单
	var roleMenus []system.SysRoleMenu
	err = global.GGB_DB.Where("sys_role_id = ?", sysRoleId).Find(&roleMenus).Error
	if err != nil {
		return
	}

	var menuIds []uint
	for _, roleMenu := range roleMenus {
		menuIds = append(menuIds, roleMenu.SysMenuID)
	}

	// 根据菜单id查找菜单
	err = global.GGB_DB.Where("id in (?)", menuIds).Order("sort").Find(&menus).Error

	return menus, err
}

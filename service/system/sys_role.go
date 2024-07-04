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

// GetSysRoleList 获取角色列表
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
	var sysRole system.SysRole
	err = global.GGB_DB.First(&sysRole, req.SysRoleID).Error
	if err != nil {
		return
	}

	var sysMenus []system.SysMenu
	global.GGB_DB.Find(&sysMenus, "id IN ?", req.SysMenuIds)
	err = global.GGB_DB.Model(&sysRole).Association("Menus").Replace(&sysMenus)

	return err
}

// RoleAssignUser 角色分配给用户
func (s *SysRoleService) RoleAssignUser(req request.RoleAssignUser) (err error) {
	var sysRole system.SysRole
	err = global.GGB_DB.First(&sysRole, req.SysRoleID).Error
	if err != nil {
		return
	}

	var sysUsers []system.SysUser
	global.GGB_DB.Find(&sysUsers, "id IN ?", req.SysUserIds)
	err = global.GGB_DB.Model(&sysRole).Association("Users").Append(&sysUsers)

	return err
}

// RoleUnAssignUser 角色取消绑定用户
func (s *SysRoleService) RoleUnAssignUser(req request.RoleAssignUser) (err error) {
	var sysRole system.SysRole
	err = global.GGB_DB.First(&sysRole, req.SysRoleID).Error
	if err != nil {
		return
	}

	var sysUsers []system.SysUser
	global.GGB_DB.Find(&sysUsers, "id IN ?", req.SysUserIds)
	err = global.GGB_DB.Model(&sysRole).Association("Users").Delete(&sysUsers)

	return err
}

// GetUserByRole 获取角色绑定的用户
func (s *SysRoleService) GetUserByRole(sysRoleId uint, offset int, limit int) (list interface{}, total int64, err error) {
	err = global.GGB_DB.Model(&system.SysRoleUser{}).Where("sys_role_id = ?", sysRoleId).Count(&total).Error
	if err != nil {
		return
	}

	// 非自定义连接表的查询方式（保留以供参考）
	//var sysRole system.SysRole
	//err = global.GGB_DB.Preload("Users", func(db *gorm.DB) *gorm.DB {
	//	return db.Offset(offset).Limit(limit)
	//}).First(&sysRole, sysRoleId).Error

	var sysRoleUsers []system.SysRoleUser
	err = global.GGB_DB.Preload("SysUser").Offset(offset).Limit(limit).Find(&sysRoleUsers, 1).Error

	var sysUsers []system.SysUser
	for _, sysRoleUser := range sysRoleUsers {
		sysUsers = append(sysUsers, *sysRoleUser.SysUser)
	}

	return sysUsers, total, err
}

// GetMenuByRole 获取角色绑定的菜单
func (s *SysRoleService) GetMenuByRole(sysRoleId uint) (menus []system.SysMenu, err error) {
	var sysRole system.SysRole
	err = global.GGB_DB.Preload("Menus").First(&sysRole, sysRoleId).Error

	for _, roleMenu := range sysRole.Menus {
		menus = append(menus, *roleMenu)
	}

	return menus, err
}

package system

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
	"github.com/wangyupo/GGB/utils"
	"gorm.io/gorm"
)

// GetSysRoleList 列表
func GetSysRoleList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	name := c.Query("roleName")

	// 声明 system.SysRole 类型的变量以存储查询结果
	sysRoleList := make([]system.SysRole, 0)
	var total int64

	// 准备数据库查询
	db := global.GGB_DB.Model(&system.SysRole{})
	if name != "" {
		db = db.Where("role_name LIKE ?", "%"+name+"%")
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取分页数据
	err := db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&sysRoleList).Error
	if err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  sysRoleList,
		Total: total,
	}, c)
}

// CreateSysRole 新建
func CreateSysRole(c *gin.Context) {
	// 声明 system.SysRole 类型的变量以存储 JSON 数据
	var sysRole system.SysRole

	// 绑定 JSON 请求体中的数据到 sysRole 结构体
	if err := c.ShouldBindJSON(&sysRole); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	if !errors.Is(global.GGB_DB.Where("role_name = ?", sysRole.RoleName).First(&system.SysRole{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("角色 %s 已存在", sysRole.RoleName), c)
		return
	}
	if !errors.Is(global.GGB_DB.Where("role_code = ?", sysRole.RoleCode).First(&system.SysRole{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("角色标识码 %s 已存在", sysRole.RoleCode), c)
		return
	}

	// 创建 sysRole 记录
	if err := global.GGB_DB.Create(&sysRole).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// GetSysRole 详情
func GetSysRole(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysRole 类型的变量以存储查询结果
	var sysRole system.SysRole

	// 从数据库中查找具有指定 ID 的数据
	if err := global.GGB_DB.First(&sysRole, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(sysRole, c)
}

// UpdateSysRole 编辑
func UpdateSysRole(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysRole 类型的变量以存储查询结果
	var sysRole system.SysRole

	// 从数据库中查找具有指定 ID 的数据
	if err := global.GGB_DB.First(&sysRole, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 绑定请求参数到数据对象
	if err := c.ShouldBindJSON(&sysRole); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	if !errors.Is(global.GGB_DB.Where("id != ? AND role_name = ?", id, sysRole.RoleName).First(&system.SysRole{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("角色 %s 已存在", sysRole.RoleName), c)
		return
	}
	if !errors.Is(global.GGB_DB.Where("id != ? AND role_code = ?", id, sysRole.RoleCode).First(&system.SysRole{}).Error, gorm.ErrRecordNotFound) {
		response.FailWithMessage(fmt.Sprintf("角色标识码 %s 已存在", sysRole.RoleCode), c)
		return
	}

	// 更新用户记录
	if err := global.GGB_DB.Save(&sysRole).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// DeleteSysRole 删除
func DeleteSysRole(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 根据指定 ID 删除数据
	if err := global.GGB_DB.Delete(&system.SysRole{}, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// ChangeRoleStatus 修改角色状态
func ChangeRoleStatus(c *gin.Context) {
	var req request.ChangeRoleStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := global.GGB_DB.Model(&system.SysRole{}).
		Where("id = ?", req.ID).
		Update("status", req.Status).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithDefaultMessage(c)
}

// RoleAssignMenu 角色分配菜单
func RoleAssignMenu(c *gin.Context) {
	var req request.RoleAssignMenu
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 删掉之前的菜单
	err := global.GGB_DB.Where("role_id = ?", req.RoleID).Delete(&system.SysRoleMenu{}).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if len(req.MenuIds) > 0 {
		// 添加新菜单
		var roleMenu []system.SysRoleMenu
		for _, id := range req.MenuIds {
			roleMenu = append(roleMenu, system.SysRoleMenu{
				RoleID: req.RoleID,
				MenuID: id,
			})
		}
		err = global.GGB_DB.Create(&roleMenu).Error
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}

	response.SuccessWithDefaultMessage(c)
}

// RoleAssignUser 角色分配给用户
func RoleAssignUser(c *gin.Context) {
	var req request.RoleAssignUser
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if len(req.UserIds) > 0 {
		// 绑定用户与角色
		var roleUser []system.SysRoleUser
		for _, id := range req.UserIds {
			roleUser = append(roleUser, system.SysRoleUser{
				RoleID: req.RoleID,
				UserID: id,
			})
		}
		err := global.GGB_DB.Create(roleUser).Error
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}

	response.SuccessWithDefaultMessage(c)
}

// GetUserByRole 获取角色绑定的用户
func GetUserByRole(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	roleId := c.Query("roleId")
	if roleId == "" {
		response.FailWithMessage("缺少参数：roleId", c)
		return
	}

	var users []system.SysUser
	var total int64

	err := global.GGB_DB.Model(&system.SysRoleUser{}).Where("role_id = ?", roleId).Count(&total).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = global.GGB_DB.Model(&system.SysUser{}).
		Joins("JOIN sys_role_user ON sys_role_user.user_id = sys_user.id").
		Where("sys_role_user.role_id = ? AND sys_role_user.deleted_at IS NULL", roleId).
		Order("sys_role_user.created_at DESC").
		Offset(offset).Limit(limit).
		Find(&users).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithData(response.PageResult{
		List:  users,
		Total: total,
	}, c)
}

// RoleUnAssignUser 角色取消绑定用户
func RoleUnAssignUser(c *gin.Context) {
	var req request.RoleAssignUser
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := global.GGB_DB.Where("role_id = ? AND user_id in (?)", req.RoleID, req.UserIds).
		Delete(&system.SysRoleUser{}).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithDefaultMessage(c)
}

// GetMenuByRole 根据角色id查对应菜单
func GetMenuByRole(c *gin.Context) {
	roleId := c.Param("id")
	if roleId == "" {
		response.FailWithMessage("缺少参数：角色id", c)
		return
	}

	// 根据roleId找到对应的菜单
	var roleMenu []system.SysRoleMenu
	err := global.GGB_DB.Where("role_id = ?", roleId).Find(&roleMenu).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var menuIds []uint
	for _, menu := range roleMenu {
		menuIds = append(menuIds, menu.MenuID)
	}

	// 根据菜单id查找菜单
	var menus []system.SysMenu
	err = global.GGB_DB.Where("id in (?)", menuIds).Order("sort").Find(&menus).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithData(menus, c)
}

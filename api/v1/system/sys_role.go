package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
)

type SysRoleApi struct{}

// GetSysRoleList 列表
func (s *SysRoleApi) GetSysRoleList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	var query request.SysRoleQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := sysRoleService.GetSysRoleList(query, offset, limit)
	if err != nil {
		global.GGB_LOG.Error("获取角色列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// CreateSysRole 新建
func (s *SysRoleApi) CreateSysRole(c *gin.Context) {
	// 声明 system.SysRole 类型的变量以存储 JSON 数据
	var sysRole system.SysRole

	// 绑定 JSON 请求体中的数据到 sysRole 结构体
	if err := c.ShouldBindJSON(&sysRole); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := sysRoleService.CreateSysRole(sysRole)
	if err != nil {
		global.GGB_LOG.Error("新建角色失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// GetSysRole 详情
func (s *SysRoleApi) GetSysRole(c *gin.Context) {
	// 获取路径参数
	id, err := utils.Str2uint(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	sysRole, err := sysRoleService.GetSysRole(id)
	if err != nil {
		global.GGB_LOG.Error("获取角色详情失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(sysRole, c)
}

// UpdateSysRole 编辑
func (s *SysRoleApi) UpdateSysRole(c *gin.Context) {
	// 获取路径参数
	id, err := utils.Str2uint(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 声明 system.SysRole 类型的变量以存储查询结果
	var sysRole system.SysRole
	// 绑定请求参数到数据对象
	if err := c.ShouldBindJSON(&sysRole); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = sysRoleService.UpdateSysRole(sysRole, id)
	if err != nil {
		global.GGB_LOG.Error("编辑角色失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// DeleteSysRole 删除
func (s *SysRoleApi) DeleteSysRole(c *gin.Context) {
	// 获取路径参数
	id, _ := utils.Str2uint(c.Param("id"))

	err := sysRoleService.DeleteSysRole(id)
	if err != nil {
		global.GGB_LOG.Error("删除角色失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// ChangeRoleStatus 修改角色状态
func (s *SysRoleApi) ChangeRoleStatus(c *gin.Context) {
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：角色id", c)
		return
	}
	roleId, _ := utils.Str2uint(c.Param("id"))

	var req request.ChangeRoleStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := sysRoleService.ChangeRoleStatus(roleId, req.Status)
	if err != nil {
		global.GGB_LOG.Error("修改角色状态失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithDefaultMessage(c)
}

// RoleAssignMenu 角色分配菜单
func (s *SysRoleApi) RoleAssignMenu(c *gin.Context) {
	var req request.RoleAssignMenu
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := sysRoleService.RoleAssignMenu(req)
	if err != nil {
		global.GGB_LOG.Error("角色分配菜单失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithDefaultMessage(c)
}

// RoleAssignUser 角色分配给用户
func (s *SysRoleApi) RoleAssignUser(c *gin.Context) {
	var req request.RoleAssignUser
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := sysRoleService.RoleAssignUser(req)
	if err != nil {
		global.GGB_LOG.Error("角色分配用户失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithDefaultMessage(c)
}

// GetUserByRole 获取角色绑定的用户
func (s *SysRoleApi) GetUserByRole(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	if c.Query("sysRoleId") == "" {
		response.FailWithMessage("缺少参数：sysRoleId", c)
		return
	}
	sysRoleId, _ := utils.Str2uint(c.Query("sysRoleId"))

	list, total, err := sysRoleService.GetUserByRole(sysRoleId, offset, limit)
	if err != nil {
		global.GGB_LOG.Error("获取角色绑定的用户失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// RoleUnAssignUser 角色取消绑定用户
func (s *SysRoleApi) RoleUnAssignUser(c *gin.Context) {
	var req request.RoleAssignUser
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := sysRoleService.RoleUnAssignUser(req)
	if err != nil {
		global.GGB_LOG.Error("角色取消绑定用户失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithDefaultMessage(c)
}

// GetMenuByRole 根据角色id查对应菜单
func (s *SysRoleApi) GetMenuByRole(c *gin.Context) {
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：角色id", c)
		return
	}
	sysRoleId, _ := utils.Str2uint(c.Param("id"))

	menus, err := sysRoleService.GetMenuByRole(sysRoleId)
	if err != nil {
		global.GGB_LOG.Error("根据角色id查对应菜单败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithData(menus, c)
}

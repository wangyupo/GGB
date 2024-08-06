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

// GetSysRoleList
//
//	@Tags		SysRole
//	@Summary	获取角色列表
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		pageNumber	query		int																	true	"分页"
//	@Param		pageSize	query		int																	true	"每页条数"
//	@Success	200			{object}	response.Response{data=response.PageResult{list=[]system.SysRole}}	"返回列表，总数"
//	@Router		/system/role [GET]
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

// CreateSysRole
//
//	@Tags		SysRole
//	@Summary	新建角色
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		system.SysRole			true	"SysRole模型"
//	@Success	200		{object}	response.MsgResponse	"返回操作成功提示"
//	@Router		/system/role [POST]
func (s *SysRoleApi) CreateSysRole(c *gin.Context) {
	// 声明 system.SysRole 类型的变量以存储 JSON 数据
	var sysRole system.SysRole

	// 绑定 JSON 请求体中的数据到 sysRole 结构体
	if err := c.ShouldBindJSON(&sysRole); err != nil {
		// 错误处理
		utils.HandleValidatorError(err, c)
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

// GetSysRole
//
//	@Tags		SysRole
//	@Summary	角色详情
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id	path		int										true	"角色id（roleId）"
//	@Success	200	{object}	response.Response{data=system.SysRole}	"返回角色详情"
//	@Router		/system/role/:id [GET]
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

// UpdateSysRole
//
//	@Tags		SysRole
//	@Summary	编辑角色
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id		path		int						true	"角色id（roleId）"
//	@Param		data	body		system.SysRole			true	"SysRole模型"
//	@Success	200		{object}	response.MsgResponse	"返回操作成功提示"
//	@Router		/system/role/:id [PUT]
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
		utils.HandleValidatorError(err, c)
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

// DeleteSysRole
//
//	@Tags		SysRole
//	@Summary	删除角色
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id	path		int						true	"角色id（roleId）"
//	@Success	200	{object}	response.MsgResponse	"返回操作成功提示"
//	@Router		/system/role/:id [DELETE]
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

// ChangeRoleStatus
//
//	@Tags		SysRole
//	@Summary	修改角色状态
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id		path		int							true	"角色id（roleId）"
//	@Param		data	body		request.ChangeRoleStatus	true	"ChangeRoleStatus模型"
//	@Success	200		{object}	response.MsgResponse		"返回操作成功提示"
//	@Router		/system/role/:id/status [PATCH]
func (s *SysRoleApi) ChangeRoleStatus(c *gin.Context) {
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：角色id", c)
		return
	}
	roleId, _ := utils.Str2uint(c.Param("id"))

	var req request.ChangeRoleStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidatorError(err, c)
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

// RoleAssignMenu
//
//	@Tags		SysRole
//	@Summary	角色分配菜单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.RoleAssignMenu	true	"RoleAssignMenu模型"
//	@Success	200		{object}	response.MsgResponse	"返回操作成功提示"
//	@Router		/system/role/menu [POST]
func (s *SysRoleApi) RoleAssignMenu(c *gin.Context) {
	var req request.RoleAssignMenu
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidatorError(err, c)
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

// RoleAssignUser
//
//	@Tags		SysRole
//	@Summary	角色分配给用户
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.RoleAssignUser	true	"RoleAssignUser模型"
//	@Success	200		{object}	response.MsgResponse	"返回操作成功提示"
//	@Router		/system/role/user [POST]
func (s *SysRoleApi) RoleAssignUser(c *gin.Context) {
	var req request.RoleAssignUser
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidatorError(err, c)
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

// RoleUnAssignUser
//
//	@Tags		SysRole
//	@Summary	角色取消绑定用户
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.RoleAssignUser	true	"RoleAssignUser模型"
//	@Success	200		{object}	response.MsgResponse	"返回操作成功提示"
//	@Router		/system/role/user [DELETE]
func (s *SysRoleApi) RoleUnAssignUser(c *gin.Context) {
	var req request.RoleAssignUser
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidatorError(err, c)
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

// GetUserByRole
//
//	@Tags		SysRole
//	@Summary	获取角色绑定的用户
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id			path		int											true	"角色id（roleId）"
//	@Param		pageNumber	query		int											true	"分页"
//	@Param		pageSize	query		int											true	"每页条数"
//	@Success	200			{object}	response.Response{data=response.PageResult}	"返回用户列表，总数"
//	@Router		/system/role/:id/user [GET]
func (s *SysRoleApi) GetUserByRole(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	sysRoleId, _ := utils.Str2uint(c.Param("id"))

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

// GetMenuByRole
//
//	@Tags		SysRole
//	@Summary	获取角色绑定的菜单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id	path		int											true	"角色id（roleId）"
//	@Success	200	{object}	response.Response{data=[]system.SysMenu}	"返回菜单列表"
//	@Router		/system/role/:id/menu [GET]
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

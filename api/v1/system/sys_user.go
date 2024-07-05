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

var (
	DefaultPassword string = "123456" // 重置密码的默认密码
)

type SysUserApi struct{}

// ChangePassword 修改密码
func (s *SysUserApi) ChangePassword(c *gin.Context) {
	var req request.ChangePassword
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidatorError(err, c)
		return
	}

	// 从token获取用户id
	userId, err := utils.GetUserID(c)
	if err != nil {
		global.GGB_LOG.Error("修改失败！获取userId失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = sysUserService.ChangePassword(userId, req)
	if err != nil {
		global.GGB_LOG.Error("修改用户密码失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithMessage("密码修改成功", c)
}

// ResetPassword 重置用户登录密码
func (s *SysUserApi) ResetPassword(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("用户id不能为空", c)
		return
	}

	err := sysUserService.ResetPassword(id, DefaultPassword)
	if err != nil {
		global.GGB_LOG.Error("重置用户密码失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithMessage("密码重置成功！", c)
}

// GetSystemUserInfo 根据token获取用户信息
func (s *SysUserApi) GetSystemUserInfo(c *gin.Context) {
	id, err := utils.GetUserID(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	systemUser, err := sysUserService.GetSystemUserInfo(id)
	if err != nil {
		global.GGB_LOG.Error("获取用户信息失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithData(systemUser, c)
}

// GetSystemUserList 列表
func (s *SysUserApi) GetSystemUserList(c *gin.Context) {
	// 获取分页参数
	offset, limit := utils.GetPaginationParams(c)
	// 获取其它查询参数
	var query request.SystemUserList
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := sysUserService.GetSystemUserList(query, offset, limit)
	if err != nil {
		global.GGB_LOG.Error("获取用户列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// CreateSystemUser 新建
func (s *SysUserApi) CreateSystemUser(c *gin.Context) {
	// 声明 system.SysUser 类型的变量以存储 JSON 数据
	var systemUser system.SysUser

	// 绑定 JSON 请求体中的数据到 systemUser 结构体
	if err := c.ShouldBindJSON(&systemUser); err != nil {
		// 错误处理
		utils.HandleValidatorError(err, c)
		return
	}

	err := sysUserService.CreateSystemUser(systemUser, DefaultPassword)
	if err != nil {
		global.GGB_LOG.Error("新建用户失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// GetSystemUser 详情
func (s *SysUserApi) GetSystemUser(c *gin.Context) {
	// 获取路径参数
	id, err := utils.Str2uint(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	systemUser, err := sysUserService.GetSystemUserInfo(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(systemUser, c)
}

// UpdateSystemUser 编辑
func (s *SysUserApi) UpdateSystemUser(c *gin.Context) {
	// 获取路径参数
	id, err := utils.Str2uint(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 声明 system.SysUser 类型的变量以存储查询结果
	var systemUser system.SysUser
	// 绑定请求参数到数据对象
	if err := c.ShouldBindJSON(&systemUser); err != nil {
		// 错误处理
		utils.HandleValidatorError(err, c)
		return
	}

	err = sysUserService.UpdateSystemUser(systemUser, id)
	if err != nil {
		global.GGB_LOG.Error("编辑用户失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// DeleteSystemUser 删除用户
func (s *SysUserApi) DeleteSystemUser(c *gin.Context) {
	// 获取路径参数
	id, err := utils.Str2uint(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = sysUserService.DeleteSystemUser(id)
	if err != nil {
		global.GGB_LOG.Error("删除用户失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to deleted systemUser", c)
}

// ChangeSystemUserStatus 修改用户状态
func (s *SysUserApi) ChangeSystemUserStatus(c *gin.Context) {
	id, err := utils.Str2uint(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var req request.ChangeSystemUserStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidatorError(err, c)
		return
	}

	err = sysUserService.ChangeSystemUserStatus(id, req.Status)
	if err != nil {
		global.GGB_LOG.Error("修改用户状态失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithDefaultMessage(c)
}

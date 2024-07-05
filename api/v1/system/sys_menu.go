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

type SysMenuApi struct{}

// GetSysMenuList 列表
func (s *SysMenuApi) GetSysMenuList(c *gin.Context) {
	// 获取其它查询参数
	var query request.SysMenuQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := sysMenuService.GetSysMenuList(query)
	if err != nil {
		global.GGB_LOG.Error("获取菜单列表失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// CreateSysMenu 新建
func (s *SysMenuApi) CreateSysMenu(c *gin.Context) {
	// 声明 system.SysMenu 类型的变量以存储 JSON 数据
	var sysMenu system.SysMenu

	// 绑定 JSON 请求体中的数据到 sysMenu 结构体
	if err := c.ShouldBindJSON(&sysMenu); err != nil {
		// 错误处理
		utils.HandleValidatorError(err, c)
		return
	}

	err := sysMenuService.CreateSysMenu(sysMenu)
	if err != nil {
		global.GGB_LOG.Error("新建菜单失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// GetSysMenu 详情
func (s *SysMenuApi) GetSysMenu(c *gin.Context) {
	// 获取路径参数
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	id, _ := utils.Str2uint(c.Param("id"))

	sysMenu, err := sysMenuService.GetSysMenu(id)
	if err != nil {
		global.GGB_LOG.Error("获取菜单详情失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(sysMenu, c)
}

// UpdateSysMenu 编辑
func (s *SysMenuApi) UpdateSysMenu(c *gin.Context) {
	// 获取路径参数
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	id, _ := utils.Str2uint(c.Param("id"))

	// 声明 system.SysMenu 类型的变量以存储查询结果
	var sysMenu system.SysMenu
	// 绑定请求参数到数据对象
	if err := c.ShouldBindJSON(&sysMenu); err != nil {
		// 错误处理
		utils.HandleValidatorError(err, c)
		return
	}

	err := sysMenuService.UpdateSysMenu(sysMenu, id)
	if err != nil {
		global.GGB_LOG.Error("编辑菜单失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// DeleteSysMenu 删除
func (s *SysMenuApi) DeleteSysMenu(c *gin.Context) {
	// 获取路径参数
	if c.Param("id") == "" {
		response.FailWithMessage("缺少参数：id", c)
		return
	}
	id, _ := utils.Str2uint(c.Param("id"))

	err := sysMenuService.DeleteSysMenu(id)
	if err != nil {
		global.GGB_LOG.Error("删除菜单失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithDefaultMessage(c)
}

// MoveSysMenu 菜单排序
func (s *SysMenuApi) MoveSysMenu(c *gin.Context) {
	var req request.MoveMenu
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidatorError(err, c)
		return
	}

	// 校验DropType
	validDropTypes := map[string]bool{
		"before": true,
		"inner":  true,
		"after":  true,
	}
	if !validDropTypes[req.DropType] {
		response.FailWithMessage("移动类型错误", c)
		return
	}

	err := sysMenuService.MoveSysMenu(req)
	if err != nil {
		global.GGB_LOG.Error("移动菜单失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithDefaultMessage(c)
}

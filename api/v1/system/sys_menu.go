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

// GetSysMenuList
//
//	@Tags		SysMenu
//	@Summary	获取所有菜单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Success	200	{object}	response.Response{data=response.PageResult{list=[]system.SysMenu}}	"返回所有菜单，总数"
//	@Router		/system/menu [GET]
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

// CreateSysMenu
//
//	@Tags		SysMenu
//	@Summary	新建菜单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		system.SysMenu			true	"SysMenu模型"
//	@Success	200		{object}	response.MsgResponse	"返回操作成功提示"
//	@Router		/system/menu [POST]
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

// GetSysMenu
//
//	@Tags		SysMenu
//	@Summary	菜单详情
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id	path		int										true	"菜单id（menuId）"
//	@Success	200	{object}	response.Response{data=system.SysMenu}	"返回菜详情"
//	@Router		/system/menu/:id [GET]
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

// UpdateSysMenu
//
//	@Tags		SysMenu
//	@Summary	编辑菜单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id		path		int						true	"菜单id（menuId）"
//	@Param		data	body		system.SysMenu			true	"SysMenu模型"
//	@Success	200		{object}	response.MsgResponse	"返回操作成功提示"
//	@Router		/system/menu/:id [PUT]
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

// DeleteSysMenu
//
//	@Tags		SysMenu
//	@Summary	删除菜单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		id	path		int						true	"菜单id（menuId）"
//	@Success	200	{object}	response.MsgResponse	"返回操作成功提示"
//	@Router		/system/menu/:id [DELETE]
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

// MoveSysMenu
//
//	@Tags		SysMenu
//	@Summary	菜单排序
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	body		request.MoveMenu		true	"MoveMenu模型"
//	@Success	200		{object}	response.MsgResponse	"返回操作成功提示"
//	@Router		/system/menu/move [PUT]
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

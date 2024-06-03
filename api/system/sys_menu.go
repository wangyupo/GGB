package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
	"github.com/wangyupo/GGB/utils"
)

// GetSysMenuList 列表
func GetSysMenuList(c *gin.Context) {
	// 获取分页参数
	pageNumber, pageSize := utils.GetPaginationParams(c)
	// 获取其它查询参数
	name := c.Query("name")

	// 声明 system.SysMenu 类型的变量以存储查询结果
	sysMenuList := make([]system.SysMenu, 0)
	var total int64

	// 准备数据库查询
	db := global.DB.Model(&system.SysMenu{})
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取分页数据
	db = db.Offset((pageNumber - 1) * pageSize).Limit(pageSize)

	// 执行查询并获取结果
	if err := db.Find(&sysMenuList).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(response.PageResult{
		List:  sysMenuList,
		Total: total,
	}, c)
}

// CreateSysMenu 新建
func CreateSysMenu(c *gin.Context) {
	// 声明 system.SysMenu 类型的变量以存储 JSON 数据
	var sysMenu system.SysMenu

	// 绑定 JSON 请求体中的数据到 sysMenu 结构体
	if err := c.ShouldBindJSON(&sysMenu); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 创建 sysMenu 记录
	if err := global.DB.Create(&sysMenu).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to create sysMenu", c)
}

// GetSysMenu 详情
func GetSysMenu(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysMenu 类型的变量以存储查询结果
	var sysMenu system.SysMenu

	// 从数据库中查找具有指定 ID 的数据
	if err := global.DB.First(&sysMenu, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithData(sysMenu, c)
}

// UpdateSysMenu 编辑
func UpdateSysMenu(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 声明 system.SysMenu 类型的变量以存储查询结果
	var sysMenu system.SysMenu

	// 从数据库中查找具有指定 ID 的数据
	if err := global.DB.First(&sysMenu, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 绑定请求参数到数据对象
	if err := c.ShouldBindJSON(&sysMenu); err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 更新用户记录
	if err := global.DB.Save(&sysMenu).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to update sysMenu", c)
}

// DeleteSysMenu 删除
func DeleteSysMenu(c *gin.Context) {
	// 获取路径参数
	id := c.Param("id")

	// 根据指定 ID 删除数据
	if err := global.DB.Delete(&system.SysMenu{}, id).Error; err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 返回响应结果
	response.SuccessWithMessage("Success to deleted sysMenu", c)
}

// MoveSysMenu 菜单排序
func MoveSysMenu(c *gin.Context) {
	var req request.MoveMenu
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
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

	// 找到源菜单、目标菜单
	var originMenu, targetMenu system.SysMenu
	err := global.DB.Where("id = ?", req.OriginID).First(&originMenu).Error
	if err != nil {
		response.FailWithMessage("移动菜单未找到", c)
		return
	}
	err = global.DB.Where("id = ?", req.TargetID).First(&targetMenu).Error
	if err != nil {
		response.FailWithMessage("目标菜单未找到", c)
		return
	}

	var parentId uint
	sortOrder := targetMenu.Sort // 将目标菜单的sort（排序值）作为基准

	switch req.DropType {
	case "before":
		parentId = targetMenu.ParentId
		sortOrder -= 1
	case "after":
		parentId = targetMenu.ParentId
		sortOrder += 1
	case "inner":
		parentId = targetMenu.ID
		sortOrder = 1
	default:
		response.FailWithMessage("移动类型错误", c)
		return
	}

	// 注：使用map更新，防止gorm的updates方法把sort=0跳过
	err = global.DB.Model(&system.SysMenu{}).
		Where("id = ?", originMenu.ID).
		Updates(map[string]interface{}{
			"ParentId": parentId,
			"Sort":     sortOrder,
		}).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithMessage("菜单移动成功！", c)
}

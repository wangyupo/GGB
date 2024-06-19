package system

import (
	"errors"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
)

type SysMenuService struct{}

// GetSysMenuList 获取所有菜单
func (sysMenuService *SysMenuService) GetSysMenuList(query request.SysMenuQuery) (list interface{}, total int64, err error) {
	// 声明 system.SysMenu 类型的变量以存储查询结果
	sysMenuList := make([]system.SysMenu, 0)

	// 准备数据库查询
	db := global.GGB_DB.Model(&system.SysMenu{})
	if query.Label != "" {
		db = db.Where("label LIKE ?", "%"+query.Label+"%")
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 获取分页数据
	err = db.Find(&sysMenuList).Error

	return sysMenuList, total, err
}

// CreateSysMenu 新建菜单
func (sysMenuService *SysMenuService) CreateSysMenu(sysMenu system.SysMenu) (err error) {
	err = global.GGB_DB.Create(&sysMenu).Error
	return err
}

// GetSysMenu 菜单详情
func (sysMenuService *SysMenuService) GetSysMenu(menuId uint) (sysMenu system.SysMenu, err error) {
	err = global.GGB_DB.First(&sysMenu, menuId).Error
	return sysMenu, err
}

// UpdateSysMenu 编辑菜单
func (sysMenuService *SysMenuService) UpdateSysMenu(sysMenu system.SysMenu, menuId uint) (err error) {
	// 从数据库中查找具有指定 ID 的数据
	var oldSysMenu system.SysMenu
	err = global.GGB_DB.First(&oldSysMenu, menuId).Error
	if err != nil {
		return
	}

	var sysMenuMap = map[string]interface{}{
		"Label":    sysMenu.Label,
		"Path":     sysMenu.Path,
		"Icon":     sysMenu.Icon,
		"ParentId": sysMenu.ParentId,
		"Sort":     sysMenu.Sort,
		"Type":     sysMenu.Type,
	}

	// 更新用户记录
	err = global.GGB_DB.Model(&oldSysMenu).Updates(sysMenuMap).Error
	return err
}

// DeleteSysMenu 删除菜单
func (sysMenuService *SysMenuService) DeleteSysMenu(menuId uint) (err error) {
	err = global.GGB_DB.Where("id = ?", menuId).Delete(&system.SysMenu{}).Error
	return err
}

// MoveSysMenu 菜单排序
func (sysMenuService *SysMenuService) MoveSysMenu(req request.MoveMenu) (err error) {
	// 找到源菜单、目标菜单
	var originMenu, targetMenu system.SysMenu
	err = global.GGB_DB.Where("id = ?", req.OriginID).First(&originMenu).Error
	if err != nil {
		return errors.New("移动菜单未找到")
	}
	err = global.GGB_DB.Where("id = ?", req.TargetID).First(&targetMenu).Error
	if err != nil {
		return errors.New("目标菜单未找到")
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
		return errors.New("移动类型错误")
	}

	// 注：使用map更新，防止gorm的updates方法把sort=0跳过
	var sysMenuMap = map[string]interface{}{
		"ParentId": parentId,
		"Sort":     sortOrder,
	}
	err = global.GGB_DB.Model(&system.SysMenu{}).Where("id = ?", originMenu.ID).Updates(sysMenuMap).Error

	return err
}

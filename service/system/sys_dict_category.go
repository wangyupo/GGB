package system

import (
	"errors"
	"fmt"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
	"gorm.io/gorm"
)

type SysDictCategoryService struct{}

// GetSysDictCategoryList 获取字典类型列表
func (sysDictCategoryService *SysDictCategoryService) GetSysDictCategoryList(query request.SysDictCategoryQuery, offset int, limit int) (list interface{}, total int64, err error) {
	// 声明 system.SysDictCategory 类型的变量以存储查询结果
	sysDictCategoryList := make([]system.SysDictCategory, 0)

	// 准备数据库查询
	db := global.GGB_DB.Model(&system.SysDictCategory{})
	if query.Label != "" {
		db = db.Where("label LIKE ?", "%"+query.Label+"%")
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		// 错误处理
		return
	}

	// 获取分页数据
	err = db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&sysDictCategoryList).Error

	return sysDictCategoryList, total, err
}

// CreateSysDictCategory 创建字典类型
func (sysDictCategoryService *SysDictCategoryService) CreateSysDictCategory(req system.SysDictCategory) (err error) {
	err = global.GGB_DB.Where("label = ?", req.Label).First(&system.SysDictCategory{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("字典 %s 已存在", req.Label))
	}

	err = global.GGB_DB.Where("label_code = ?", req.LabelCode).First(&system.SysDictCategory{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("字典编码 %s 已存在", req.LabelCode))
	}

	// 创建 sysDictCategory 记录
	err = global.GGB_DB.Create(&req).Error

	return err
}

// GetSysDictCategory 获取字典类型详情
func (sysDictCategoryService *SysDictCategoryService) GetSysDictCategory(dictCategoryId uint) (sysDictCategory system.SysDictCategory, err error) {
	err = global.GGB_DB.Where("id = ?", dictCategoryId).First(&sysDictCategory).Error
	return sysDictCategory, err
}

// UpdateSysDictCategory 更新字典类型
func (sysDictCategoryService *SysDictCategoryService) UpdateSysDictCategory(req system.SysDictCategory, dictCategoryId uint) (err error) {
	var oldSysDictCategory system.SysDictCategory

	// 从数据库中查找具有指定 ID 的数据
	err = global.GGB_DB.Where("id = ?", dictCategoryId).First(&oldSysDictCategory).Error
	if err != nil {
		return err
	}

	err = global.GGB_DB.Where("label = ? AND id != ?", req.Label, dictCategoryId).First(&system.SysDictCategory{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("字典 %s 已存在", req.Label))
	}

	err = global.GGB_DB.Where("label_code = ? AND id != ?", req.LabelCode, dictCategoryId).First(&system.SysDictCategory{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("字典编码 %s 已存在", req.LabelCode))
	}

	sysDictCategoryMap := map[string]interface{}{
		"Label":       req.Label,
		"LabelCode":   req.LabelCode,
		"Description": req.Description,
	}
	// 更新用户记录
	err = global.GGB_DB.Model(&oldSysDictCategory).Updates(sysDictCategoryMap).Error

	return err
}

// DeleteSysDictCategory 删除字典类型
func (sysDictCategoryService *SysDictCategoryService) DeleteSysDictCategory(dictCategoryId uint) (err error) {
	err = global.GGB_DB.Delete(&system.SysDictCategory{}, dictCategoryId).Error
	return err
}

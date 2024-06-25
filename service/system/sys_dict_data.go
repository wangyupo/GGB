package system

import (
	"errors"
	"fmt"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/system"
	"github.com/wangyupo/GGB/model/system/request"
	"gorm.io/gorm"
)

type SysDictDataService struct{}

// GetSysDictDataList 获取字典数据列表
func (s *SysDictDataService) GetSysDictDataList(query request.SysDictDataQuery, offset int, limit int) (list interface{}, total int64, err error) {
	// 声明 system.SysDictData 类型的变量以存储查询结果
	sysDictDataList := make([]system.SysDictData, 0)

	// 准备数据库查询
	db := global.GGB_DB.Model(&system.SysDictData{}).Where("category_id = ?", query.CategoryId)
	if query.Label != "" {
		db = db.Where("label LIKE ?", "%"+query.Label+"%")
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 获取分页数据
	err = db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&sysDictDataList).Error

	return sysDictDataList, total, err
}

// CreateSysDictData 新建字典数据
func (s *SysDictDataService) CreateSysDictData(req system.SysDictData) (err error) {
	err = global.GGB_DB.Where("label = ? AND category_id = ?", req.Label, req.CategoryID).First(&system.SysDictData{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("字典键 %s 已存在", req.Label))
	}

	err = global.GGB_DB.Where("value = ? AND category_id = ?", req.Value, req.CategoryID).First(&system.SysDictData{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("字典值 %s 已存在", req.Value))
	}

	// 创建 sysDictData 记录
	err = global.GGB_DB.Create(&req).Error

	return err
}

// GetSysDictData 获取字典数据
func (s *SysDictDataService) GetSysDictData(sysDictDataId uint) (sysDictData system.SysDictData, err error) {
	err = global.GGB_DB.First(&sysDictData, sysDictDataId).Error
	return sysDictData, err
}

// UpdateSysDictData 编辑字典数据
func (s *SysDictDataService) UpdateSysDictData(req system.SysDictData, sysDictDataId uint) (err error) {
	// 从数据库中查找具有指定 ID 的数据
	var oldSysDictData system.SysDictData
	err = global.GGB_DB.Where("id = ?", sysDictDataId).First(&oldSysDictData).Error
	if err != nil {
		return
	}

	err = global.GGB_DB.Where("label = ? AND id != ?", req.Label, sysDictDataId).First(&system.SysDictData{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("字典键 %s 已存在", req.Label))
	}

	err = global.GGB_DB.Where("value = ? AND id != ?", req.Value, sysDictDataId).First(&system.SysDictData{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("字典值 %s 已存在", req.Value))
	}

	sysDictDataMap := map[string]interface{}{
		"CategoryID":  req.CategoryID,
		"Label":       req.Label,
		"Value":       req.Value,
		"Description": req.Description,
	}
	// 更新用户记录
	err = global.GGB_DB.Model(&oldSysDictData).Updates(sysDictDataMap).Error

	return err
}

// DeleteSysDictData 删除字典数据
func (s *SysDictDataService) DeleteSysDictData(sysDictDataId uint) (err error) {
	err = global.GGB_DB.Delete(&system.SysDictData{}, sysDictDataId).Error
	return err
}

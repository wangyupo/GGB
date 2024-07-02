package common

import (
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common"
	"github.com/wangyupo/GGB/model/common/request"
	"github.com/wangyupo/GGB/utils"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
)

type TranscriptService struct{}

// ImportByExcel 通过Excel导入数据
func (t *TranscriptService) ImportByExcel(file multipart.File, userId uint) (err error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return
	}
	defer f.Close()

	// 读取Excel第一个sheet的所有行
	sheetName := f.GetSheetName(0) // 获取第一个表单的sheetName
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return
	}

	// 遍历行并储存数据
	keysOrder := []string{"name", "language", "math", "english", "geography", "politics"}
	var transcripts []common.Transcript
	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue // 跳过标题行
		}
		var transcript common.Transcript
		transcript.UserId = userId
		for cellIndex, cellValue := range row {
			switch keysOrder[cellIndex] {
			case "name":
				transcript.Name = cellValue
			case "language":
				language, _ := utils.Str2uint(cellValue)
				transcript.Language = language
			case "math":
				math, _ := utils.Str2uint(cellValue)
				transcript.Math = math
			case "english":
				english, _ := utils.Str2uint(cellValue)
				transcript.English = english
			case "geography":
				geography, _ := utils.Str2uint(cellValue)
				transcript.Geography = geography
			case "politics":
				politics, _ := utils.Str2uint(cellValue)
				transcript.Politics = politics
			}
		}
		transcripts = append(transcripts, transcript)
	}
	err = global.GGB_DB.Create(&transcripts).Error

	return
}

// GetTranscriptList 获取成绩列表
func (t *TranscriptService) GetTranscriptList(query request.TranscriptQuery, offset int, limit int) (list interface{}, total int64, err error) {
	// 声明 common.Transcript 类型的变量以存储查询结果
	transcriptList := make([]common.Transcript, 0)

	// 准备数据库查询
	db := global.GGB_DB.Model(&common.Transcript{})
	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}

	// 获取总数
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 获取分页数据
	err = db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&transcriptList).Error

	return transcriptList, total, err
}

// CreateTranscript 创建成绩
func (t *TranscriptService) CreateTranscript(req common.Transcript) (err error) {
	//err = global.GGB_DB.Where("label = ?", req.Label).First(&common.Transcript{}).Error
	//if !errors.Is(err, gorm.ErrRecordNotFound) {
	//	return errors.New(fmt.Sprintf("成绩 %s 已存在", req.Label))
	//}

	// 创建 transcript 记录
	err = global.GGB_DB.Create(&req).Error

	return err
}

// DeleteTranscript 删除成绩
func (t *TranscriptService) DeleteTranscript(transcriptId uint) (err error) {
	err = global.GGB_DB.Where("id = ?", transcriptId).Delete(&common.Transcript{}).Error
	return err
}

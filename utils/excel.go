package utils

import (
	"errors"
	"fmt"
	"github.com/wangyupo/GGB/global"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"reflect"
	"strings"
	"time"
)

// ExportExcelByTemplate 根据指定模板导出Excel
func ExportExcelByTemplate(templateName string, list interface{}) (err error) {
	// 1-读取Excel模板Sheet1首行
	templateFile, err := excelize.OpenFile(global.GGB_CONFIG.Excel.TemplateDir + templateName)
	if err != nil {
		global.GGB_LOG.Error(fmt.Sprintf("读取%s的Excel模板失败", templateName), zap.Error(err))
		return err
	}
	defer func() {
		if err := templateFile.Close(); err != nil {
			global.GGB_LOG.Error(fmt.Sprintf("关闭%s的Excel模板失败", templateName), zap.Error(err))
		}
	}()
	allRows, err := templateFile.GetRows("Sheet1")
	if err != nil {
		global.GGB_LOG.Error(fmt.Sprintf("获取%s的Excel的所有行失败", templateName), zap.Error(err))
		return err
	}
	if len(allRows) == 0 {
		return errors.New(fmt.Sprintf("%s的Excel模板未设置首行", templateName))
	}

	// 1-1 裁切中文名和字段名
	firstRowTitle := make([]string, 0, len(allRows[0])) // 首行中文
	firstRowKeys := make([]string, 0, len(allRows[0]))  // 数据key
	for _, colValue := range allRows[0] {
		result := strings.Split(colValue, "|")
		firstRowTitle = append(firstRowTitle, result[0])
		firstRowKeys = append(firstRowKeys, result[1])
	}

	// 2-创建新的Excel文件，并将读取到的首行设置为新Excel的Sheet1的首行
	newFile := excelize.NewFile()
	defer func() {
		if err := newFile.Close(); err != nil {
			global.GGB_LOG.Error(fmt.Sprintf("关闭%s的新Excel失败", templateName), zap.Error(err))
		}
	}()
	sheetName := "Sheet1"

	// 2-1 给新建的Excel设置首行
	firstRowStyle, _ := newFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#FFFF00"}, // 明黄色, 可以替换为你喜欢的颜色代码
		},
		Font: &excelize.Font{
			Bold: true, // 设置字体为加粗
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left", // 水平方向上文本左对齐
		},
	})
	for colIndex, cellValue := range firstRowTitle {
		cellName, _ := excelize.CoordinatesToCellName(colIndex+1, 1)
		_ = newFile.SetCellValue(sheetName, cellName, cellValue)               // 填充首行内容
		_ = newFile.SetCellStyle(sheetName, cellName, cellName, firstRowStyle) // 设置首行样式
	}

	// 2-2 设置首行冻结
	_ = newFile.SetPanes("Sheet1", &excelize.Panes{
		Freeze:      true,         // 开启冻结功能
		Split:       false,        // 不使用拆分视图功能
		XSplit:      0,            // X轴（水平方向）上不冻结，所以设为0
		YSplit:      1,            // Y轴（垂直方向）上冻结首行，所以设置为1。这意味着在第一行下方冻结
		TopLeftCell: "A2",         // 视图冻结后，最上方可见的首个单元格设置为A2，与YSplit结合，实际冻结顶端行
		ActivePane:  "bottomLeft", // 当前活动的窗格设置为左下方
	})

	//for colIndex := 1; colIndex <= len(firstRowTitle); colIndex++ {
	//	colName, _ := excelize.ColumnNumberToName(colIndex)                    // 将列索引号转换为列名称
	//	cellName := colName + "1"                                              // 生成单元格名称
	//	_ = newFile.SetCellStyle(sheetName, cellName, cellName, firstRowStyle) // 应用样式
	//}

	// 3-根据传入的list依次填充
	//contentRowStyle, _ := newFile.NewStyle(&excelize.Style{
	//	Alignment: &excelize.Alignment{
	//		Horizontal: "left", // 水平方向上文本左对齐
	//	},
	//})
	//dateTimeFormat := "yyyy-mm-dd hh:mm:ss"
	//dateTimeStyle, _ := newFile.NewStyle(&excelize.Style{
	//	CustomNumFmt: &dateTimeFormat,
	//})
	for rowIndex, rowData := range list.([]map[string]interface{}) {
		for colIndex, key := range firstRowKeys {
			for _, _ = range rowData {
				// 写入单元格内容
				cellRef, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2)
				_ = newFile.SetCellValue(sheetName, cellRef, rowData[key])

				// 设置单元格样式
				//cellName, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+1)
				//_ = newFile.SetCellStyle(sheetName, cellName, cellName, contentRowStyle)
			}
		}
	}

	// 4-保存文件
	outputFileName := MD5V([]byte(strings.Split(templateName, ".")[0] + time.Now().Format(time.TimeOnly)))
	err = newFile.SaveAs(global.GGB_CONFIG.Excel.OutputDir + outputFileName + ".xlsx")
	if err != nil {
		global.GGB_LOG.Error(fmt.Sprintf("保存%s的新Excel失败", templateName), zap.Error(err))
	}

	return nil
}

func IsTimeType(value interface{}) bool {
	typeOfValue := reflect.TypeOf(value)
	_, ok := value.(time.Time)
	return typeOfValue.String() == "time.Time" || ok
}

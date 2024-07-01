package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/wangyupo/GGB/global"
	"github.com/xuri/excelize/v2"
)

// ExportExcelByTemplate 根据指定模板导出Excel
func ExportExcelByTemplate(list [][]interface{}) (filePath string, err error) {
	// 1-新建Excel
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
		}
	}()
	sheetName := "Sheet1"

	// 1-1 设置首行冻结
	_ = f.SetPanes("Sheet1", &excelize.Panes{
		Freeze:      true,         // 开启冻结功能
		Split:       false,        // 不使用拆分视图功能
		XSplit:      0,            // X轴（水平方向）上不冻结，所以设为0
		YSplit:      1,            // Y轴（垂直方向）上冻结首行，所以设置为1。这意味着在第一行下方冻结
		TopLeftCell: "A2",         // 视图冻结后，最上方可见的首个单元格设置为A2，与YSplit结合，实际冻结顶端行
		ActivePane:  "bottomLeft", // 当前活动的窗格设置为左下方
	})

	// 1-2 设置首行样式
	firstRowStyle, _ := f.NewStyle(&excelize.Style{
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

	// 1-3 设置内容样式
	contentRowStyle, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left", // 水平方向上文本左对齐
		},
	})

	// 1-4 填充数据和样式
	for rowIndex, rowData := range list {
		cellNames, _ := generateCellNames(rowIndex+1, len(list[0]))
		startCell, _ := excelize.JoinCellName("A", rowIndex+1)
		_ = f.SetSheetRow(sheetName, startCell, &rowData) // 填充数据
		style := contentRowStyle
		if rowIndex == 0 {
			style = firstRowStyle
		}
		_ = f.SetCellStyle(sheetName, cellNames[0], cellNames[len(cellNames)-1], style) // 填充样式

	}

	// 2-保存文件
	filePath = global.GGB_CONFIG.Excel.OutputDir + uuid.New().String() + ".xlsx"
	err = f.SaveAs(filePath)
	if err != nil {
		global.GGB_LOG.Error("保存Excel失败")
	}

	return filePath, nil
}

// generateCellNames 根据给定的行号与单元格数量，生成一个含有 Excel 单元格名称的字符串切片。
// 参数: rowNum: 指定希望生成单元格名称所在的行号。n: 需要生成的单元格数量。
// 返回值：例如，generateCellNames(1, 3) 将返回 ["A1", "B1", "C1"], nil
func generateCellNames(rowNum int, n int) ([]string, error) {
	// 创建一个字符串切片 cellNames 用于存储生成的单元格名称，初始长度为 0，容量为 n
	cellNames := make([]string, 0, n)

	// 循环 n 次，生成 n 个单元格名称
	for i := 0; i < n; i++ {
		// 使用 excelize 库的 ColumnNumberToName 函数将列号 (i+1) 转换为对应的字母表示，例如 1 -> "A", 2 -> "B"
		colName, err := excelize.ColumnNumberToName(i + 1)

		// 如果转换过程中出现错误，返回 nil 和错误信息
		if err != nil {
			return nil, err
		}

		// 使用 fmt.Sprintf 函数拼接列字母和行号，生成单元格名称，例如 "A1", "B1"
		cellName := fmt.Sprintf("%s%d", colName, rowNum)

		// 将生成的单元格名称追加到 cellNames 切片中
		cellNames = append(cellNames, cellName)
	}

	// 返回生成的单元格名称切片和 nil 错误
	return cellNames, nil
}

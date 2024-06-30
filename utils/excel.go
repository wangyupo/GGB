package utils

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/wangyupo/GGB/global"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// ExportExcelByTemplate 根据指定模板导出Excel
func ExportExcelByTemplate(templateName string, list interface{}) (filePath string, err error) {
	// 1-读取Excel模板Sheet1首行
	templateFile, err := excelize.OpenFile(global.GGB_CONFIG.Excel.TemplateDir + templateName)
	if err != nil {
		global.GGB_LOG.Error(fmt.Sprintf("读取%s的Excel模板失败", templateName), zap.Error(err))
		return
	}
	defer func() {
		if err := templateFile.Close(); err != nil {
			global.GGB_LOG.Error(fmt.Sprintf("关闭%s的Excel模板失败", templateName), zap.Error(err))
		}
	}()
	allRows, _ := templateFile.GetRows("Sheet1")

	// 1-1 根据Excel模板首行，解析首行、需要映射中文枚举、日期格式化等备用
	firstRowTitle := make([]interface{}, 0, len(allRows[0])) // 首行中文
	firstRowKeys := make([]interface{}, 0, len(allRows[0]))  // 数据key
	firstRowEnums := make(map[interface{}]interface{})       // 数据枚举
	firstRowDateFormat := make(map[interface{}]interface{})  // 日期格式化
	for colIndex, colValue := range allRows[0] {
		var colValueJson map[string]interface{}
		_ = json.Unmarshal([]byte(colValue), &colValueJson)
		firstRowTitle = append(firstRowTitle, colValueJson["title"])
		firstRowKeys = append(firstRowKeys, colValueJson["key"])
		if _, ok := colValueJson["enum"]; ok {
			firstRowEnums[colIndex] = colValueJson["enum"]
		}
		if _, ok := colValueJson["dateFormat"]; ok {
			firstRowDateFormat[colIndex] = colValueJson["dateFormat"]
		}
	}

	// 2-根据Excel首行的key过滤数据、映射中文，输出包含标题的整体切片
	mergedSlice := make([][]interface{}, 0)
	mergedSlice = append(mergedSlice, firstRowTitle)
	for _, listItem := range list.([]map[string]interface{}) {
		contentItemList := make([]interface{}, 0, len(firstRowKeys))
		for colIndex, key := range firstRowKeys {
			value := listItem[key.(string)]
			if innerMap, ok := firstRowEnums[colIndex]; ok { // 映射中文
				// 检查value的类型并转换为字符串
				var valueKey string
				switch v := value.(type) {
				case int:
					valueKey = strconv.Itoa(v)
				case uint:
					valueKey = strconv.FormatUint(uint64(v), 10)
				}
				value = innerMap.(map[string]interface{})[valueKey]
			}
			if _, ok := firstRowDateFormat[colIndex]; ok { // 格式化日期
				switch firstRowDateFormat[colIndex] {
				case "DateTime":
					value = value.(time.Time).Format(time.DateTime)
				case "DateOnly":
					value = value.(time.Time).Format(time.DateOnly)
				case "TimeOnly":
					value = value.(time.Time).Format(time.TimeOnly)
				default:
					value = value.(time.Time).Format(time.DateTime)
				}
			}
			contentItemList = append(contentItemList, value)
		}
		mergedSlice = append(mergedSlice, contentItemList)
	}

	// 3-新建Excel
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
		}
	}()
	sheetName := "Sheet1"

	// 3-1 设置首行冻结
	_ = f.SetPanes("Sheet1", &excelize.Panes{
		Freeze:      true,         // 开启冻结功能
		Split:       false,        // 不使用拆分视图功能
		XSplit:      0,            // X轴（水平方向）上不冻结，所以设为0
		YSplit:      1,            // Y轴（垂直方向）上冻结首行，所以设置为1。这意味着在第一行下方冻结
		TopLeftCell: "A2",         // 视图冻结后，最上方可见的首个单元格设置为A2，与YSplit结合，实际冻结顶端行
		ActivePane:  "bottomLeft", // 当前活动的窗格设置为左下方
	})

	// 3-2 设置首行样式
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

	// 3-3 设置内容样式
	contentRowStyle, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left", // 水平方向上文本左对齐
		},
	})

	// 3-4 填充数据和样式
	for rowIndex, rowData := range mergedSlice {
		cellNames, _ := generateCellNames(rowIndex+1, len(firstRowKeys))
		startCell, _ := excelize.JoinCellName("A", rowIndex+1)
		_ = f.SetSheetRow(sheetName, startCell, &rowData) // 填充数据
		style := contentRowStyle
		if rowIndex == 0 {
			style = firstRowStyle
		}
		_ = f.SetCellStyle(sheetName, cellNames[0], cellNames[len(cellNames)-1], style) // 填充样式

	}

	// 4-保存文件
	filePath = global.GGB_CONFIG.Excel.OutputDir + uuid.New().String() + ".xlsx"
	err = f.SaveAs(filePath)
	if err != nil {
		global.GGB_LOG.Error(fmt.Sprintf("保存%s的新Excel失败", templateName), zap.Error(err))
	}

	return filePath, nil
}

// generateCellNames 根据给定的行号与单元格数量，生成一个含有 Excel 单元格名称的字符串切片。
//
// 此函数首先创建一个空的字符串切片以保存生成的单元格名称。接着，通过遍历指定次数 n，
// 使用 excelize 库的 ColumnNumberToName 函数将每一次迭代的列号转换为 Excel 中的列字母标识（例如，1 转换为 "A"，2 转换为 "B" 等）。
// 再利用 fmt.Sprintf 合并列字母与行号，以生成单元格名称（如 "A1", "B1" 等）。最终，这些单元格名称将被收集并
// 存放于一个切片中，该切片随后被返回作为函数结果。
//
// 参数:
//
//	rowNum: 指定希望生成单元格名称所在的行号。
//	n: 需要生成的单元格数量。
//
// 返回值：
//
//	([]string, error): 第一个返回值为生成的单元格名称切片，第二个返回值为在执行过程中可能出现的错误。
//	                   正常情况下，错误返回值为 nil，仅当列编号转换失败时才会返回 non-nil 错误。
//		               例如，generateCellNames(1, 3) 将返回 ["A1", "B1", "C1"], nil
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

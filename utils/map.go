package utils

import "reflect"

// ExcludeNestedFields 递归展开结构体并排除指定的多个嵌套字段
func ExcludeNestedFields(obj interface{}, excludeFields []string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := structToMap(reflect.ValueOf(obj), reflect.TypeOf(obj), excludeFields, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// structToMap 递归转换结构体为map
func structToMap(v reflect.Value, t reflect.Type, excludeFields []string, result map[string]interface{}) error {
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		jsonTag := field.Tag.Get("json")

		// 检查是否要排除的字段
		if shouldExcludeField(field.Name, excludeFields) || jsonTag == "-" {
			continue
		}

		// 递归展开匿名字段
		if field.Anonymous {
			err := structToMap(fieldValue, fieldValue.Type(), excludeFields, result)
			if err != nil {
				return err
			}
			continue
		}

		// 使用 JSON 标签作为键
		if jsonTag != "" && jsonTag != "-" {
			result[jsonTag] = fieldValue.Interface()
		} else {
			result[field.Name] = fieldValue.Interface()
		}
	}

	return nil
}

// shouldExcludeField 检查字段是否在排除列表中
func shouldExcludeField(fieldName string, excludeFields []string) bool {
	for _, excludeField := range excludeFields {
		if fieldName == excludeField {
			return true
		}
	}
	return false
}

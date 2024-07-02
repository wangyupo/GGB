package utils

import (
	"encoding/json"
	"fmt"
)

// Struct2Json 结构体转json
func Struct2Json(obj interface{}) string {
	str, err := json.Marshal(obj)
	if err != nil {
		panic(fmt.Sprintf("[Struct2Json]转换异常: %v", err))
	}
	return string(str)
}

// Json2Struct json转结构体
func Json2Struct(str string, obj interface{}) {
	err := json.Unmarshal([]byte(str), obj)
	if err != nil {
		panic(fmt.Sprintf("[Json2Struct]转换异常: %v", err))
	}
}

// JsonI2Struct json interface转为结构体
func JsonI2Struct(str interface{}, obj interface{}) {
	JsonStr := str.(string)
	Json2Struct(JsonStr, obj)
}

// Json2Map json转map
func Json2Map(jsonStr string) (m map[string]string, err error) {
	err = json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return nil, err
	}
	return
}

// Map2Json map转json
func Map2Json(m map[string]string) (string, error) {
	result, err := json.Marshal(m)
	if err != nil {
		return "", nil
	}
	return string(result), nil
}

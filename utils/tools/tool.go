package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// JsonString2Map - 将json的字符串转为map
func JsonString2Map(str string) (map[string]interface{}, error) {
	mapData := make(map[string]interface{})
	if str == "" {
		return mapData, nil
	}
	if err := json.Unmarshal([]byte(str), &mapData); err != nil {
		return nil, err
	}
	return mapData, nil
}

// Number2String - 数值类型转整数的字符串
func Number2String(val interface{}) string {
	switch val.(type) {
	case float64, float32:
		return fmt.Sprintf("%0.f", val)
	case int, int64, int8, int32, int16, uint:
		return fmt.Sprintf("%d", val)
	case string:
		return fmt.Sprintf("%s", val)
	}
	return ""
}

// Struct2Map - 结构体转map key tag名称 value 对应字段值
func Struct2Map(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, errors.New(fmt.Sprintf("ToMap only accepts struct or struct pointer; got %T", v))
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

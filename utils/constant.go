package utils

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
)

//struct反射
func structCom(obj interface{}) (reflect.Type, reflect.Value) {
	t := reflect.TypeOf(obj) // 获取 obj 的类型信息
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr { // 如果是指针，则获取其所指向的元素
		t = t.Elem()
		v = v.Elem()
	}
	return t, v
}

//获取struct 不是空值的字段共和
func StructLenNonull(obj interface{}) int {
	t, v := structCom(obj)
	count := 0
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			if v.Field(i).Interface() != "" {
				count++
			}
		}
	}
	return count
}

//struct 转为map map[string]interface{}
func Struct2MapInterface(obj interface{}) map[string]interface{} {
	t, v := structCom(obj)
	var data = make(map[string]interface{})
	if t.Kind() == reflect.Struct { // 只有结构体可以获取其字段信息
		for i := 0; i < t.NumField(); i++ {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	return data
}

//struct 转为map map[string]string
func Struct2MapString(obj interface{}) map[string]string {
	t, v := structCom(obj)
	var data = make(map[string]string)
	if t.Kind() == reflect.Struct { // 只有结构体可以获取其字段信息
		for i := 0; i < t.NumField(); i++ {
			data[t.Field(i).Name] = v.Field(i).String()
		}

	}
	return data
}

//struct转json,并格式化
func Struct2json(obj interface{}) string {
	bs, _ := json.Marshal(obj)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	return out.String()
}

//字符串拼接  StringBuilder([]string{a,b,c})
func StringBuilder(p []string) string {
	var b strings.Builder
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}

//切片去重
//使用map主键唯一性
func RemoveRepMap(arr []string) []string {
	resArr := make([]string, 0)
	tmpMap := make(map[string]struct{})
	for _, val := range arr {
		if _, ok := tmpMap[val]; !ok {
			resArr = append(resArr, val)
			tmpMap[val] = struct{}{}
		}

	}
	return resArr
}

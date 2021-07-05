/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package structureconv

import (
	"reflect"
)

// 结构体转map
func StructToMapByNoBlank(obj interface{}) map[string]interface{} {
	if reflect.TypeOf(obj).Kind() != reflect.Struct {
		return nil
	}
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj2.Field(i).CanInterface() && !IsBlank(obj2.Field(i).Interface()) {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

// 是否为空
func IsBlank(obj interface{}) bool {
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	case reflect.Invalid:
		return true
	}

	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

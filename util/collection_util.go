package util

import (
	"reflect"
)

// 判断elem是否在collection中，collection支持的类型arrary,slice,map
func Contains(collection interface{}, elem interface{}) bool {
	targetValue := reflect.ValueOf(collection)
	switch reflect.TypeOf(collection).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == elem {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(elem)).IsValid() {
			return true
		}
	}

	return false
}

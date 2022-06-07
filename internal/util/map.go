package util

import (
	"reflect"
)

// ToMap uses json tags on struct fields to decide which fields
// to add to the returned map.
// inspired by https://stackoverflow.com/a/23598731/10633734
func ToMap(in interface{}) map[string]interface{} {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return out
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		if tagv := fi.Tag.Get("json"); tagv != "" {
			out[tagv] = v.Field(i).Interface()
		}
	}

	return out
}

package godf

import (
	"fmt"
	"reflect"
)

func AutoCast(data interface{}) any {
	switch reflect.ValueOf(data).Kind() {
	case reflect.Float64:
		return CastFloat64(data)
	case reflect.Int:
		return CastInt(data)
	case reflect.String:
		return CastString(data)
	case reflect.Slice:
		sliceType := reflect.TypeOf(data).Elem()
		switch sliceType.Kind() {
		case reflect.Float64:
			return CastArrayFloat64(data)
		case reflect.Int:
			return CastArrayInt(data)
		case reflect.String:
			return CastArrayString(data)
		default:
			panic(fmt.Sprintf("Datatype %s not supported yet", sliceType.Kind()))
		}
	default:
		panic(fmt.Sprintf("Datatype %s not supported yet", reflect.TypeOf(data).Kind()))
	}
}

func CastString(data interface{}) string {
	return reflect.ValueOf(data).String()
}

func CastFloat64(data interface{}) float64 {
	return reflect.ValueOf(data).Float()
}

func CastInt(data interface{}) int {
	return int(reflect.ValueOf(data).Int())
}

func CastArrayFloat64(data interface{}) []float64 {
	return reflect.ValueOf(data).Interface().([]float64)
}

func CastArrayString(data interface{}) []string {
	return reflect.ValueOf(data).Interface().([]string)
}

func CastArrayInt(data interface{}) []int {
	return reflect.ValueOf(data).Interface().([]int)
}

func CastHeaders(headers []string) []interface{} {
	var intfHeaders []interface{}
	for _, s := range headers {
		intfHeaders = append(intfHeaders, interface{}(s))
	}

	return intfHeaders
}

func GetDatatype(data interface{}) reflect.Kind {
	return reflect.TypeOf(data).Kind()
}

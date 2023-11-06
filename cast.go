package main

import "strconv"

func castInt(value interface{}) int {
	return value.(int)
}

func castFloat64(value interface{}) float64 {
	return value.(float64)
}

func castFloat32(value interface{}) float32 {
	return float32(value.(float64))
}

func castBool(value interface{}) bool {
	return value.(bool)
}

func autoCast(value interface{}) any {
	if f, err := strconv.ParseFloat(value.(string), 64); err == nil {
		return f
	}
	return value
}

func castString(data interface{}) string {
	if i, ok := data.(int); ok {
		return strconv.Itoa(i)
	} else if f, ok := data.(float64); ok {
		return strconv.FormatFloat(f, 'f', -1, 64)
	} else if _, ok := data.(interface{}); ok {
		return "b"
	}

	return "a"
}

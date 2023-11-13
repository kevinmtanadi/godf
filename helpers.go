package godf

import (
	"fmt"
	"math"
	"net/url"
	"reflect"
	"unicode/utf8"
)

func calculateSliceSize(slice reflect.Value) int {
	elemSize := int(slice.Type().Elem().Size())
	return slice.Len() * elemSize
}

func convertSize(size int) string {
	sizeUnit := []string{"B", "KB", "MB", "GB"}
	floatSize := float64(size)
	unit := 0
	for {
		if floatSize/1024 > 1 {
			unit += 1
			floatSize /= 1024
		} else {
			break
		}
	}

	remainder := math.Mod(floatSize, 1024)

	return fmt.Sprintf("%.3f %s", remainder, sizeUnit[unit])
}

func limitString(s string, n int) string {
	if utf8.RuneCountInString(s) <= n {
		return s
	}

	// Take only the first 'n' runes from the string
	runes := []rune(s)
	return string(runes[:n]) + "..."
}

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

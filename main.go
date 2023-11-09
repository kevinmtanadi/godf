package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	y := []float64{1.1, 1.2, 1.3, 1.4, 1.5}

	test := DataFrame(map[string]interface{}{
		"x": x,
		"y": y,
	})

	appendData := [][]float64{{0.6, 1.6}, {0.7, 1.7}}

	checkSliceType(appendData)

	test.Append(appendData)
	test.Show()

	// test.WriteCSV("test.csv")

	// df := ReadCSV("data.csv")

	// df.Show()

	// df.Show()

}

// Features to be added
// Merging df (same row || col)
// Insert new data to df (works but need better data casting)

func checkSliceType(slice interface{}) {
	sliceValue := reflect.ValueOf(slice)
	sliceType := sliceValue.Type()

	if sliceType.Kind() != reflect.Slice {
		fmt.Println("Not a slice.")
		return
	}

	// Check if it's a 1D or 2D slice
	if sliceValue.Len() > 0 {
		if reflect.ValueOf(sliceValue.Index(0).Interface()).Kind() == reflect.Slice {
			fmt.Println("2D slice")
		} else {
			fmt.Println("1D slice")
		}
	} else {
		fmt.Println("Empty slice.")
	}
}

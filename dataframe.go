package main

import (
	"fmt"
	"reflect"
)

type dataframe struct {
	headers []string
	rows    [][]interface{}
}

func DataFrame(data map[string]interface{}) *dataframe {
	df := dataframe{}

	for header := range data {
		df.headers = append(df.headers, header)
	}

	numRows := -1

	// Iterate over the data map and construct rows
	for _, v := range data {
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			sliceValue := reflect.ValueOf(v)
			if numRows == -1 {
				numRows = sliceValue.Len()
			} else if numRows != sliceValue.Len() {
				panic("All slices must have the same length.")
			}

			var row []interface{}
			for i := 0; i < sliceValue.Len(); i++ {
				row = append(row, sliceValue.Index(i).Interface())
			}
			df.rows = append(df.rows, row)
		} else {
			panic("Values in the map must be slices.")
		}
	}

	return &df
}

func (d *dataframe) Head(n ...int) {
	length := 10
	if len(n) > 0 {
		length = n[0]
	}

	for i, row := range d.rows {
		if i == 0 && d.headers != nil {
			fmt.Println(d.headers)
		}

		fmt.Println(row)
		if i == length-1 {
			break
		}
		i++
	}
}

func (d *dataframe) Enable() {

}

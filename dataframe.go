package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
)

type dataframe struct {
	headers []string
	data    [][]interface{}
}

// DataFrame initializes and returns a dataframe
//
// This function takes a map[string]interface{} as a variable
// where the string will be the header name and the interface{}
// will be the contained data
//
//	Returns a pointer to the dataframe
func DataFrame(data map[string]interface{}) *dataframe {
	df := dataframe{}
	headers := make([]string, 0)
	var rows [][]interface{}

	maxRows := 0

	for _, value := range data {
		val := reflect.ValueOf(value)
		if val.Kind() == reflect.Slice && val.Len() > maxRows {
			maxRows = val.Len()
		}
	}

	for key, value := range data {
		headers = append(headers, key)

		val := reflect.ValueOf(value)
		row := make([]interface{}, maxRows)

		for i := 0; i < maxRows; i++ {
			if i < val.Len() {
				row[i] = val.Index(i).Interface()
			} else {
				row[i] = nil
			}
		}

		rows = append(rows, row)
	}

	df.headers = headers
	df.data = rows

	return &df
}

// Transpose transpose the data not including the headers
//
//	Returns a pointer to the dataframe
func (d *dataframe) Transpose() *dataframe {
	if len(d.data) == 0 {
		return nil
	}

	df := dataframe{}

	numRows, numCols := len(d.data), len(d.data[0])

	transposed := make([][]interface{}, numCols)
	for j := 0; j < numCols; j++ {
		transposed[j] = make([]interface{}, numRows)
		for i := 0; i < numRows; i++ {
			transposed[j][i] = d.data[i][j]
		}
	}

	df.headers = d.headers
	df.data = transposed

	return &df
}

// Shape returns the shape of the data in the dataframe
//
//	Returns row, col of the table
func (d *dataframe) Shape() (row int, col int) {
	return len(d.data), len(d.data[0])
}

// Show renders the dataframe
func (d *dataframe) Show() {
	if d.data == nil {
		fmt.Println("dataframe is empty")
		return
	}
	df := d.Transpose()

	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetOutputMirror(os.Stdout)
	indexedHeader := append([]interface{}{"#"}, CastHeaders(d.headers)...)
	t.AppendHeader(table.Row(indexedHeader))

	if df.data != nil {
		for idx, row := range df.data {
			indexedRow := append([]interface{}{idx + 1}, row...)
			t.AppendRow(table.Row(indexedRow))
		}
	}

	t.Render()
}

// GetCol returns a new dataframe with selected columns
//
// Can receive single or multiple headers of type
//
//   - string : headers name
//
//     Example of usage:
//     df.GetCol("first_name", "last_name")
func (d *dataframe) GetCol(headers ...string) *dataframe {
	if len(headers) == 0 {
		panic("call of godf.GetCol requires at least one index")
	}

	var colNum []int

	for _, h := range headers {
		for i, header := range d.headers {
			if header == h {
				colNum = append(colNum, i)
			}
		}
	}

	newHeader := make([]string, len(colNum))

	df := dataframe{}

	for i, v := range colNum {
		newHeader[i] = d.headers[v]
	}

	df.headers = newHeader

	newData := [][]interface{}{}
	for _, col := range colNum {
		newRow := d.data[col]
		newData = append(newData, newRow)
	}

	df.data = newData

	return &df
}

// DropCol drop the inputted columns directly from the current dataframe
//
// Can receive single or multiple headers of type
//   - string : headers name
//   - int : headers position
//
// Using string is recommended
//
//	Example of usage:
//	df.DropCol("first_name", "last_name")
func (d *dataframe) DropCol(headers ...interface{}) {
	var colNum int

	for _, h := range headers {
		switch v := h.(type) {
		case int:
			colNum = v
		case string:
			for i, header := range d.headers {
				if header == v {
					colNum = i
				}
			}
		default:
			panic("unsupported header name datatype")
		}

		newHeader := make([]string, len(d.headers)-1)
		copy(newHeader[:colNum], d.headers[:colNum])
		copy(newHeader[colNum:], d.headers[colNum+1:])

		d.headers = newHeader

		newData := [][]interface{}{}
		for idx, col := range d.data {
			newCol := make([]interface{}, len(col))
			if idx != colNum {
				newCol = col
			}
			newData = append(newData, newCol)
		}

		d.data = newData
	}

}

// GetRow returns a new dataframe of inputted indexes
// and can handle single and multiple rows
//
//	Example of usage:
//	df.GetRow(1, 2, 3, 4, 5)
func (d *dataframe) GetRow(idx ...int) *dataframe {
	if len(idx) == 0 {
		panic("call of godf.GetRow requires at least one index")
	}

	df := dataframe{}
	df.headers = d.headers

	// result := [][]interface{}{}
	row, _ := d.Shape()
	for i := 0; i < row; i++ {
		data := []interface{}{}
		for _, v := range idx {
			data = append(data, d.data[i][v-1])
		}
		df.data = append(df.data, data)
	}

	return &df
}

// DropRow drops the inputted indexes directly from the current dataframe.
// Can handle single and multiple rows
//
//	Example of usage:
//	df.DropRow(1, 2, 3, 4, 5)
func (d *dataframe) DropRow(idx ...int) {
	if len(idx) == 0 {
		panic("call of godf.DropRow requires at least one index")
	}

	_, col := d.Shape()

	indexes := revertSlice(col, idx)
	df := d.GetRow(indexes...)
	d.data = df.data
}

// ExtractData returns the raw data as [][]interface{}
func (d *dataframe) ExtractData() [][]interface{} {
	return d.data
}

func stringify(data []interface{}) []string {
	line := make([]string, len(data))

	for idx, x := range data {
		if i, ok := x.(int); ok {
			line[idx] = strconv.Itoa(i)
		} else if f, ok := x.(float64); ok {
			line[idx] = strconv.FormatFloat(f, 'f', -1, 64)
		} else {
			line[idx] = x.(string)
		}
	}

	return line
}

// revertSlice returns an array of numbers from 1 to n
// without the numbers in the slice
func revertSlice(n int, slice []int) []int {
	excluded := make(map[int]struct{})
	for _, num := range slice {
		excluded[num] = struct{}{}
	}

	// Generate the sequence while excluding numbers in the array
	sequence := make([]int, 0)
	for i := 1; i <= n; i++ {
		_, exists := excluded[i]
		if !exists {
			sequence = append(sequence, i)
		}
	}

	return sequence
}

// Append appends new data to the current dataframe
//
//	Can receive single or multiple data
//	Single data would be in form of 1D slice
//	Multiple data would be in form of 2D slice
//
//	Example of usage:
//	df.Append([]int{1, 2, 3})
//	df.Append([][]int{{1, 2}, {3, 4}, {5, 6}})
func (d *dataframe) Append(data interface{}) {

	val := reflect.ValueOf(data)
	length := val.Len()
	colNum := len(d.data)

	if val.Len() > 0 {
		if reflect.ValueOf(val.Index(0).Interface()).Kind() == reflect.Slice {
			// multiple data inputted
			for i := 0; i < length; i++ {
				inputColNum := reflect.ValueOf(val.Index(i).Interface()).Len()
				if inputColNum != colNum {
					panic(fmt.Sprintf("number of columns on input data row %d does not match with existing data: {%d, %d}", i+1, inputColNum, colNum))
				}
				for j := 0; j < colNum; j++ {
					data := val.Index(i).Index(j).Interface()
					d.data[j] = append(d.data[j], data)
				}
			}
		} else {
			// only a single data inputted
			if length != len(d.data) {
				panic(fmt.Sprintf("number of columns on input data does not match with existing data: {%d, %d}", length, colNum))
			}

			for i := 0; i < length; i++ {
				if i < val.Len() {
					d.data[i] = append(d.data[i], val.Index(i).Interface())
				}
			}
		}
	} else {
		panic("given empty slice")
	}
}

// Merge will merge 2 dataframes vertically, appending the data
//
//	Example of usage:
//	mergedDf := df1.Merge(df2)
func (d1 *dataframe) Merge(d2 *dataframe) *dataframe {
	df := dataframe{}

	if len(d1.headers) != len(d2.headers) {
		panic(fmt.Sprintf("data length do not match {%d | %d}", len(d1.headers), len(d2.headers)))
	}

	for i := range d1.headers {
		if d1.headers[i] != d2.headers[i] {
			panic(fmt.Sprintf("headers do not match {%s | %s}", d1.headers[i], d2.headers[i]))
		}
	}

	for i, data := range d2.data {
		d1.data[i] = append(d1.data[i], data...)
	}

	df.headers = d1.headers
	df.data = d1.data

	return &df
}

// Join joins 2 dataframes horizontally
//
//	Example of usage:
//	joinedDf := df1.Join(df2)
func (d1 *dataframe) Join(d2 *dataframe) *dataframe {
	df := dataframe{}

	d1Row, _ := d1.Shape()
	d2Row, _ := d2.Shape()

	if d1Row != d2Row {
		panic(fmt.Sprintf("data length do not match {%d | %d}", d1Row, d2Row))
	}

	df.headers = d1.headers
	df.headers = append(df.headers, d2.headers...)

	df.data = d1.data
	df.data = append(df.data, d2.data...)

	return &df
}

// Limit returns a new dataframe with the first n rows
//
//	Example of usage:
//	limitedDf := df.Limit(5)
func (d *dataframe) Limit(n int) *dataframe {
	df := dataframe{}

	df.headers = d.headers
	df.data = make([][]interface{}, len(d.data))

	for i := range d.data {
		if i == n {
			break
		}

		row := d.GetRow(i + 1).ExtractData()
		for j, d := range row {
			df.data[j] = append(df.data[j], d[0])
		}
	}

	return &df
}

package main

import (
	"encoding/csv"
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

func ReadCSV(filename string) *dataframe {
	df := dataframe{}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	line := 0
	for {
		data, err := reader.Read()
		if err != nil {
			break
		}

		if line == 0 {
			df.headers = data
		} else {
			df.data = append(df.data, []interface{}{})
			for _, s := range data {
				castedData := AutoCast(s)
				df.data[line-1] = append(df.data[line-1], castedData)
			}
		}
		line++
	}

	df = *df.Transpose()
	return &df
}

func (d *dataframe) WriteCSV(path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func (d *dataframe) Shape() (row int, col int) {
	return len(d.data), len(d.data[0])
}

func (d *dataframe) Show() {
	df := d.Transpose()

	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetOutputMirror(os.Stdout)
	indexedHeader := append([]interface{}{"#"}, CastHeaders(d.headers)...)
	t.AppendHeader(table.Row(indexedHeader))

	for idx, row := range df.data {
		indexedRow := append([]interface{}{idx + 1}, row...)
		t.AppendRow(table.Row(indexedRow))
	}

	t.Render()
}

func (d *dataframe) GetCol(headers ...interface{}) *dataframe {
	if len(headers) == 0 {
		panic("GetCol requires at least one header")
	}

	var colNum []int

	for _, h := range headers {
		switch v := h.(type) {
		case int:
			colNum = append(colNum, v)
		case string:
			for i, header := range d.headers {
				if header == v {
					colNum = append(colNum, i)
				}
			}
		default:
			panic("Unsupported header name datatype")
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
			panic("Unsupported header name datatype")
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

func (d *dataframe) GetRow(idx ...int) *dataframe {
	if len(idx) == 0 {
		panic("GetRow requires at least one index")
	}

	df := dataframe{}
	df.headers = d.headers

	// result := [][]interface{}{}
	fmt.Println(idx)
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

func (d *dataframe) DropRow(idx ...int) {
	if len(idx) == 0 {
		panic("DropRow requires at least one index")
	}

	_, col := d.Shape()

	indexes := revertSlice(col, idx)
	df := d.GetRow(indexes...)
	d.data = df.data
}

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

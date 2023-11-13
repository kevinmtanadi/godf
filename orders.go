package godf

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Sort sorts the data according to a certain column. The column must be numeric
//
//	Example of usage:
//	df.Sort("x")
//	this will sort all the table, according to the column x
func (d *dataframe) Sort(header string) {
	df := d.Transpose()

	colNum := -1
	for i, h := range d.headers {
		if h == header {
			colNum = i
		}
	}
	if colNum == -1 {
		fmt.Printf("godf.Sort: column name %s not found\n", header)
		return
	}

	df.data = sortByColumn(df.data, colNum)
	df = df.Transpose()

	d.data = df.data
}

func sortByColumn(data [][]interface{}, columnIndex int) [][]interface{} {
	sort.Slice(data, func(i, j int) bool {
		valueI := data[i][columnIndex]
		valueJ := data[j][columnIndex]

		if i, ok := valueI.(int); ok {
			valueI = float64(i)
		}

		if j, ok := valueJ.(int); ok {
			valueJ = float64(j)
		}

		// Assuming the column values are of type float64 for simplicity
		return valueI.(float64) < valueJ.(float64)
	})

	return data
}

// Scramble will shuffle the table randomly. A seed can be send
// but if not, it will be randomly generated
//
//	Example of usage:
//	df.Scramble(1000)
func (d *dataframe) Scramble(seed ...int) {
	if len(seed) > 0 {
		rand.Seed(int64(seed[0]))
	} else {
		rand.Seed(time.Now().UnixNano())
	}

	df := d.Transpose()
	n := len(df.data)

	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		df.data[i], df.data[j] = df.data[j], df.data[i]
	}

	df = df.Transpose()
	d.data = df.data
}

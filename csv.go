package godf

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

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
				castedData := castDataType(s)
				df.data[line-1] = append(df.data[line-1], castedData)
			}
		}
		line++
	}

	df = *df.Transpose()
	return &df
}

func castDataType(s string) interface{} {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}

	return s
}

func (d *dataframe) WriteCSV(path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	df := d.Transpose()

	err = writer.Write(df.headers)
	if err != nil {
		panic(err)
	}

	for _, row := range df.data {
		err := writer.Write(stringify(row))
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Generated CSV file: ", path)
}

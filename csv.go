package godf

import (
	"encoding/csv"
	"fmt"
	"os"
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

package godf

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// Read CSV and parse it into a dataframe
//
//	Support reading from local file or URL
func ReadCSV(filename string) *dataframe {

	if !isURL(filename) {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		return parseCSV(reader)
	}

	resp, err := http.Get(filename)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	return parseCSV(reader)
}

func parseCSV(reader *csv.Reader) *dataframe {
	line := 0
	df := dataframe{}
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
				if s != "" {
					castedData := castDataType(s)
					df.data[line-1] = append(df.data[line-1], castedData)
				} else {
					df.data[line-1] = append(df.data[line-1], nil)
				}
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

// WriteCSV writes the dataframe to a CSV file
func (d *dataframe) WriteCSV(outputPath string) {
	file, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// df := d.Transpose()

	err = writer.Write(d.headers)
	if err != nil {
		panic(err)
	}

	for _, row := range d.data {
		err := writer.Write(stringify(row))
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Generated CSV file: ", outputPath)
}

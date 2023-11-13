package godf

import (
	"fmt"
	"reflect"
)

type Filter interface {
	applyFilter(d *dataframe) *dataframe
}

// Eq filter rows where a data in the given column is equal to given value
//
//	Example of usage:
//	df.Where(Eq{"x": 5})
//	will only get data with x = 5
type Eq map[string]interface{}

func (filter Eq) applyFilter(d *dataframe) *dataframe {
	df := dataframe{}
	df.headers = d.headers
	df.data = make([][]interface{}, len(d.data))

	for fHeader, fValue := range filter {
		if reflect.TypeOf(fValue).Kind() == reflect.Int {
			fValue = float64(fValue.(int))
		}
		for colIdx, dHeader := range d.headers {
			if fHeader == dHeader {
				for i, row := range d.data[colIdx] {
					if row.(float64) == fValue.(float64) {
						rowData := d.GetRow(i + 1)
						for i, data := range rowData.data {
							df.data[i] = append(df.data[i], data[0])
						}
					}
				}
			}
		}
	}

	return &df
}

// NotEq filter rows where a data in the given column is not equal to given value
//
//	Example of usage:
//	df.Where(NotEq{"x": 5})
//	will get all data except the data with x = 5
type NotEq map[string]interface{}

func (filter NotEq) applyFilter(d *dataframe) *dataframe {
	df := dataframe{}
	df.headers = d.headers
	df.data = make([][]interface{}, len(d.data))

	for fHeader, fValue := range filter {
		if reflect.TypeOf(fValue).Kind() == reflect.Int {
			fValue = float64(fValue.(int))
		}
		for colIdx, dHeader := range d.headers {
			if fHeader == dHeader {
				for i, row := range d.data[colIdx] {
					if row.(float64) != fValue.(float64) {
						rowData := d.GetRow(i + 1)
						for i, data := range rowData.data {
							df.data[i] = append(df.data[i], data[0])
						}
					}
				}
			}
		}
	}

	return &df
}

// GT filter rows where a data in the given column is greater than the given value
//
//	Example of usage:
//	df.Where(GT{"x": 5})
//	will only get data with x > 5
type GT map[string]interface{}

func (filter GT) applyFilter(d *dataframe) *dataframe {
	df := dataframe{}
	df.headers = d.headers
	df.data = make([][]interface{}, len(d.data))

	for fHeader, fValue := range filter {
		if reflect.TypeOf(fValue).Kind() == reflect.Int {
			fValue = float64(fValue.(int))
		}
		for colIdx, dHeader := range d.headers {
			if fHeader == dHeader {
				for i, row := range d.data[colIdx] {
					if row.(float64) > fValue.(float64) {
						rowData := d.GetRow(i + 1)
						for i, data := range rowData.data {
							df.data[i] = append(df.data[i], data[0])
						}
					}
				}
			}
		}
	}

	return &df
}

// GTE filter rows where a data in the given column is greater than or equal to the given value
//
//	Example of usage:
//	df.Where(GTE{"x": 5})
//	will only get data with x >= 5
type GTE map[string]interface{}

func (filter GTE) applyFilter(d *dataframe) *dataframe {
	df := dataframe{}
	df.headers = d.headers
	df.data = make([][]interface{}, len(d.data))

	for fHeader, fValue := range filter {
		if reflect.TypeOf(fValue).Kind() == reflect.Int {
			fValue = float64(fValue.(int))
		}
		for colIdx, dHeader := range d.headers {
			if fHeader == dHeader {
				for i, row := range d.data[colIdx] {
					if row.(float64) >= fValue.(float64) {
						rowData := d.GetRow(i + 1)
						for i, data := range rowData.data {
							df.data[i] = append(df.data[i], data[0])
						}
					}
				}
			}
		}
	}

	return &df
}

// LT filter rows where a data in the given column is less than the given value
//
//	Example of usage:
//	df.Where(LT{"x": 5})
//	will only get data with x < 5
type LT map[string]interface{}

func (filter LT) applyFilter(d *dataframe) *dataframe {
	df := dataframe{}
	df.headers = d.headers
	df.data = make([][]interface{}, len(d.data))

	for fHeader, fValue := range filter {
		if reflect.TypeOf(fValue).Kind() == reflect.Int {
			fValue = float64(fValue.(int))
		}
		for colIdx, dHeader := range d.headers {
			if fHeader == dHeader {
				for i, row := range d.data[colIdx] {
					if row.(float64) < fValue.(float64) {
						rowData := d.GetRow(i + 1)
						for i, data := range rowData.data {
							df.data[i] = append(df.data[i], data[0])
						}
					}
				}
			}
		}
	}

	return &df
}

// LTE filter rows where a data in the given column is less than or equal to the given value
//
//	Example of usage:
//	df.Where(LTE{"x": 5})
//	will only get data with x <= 5
type LTE map[string]interface{}

func (filter LTE) applyFilter(d *dataframe) *dataframe {
	df := dataframe{}
	df.headers = d.headers
	df.data = make([][]interface{}, len(d.data))

	for fHeader, fValue := range filter {
		if reflect.TypeOf(fValue).Kind() == reflect.Int {
			fValue = float64(fValue.(int))
		}
		for colIdx, dHeader := range d.headers {
			if fHeader == dHeader {
				for i, row := range d.data[colIdx] {
					if row.(float64) <= fValue.(float64) {
						rowData := d.GetRow(i + 1)
						for i, data := range rowData.data {
							df.data[i] = append(df.data[i], data[0])
						}
					}
				}
			}
		}
	}

	return &df
}

// In filter all data which value is in the given slice
//
//	Example of usage:
//	df.Where(In{"x": []int{1, 2, 3, 4}})
//	will only get data with x = [1, 2, 3, 4]
type In map[string]interface{}

func (filter In) applyFilter(d *dataframe) *dataframe {
	df := dataframe{}
	df.headers = d.headers
	df.data = make([][]interface{}, len(d.data))

	for fHeader, fValue := range filter {
		if reflect.TypeOf(fValue).Kind() != reflect.Slice {
			fmt.Println("In only support slice")
			return nil
		}
		for colIdx, dHeader := range d.headers {
			if fHeader == dHeader {
				for i, row := range d.data[colIdx] {
					if inArray(row, fValue) {
						rowData := d.GetRow(i + 1)
						for i, data := range rowData.data {
							df.data[i] = append(df.data[i], data[0])
						}
					}
				}
			}
		}
	}

	return &df
}

func inArray(value interface{}, array interface{}) bool {
	s := reflect.ValueOf(array)
	for i := 0; i < s.Len(); i++ {
		var fData float64
		data := s.Index(i).Interface()
		if reflect.TypeOf(data).Kind() == reflect.Int {
			fData = float64(data.(int))
		} else if reflect.TypeOf(data).Kind() == reflect.Float64 {
			fData = data.(float64)
		} else {
			panic("Currently In only supports int and float64")
		}

		if reflect.DeepEqual(value, fData) {
			return true
		}
	}

	return false
}

func inArrayString(needle string, haystack []string) bool {
	for _, h := range haystack {
		if needle == h {
			return true
		}
	}

	return false
}

// Where
//
//	Filter rows with given filter
//
//	Example of usage:
//	df.Where(GT{"x": 5})
//	will only get data with x > 5
func (d *dataframe) Where(filter Filter) *dataframe {
	if filter == nil {
		return nil
	}

	return filter.applyFilter(d)
}

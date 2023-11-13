package godf

import (
	"reflect"
)

type Preprocessing interface {
	Standardize()
	Normalize()
	OneHotEncode()
}

// if no headers are provided, all non string data will be standardized
func (d *dataframe) Standardize(headers ...string) *dataframe {
	df := d

	for i, header := range df.headers {
		if inArrayString(header, headers) {
			standardized := standardize(df.data[i])
			df.data[i] = []interface{}{}
			for _, s := range standardized {
				df.data[i] = append(df.data[i], s)
			}
		}
	}

	return df
}

func (d *dataframe) Normalize(headers ...string) *dataframe {
	df := d

	for i, header := range df.headers {
		if inArrayString(header, headers) {
			standardized := normalize(df.data[i])
			df.data[i] = []interface{}{}
			for _, s := range standardized {
				df.data[i] = append(df.data[i], s)
			}
		}
	}

	return df
}

// OneHotEncode will encode categorical data (string) into numerical data.
//
// If no headers given, it will one hot encode all string data
func (d *dataframe) OneHotEncode(headers ...string) {
	if len(headers) == 0 {
		// encode all string header
		for i := range d.data {
			if reflect.TypeOf(d.data[i][0]).Kind() == reflect.String {
				d.data[i] = oneHotEncode(d.data[i])
			}
		}
	} else {
		for i, h := range d.headers {
			if inArrayString(h, headers) {
				if reflect.TypeOf(d.data[i][0]).Kind() == reflect.String {
					d.data[i] = oneHotEncode(d.data[i])
				}
			}
		}
	}
}

func oneHotEncode(data []interface{}) []interface{} {
	encodeMap := make(map[string]int)

	encoded := make([]interface{}, len(data))
	for i, v := range data {
		if _, ok := encodeMap[v.(string)]; !ok {
			encodeMap[v.(string)] = len(encodeMap)
		}
		encoded[i] = encodeMap[v.(string)]
	}

	return encoded
}

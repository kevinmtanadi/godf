package godf

type Preprocessing interface {
	Standardize()
	Normalize()
	OneHotEncode()
}

// preprocessing is still not working, do not use.
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

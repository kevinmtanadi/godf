package godf

func (d *dataframe) FillNA(method func(d *dataframe)) {
	method(d)
}

func Zero(d *dataframe) {
	for i := range d.data {
		for j := range d.data[i] {
			if d.data[i][j] == nil {
				d.data[i][j] = 0
			}
		}
	}
}

func Average(d *dataframe) {
	for i := range d.data {
		for j := range d.data[i] {
			if d.data[i][j] == nil {
				d.data[i][j] = mean(convert1DFloat(d.data[i]))
			}
		}
	}
}

func Median(d *dataframe) {
	for i := range d.data {
		for j := range d.data[i] {
			if d.data[i][j] == nil {
				d.data[i][j] = median(convert1DFloat(d.data[i]))
			}
		}
	}
}

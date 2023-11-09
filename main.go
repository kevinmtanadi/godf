package main

func main() {
	// x := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	// y := []float64{1.1, 1.2, 1.3, 1.4, 1.5}

	// test := DataFrame(map[string]interface{}{
	// 	"x": x,
	// 	"y": y,
	// })

	df := ReadCSV("data.csv")

	df.Show()

	df.WriteCSV("output.csv")

	// df.Show()
}

// Features to be added
// Merging df (same row || col)
// Insert new data to df

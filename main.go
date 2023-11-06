package main

import "fmt"

func main() {
	X := []float64{0.08, 0.26, 0.45, 0.60}
	y := []float64{0, 0, 1, 1}
	df := DataFrame(map[string]interface{}{
		"X": X,
		"Y": y,
	})

	df.Head()
	fmt.Println(df)
	df.Enable()
}

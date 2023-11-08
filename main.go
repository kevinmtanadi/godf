package main

func main() {
	x := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	y := []float64{1.1, 1.2, 1.3, 1.4, 1.5}

	df := DataFrame(map[string]interface{}{"X": x, "Y": y})

	df.DropRow(2, 3, 5)
	df.Show()
}

// Buggy func
// GetRow
// DropRow

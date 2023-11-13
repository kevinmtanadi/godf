package godf

import "math"

func standardize(x []interface{}) []float64 {

	standardizedData := []float64{}

	castedData := []float64{}

	for _, v := range x {
		var casted float64
		switch w := v.(type) {
		case float64:
			casted = w
		case int:
			casted = float64(w)
		}
		castedData = append(castedData, casted)
	}

	avg := mean(castedData)
	for _, v := range castedData {
		result := (v - avg) / std(castedData, avg)
		standardizedData = append(standardizedData, result)
	}

	return standardizedData
}

func std(x []float64, mean float64) float64 {
	std := 0.

	for _, v := range x {
		std += math.Pow(v-mean, 2)
	}

	return math.Sqrt(std / float64(len(x)))
}

func normalize(x []interface{}) []float64 {
	normalizedData := []float64{}

	castedData := []float64{}

	for _, v := range x {
		var casted float64
		switch w := v.(type) {
		case float64:
			casted = w
		case int:
			casted = float64(w)
		}
		castedData = append(castedData, casted)
	}

	min := min(castedData)
	max := max(castedData)

	for _, v := range castedData {
		result := (v - min) / (max - min)
		normalizedData = append(normalizedData, result)
	}

	return normalizedData
}

func mean(x []float64) float64 {
	sum := 0.

	for _, v := range x {
		sum += v
	}

	return sum / float64(len(x))
}

func min(x []float64) float64 {
	min := math.Inf(1)
	for _, v := range x {
		if v < min {
			min = v
		}
	}

	return min
}

func max(x []float64) float64 {
	max := math.Inf(-1)
	for _, v := range x {
		if v > max {
			max = v
		}
	}

	return max
}

func sum(x []float64) float64 {
	res := 0.
	for _, v := range x {
		res += v
	}

	return res
}

func dot(x, y []float64) float64 {
	result := 0.0
	for i := 0; i < len(x); i++ {
		result += x[i] * y[i]
	}

	return result
}

func arrayMultiplication(x, y []float64) []float64 {
	if len(x) != len(y) {
		panic("ArrayMultiplication: len(x) != len(y)")
	}

	result := make([]float64, len(x))

	for i := 0; i < len(x); i++ {
		result[i] = x[i] * y[i]
	}

	return result
}

func correlation(x, y []float64) float64 {
	n := float64(len(x))

	pembilang := (n*dot(x, y) - (sum(x) * sum(y)))
	pembagi1 := n*sum(arrayMultiplication(x, x)) - math.Pow(sum(x), 2)
	pembagi2 := n*sum(arrayMultiplication(y, y)) - math.Pow(sum(y), 2)

	corr := pembilang / math.Sqrt(pembagi1*pembagi2)
	return corr
}

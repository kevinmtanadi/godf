package godf

import "math"

func standardize(x []interface{}) []float64 {

	standardizedData := []float64{}

	castedData := []float64{}

	for _, v := range x {
		casted := v.(float64)
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
		casted := v.(float64)
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

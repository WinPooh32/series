package series

import "github.com/WinPooh32/math"

func Mean(data Data) float32 {
	var mean float32
	var count int
	var items = data.Data()
	var inv = 1.0 / float32(len(items))
	for _, v := range items {
		if math.IsNaN(v) {
			continue
		}
		mean += v * inv
		count++
	}
	if count == 0 {
		return math.NaN()
	}
	return mean
}

func Sum(data Data) float32 {
	var sum float32
	var count int
	var items = data.Data()
	for _, v := range items {
		if math.IsNaN(v) {
			continue
		}
		sum += v
		count++
	}
	if count == 0 {
		return math.NaN()
	}
	return sum
}

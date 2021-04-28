package series

import (
	"github.com/WinPooh32/math"
)

func Mean(data Data) float32 {
	var (
		count int
		mean  float32
		items = data.Data()
		inv   = 1.0 / float32(len(items))
	)
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
	var (
		sum   float32
		count int
		items = data.Data()
	)
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

func Min(data Data) float32 {
	var (
		min   float32 = math.MaxFloat32
		count int
		items = data.Data()
	)
	for _, v := range items {
		if math.IsNaN(v) {
			continue
		}
		if v < min {
			min = v
		}
		count++
	}
	if count == 0 {
		return math.NaN()
	}
	return min
}

func Max(data Data) float32 {
	var (
		max   float32 = -math.MaxFloat32
		count int
		items = data.Data()
	)
	for _, v := range items {
		if math.IsNaN(v) {
			continue
		}
		if v > max {
			max = v
		}
		count++
	}
	if count == 0 {
		return math.NaN()
	}
	return max
}

package series

import (
	"github.com/WinPooh32/math"
)

// Mean returns mean of data's values.
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

// Sum returns sum of data's values.
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

// Min returns minimum value.
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

// Max returns maximum value.
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

// Std returns standard deviation.
// Normalized by n-1.
func Std(data Data) float32 {
	var (
		count  int
		items  = data.Data()
		mean   = Mean(data)
		inv    = 1.0 / float32(len(items)-1)
		stdDev float32
	)
	for _, v := range items {
		if math.IsNaN(v) {
			continue
		}
		d := v - mean
		stdDev += (d * d) * inv
		count++
	}
	if count == 0 {
		return math.NaN()
	}
	return stdDev
}

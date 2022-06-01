package series

import (
	"github.com/WinPooh32/series/math"
)

// AggregateFunc is applied aggregation function.
type AggregateFunc func(data Data) Dtype

// Mean returns mean of data's values.
func Mean(data Data) Dtype {
	var (
		count int
		mean  Dtype
		items = data.Values()
		inv   = 1.0 / Dtype(len(items))
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
func Sum(data Data) Dtype {
	var (
		sum   Dtype
		count int
		items = data.Values()
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
func Min(data Data) Dtype {
	var (
		min   Dtype = maxReal
		count int
		items = data.Values()
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
func Max(data Data) Dtype {
	var (
		max   Dtype = -maxReal
		count int
		items = data.Values()
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

// Argmin returns index of the smallest value of series data.
// If the minimum is achieved in multiple locations, the first row position is returned.
func Argmin(data Data) int {
	var (
		min   Dtype = maxReal
		pos   int
		items = data.Values()
	)
	for i, v := range items {
		if math.IsNaN(v) {
			continue
		}
		if v < min {
			min = v
			pos = i
		}
	}
	return pos
}

// Argmax returns index of the biggest value of series data.
// If the maximum is achieved in multiple locations, the first row position is returned.
func Argmax(data Data) int {
	var (
		max   Dtype = -maxReal
		pos   int   = -1
		items       = data.Values()
	)
	for i, v := range items {
		if math.IsNaN(v) {
			continue
		}
		if v > max {
			max = v
			pos = i
		}
	}
	return pos
}

// Std returns standard deviation.
// Normalized by n-1.
func Std(data Data, mean Dtype) Dtype {
	var (
		count  int
		items  = data.Values()
		inv    = 1.0 / Dtype(len(items)-1)
		stdDev Dtype
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
	return math.Sqrt(stdDev)
}

func First(data Data) Dtype {
	items := data.Values()
	if len(items) == 0 {
		return math.NaN()
	}
	return items[0]
}

func Last(data Data) Dtype {
	items := data.Values()
	if len(items) == 0 {
		return math.NaN()
	}
	return items[len(items)-1]
}

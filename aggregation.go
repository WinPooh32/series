package series

import (
	"github.com/WinPooh32/series/math"
)

// AggregateFunc is applied aggregation function.
type AggregateFunc func(data Data) dtype

// Mean returns mean of data's values.
func Mean(data Data) dtype {
	var (
		count int
		mean  dtype
		items = data.Data()
		inv   = 1.0 / dtype(len(items))
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
func Sum(data Data) dtype {
	var (
		sum   dtype
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
func Min(data Data) dtype {
	var (
		min   dtype = maxReal
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
func Max(data Data) dtype {
	var (
		max   dtype = -maxReal
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

// Argmin returns index of the smallest value of series data.
// If the minimum is achieved in multiple locations, the first row position is returned.
func Argmin(data Data) int {
	var (
		min   dtype = maxReal
		pos   int
		items = data.Data()
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
		max   dtype = -maxReal
		pos   int   = -1
		items       = data.Data()
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
func Std(data Data, mean dtype) dtype {
	var (
		count  int
		items  = data.Data()
		inv    = 1.0 / dtype(len(items)-1)
		stdDev dtype
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

func First(data Data) dtype {
	items := data.Data()
	if len(items) == 0 {
		return math.NaN()
	}
	return items[0]
}

func Last(data Data) dtype {
	items := data.Data()
	if len(items) == 0 {
		return math.NaN()
	}
	return items[len(items)-1]
}

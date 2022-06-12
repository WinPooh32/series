package series

import (
	"github.com/WinPooh32/series/math"
)

// AggregateFunc is applied aggregation function.
type AggregateFunc func(data Data) DType

// Mean returns mean of data's values.
func Mean(data Data) DType {
	var (
		count int
		mean  DType
		items = data.Values()
		inv   = 1.0 / DType(len(items))
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
func Sum(data Data) DType {
	var (
		sum   DType
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
func Min(data Data) DType {
	var (
		min   DType = maxReal
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
func Max(data Data) DType {
	var (
		max   DType = -maxReal
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

// Median returns median value of series.
// Linear interpolation is used for odd length.
func Median(data Data) DType {
	values := data.Values()

	if len(values) == 0 {
		return math.NaN()
	}

	if len(values) == 1 {
		return values[0]
	}

	if len(values)%2 == 0 {
		i := len(values) / 2
		return (values[i-1] + values[i]) / 2
	}

	return values[len(values)/2]
}

// Argmin returns offset of the smallest value of series data.
// If the minimum is achieved in multiple locations, the first row position is returned.
func Argmin(data Data) int {
	var (
		min   DType = maxReal
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

// Argmax returns offset of the biggest value of series data.
// If the maximum is achieved in multiple locations, the first row position is returned.
func Argmax(data Data) int {
	var (
		max   DType = -maxReal
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

// Variance returns variance of values.
// Ddof - Delta Degrees of Freedom. The divisor used in calculations is N - ddof,
// where N represents the number of elements.
func Variance(data Data, mean DType, ddof int) DType {
	var (
		dev   DType
		count int

		values = data.Values()
	)
	for _, v := range values {
		if math.IsNaN(v) {
			continue
		}
		d := v - mean
		dev += (d * d)
		count++
	}
	if count == 0 {
		return math.NaN()
	}
	return dev / DType(len(values)-ddof)
}

// Std returns standard deviation.
// Ddof - Delta Degrees of Freedom. The divisor used in calculations is N - ddof,
// where N represents the number of elements.
func Std(data Data, mean DType, ddof int) DType {
	return math.Sqrt(Variance(data, mean, ddof))
}

func First(data Data) DType {
	items := data.Values()
	if len(items) == 0 {
		return math.NaN()
	}
	return items[0]
}

func Last(data Data) DType {
	items := data.Values()
	if len(items) == 0 {
		return math.NaN()
	}
	return items[len(items)-1]
}

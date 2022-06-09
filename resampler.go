package series

import (
	"time"

	"github.com/WinPooh32/series/math"
)

// ResampleOrigin is the timestamp (milliseconds) on which to adjust the grouping.
// The timezone of origin must match the timezone of the index.
type ResampleOrigin int

const (
	// OriginEpoch is 1970-01-01.
	OriginEpoch ResampleOrigin = iota
	// OriginStart is the first value of the timeseries.
	OriginStart
	// OriginStartDay is the first day at midnight of the timeseries.
	OriginStartDay
)

// InterpolationMethod is the method of filling NaN values.
type InterpolationMethod int

const (
	// InterpolationLinear fills NaNs by linear interpolation method.
	InterpolationLinear InterpolationMethod = iota
	// InterpolationPad fills NaNs by existing values.
	InterpolationPad
	// InterpolationNone doesn't fill NaNs.
	InterpolationNone
)

// Resampler resamples time-series data.
// Not full groups will are filled by NaNs.
type Resampler struct {
	data   Data
	freq   int64
	origin ResampleOrigin
}

// Sum applies sum function to sample group.
func (res Resampler) Sum() Data {
	return res.downsample(Sum)
}

// Mean applies mean function to sample group.
func (res Resampler) Mean() Data {
	return res.downsample(Mean)
}

// Min applies min function to sample group.
func (res Resampler) Min() Data {
	return res.downsample(Min)
}

// Max applies max function to sample group.
func (res Resampler) Max() Data {
	return res.downsample(Max)
}

// First applies first function to sample group.
func (res Resampler) First() Data {
	return res.downsample(First)
}

// Last applies last function to sample group.
func (res Resampler) Last() Data {
	return res.downsample(Last)
}

// Apply applies custom function to sample group.
func (res Resampler) Apply(agg AggregateFunc) Data {
	return res.downsample(agg)
}

// Interpolate fills all NaNs between known values after applied upsamping.
func (res Resampler) Interpolate(method InterpolationMethod) Data {
	result := res.upsample()

	switch method {
	case InterpolationLinear:
		return result.Lerp()
	case InterpolationPad:
		return result.Pad()
	case InterpolationNone:
		return result
	default:
		return result
	}
}

func (res Resampler) upsample() Data {
	index := res.data.index
	values := res.data.values

	firstIdx := index[0]
	lastIdx := index[len(index)-1]

	oldFreq := Dtype(res.data.freq)
	newFreq := Dtype(res.freq)

	freq := math.Ceil(oldFreq / newFreq)

	newCap := int(lastIdx-firstIdx) + 1

	var (
		newIndex []int64
		newData  []Dtype
	)

	if cap(index) >= newCap {
		newIndex = index[:0]
	} else {
		newIndex = make([]int64, 0, newCap)
	}

	newIndex = res.reindex(newIndex, firstIdx, lastIdx, int(newFreq))

	if cap(values) >= newCap {
		newData = values[:0]
	} else {
		newData = make([]Dtype, 0, newCap)
	}

	newData = res.fillData(newData[:newCap], values, int(freq))

	return MakeData(res.freq, newIndex, newData)
}

func (Resampler) reindex(dst []int64, startValue, endValue int64, freq int) []int64 {
	for value := startValue; value <= endValue; value += int64(freq) {
		dst = append(dst, value)
	}
	return dst
}

func (Resampler) fillData(dst, src []Dtype, step int) []Dtype {
	// under the hood src and dst can be same array,
	// then fill dst at backward direction.
	i := len(dst) - 1
	j := len(src) - 1

	for i >= 0 && j >= 0 {
		dst[i] = src[j]

		// Fill new values by NaNs.
		next := i - step

		beg := next
		end := i

		if beg < 0 {
			beg = 0
		}

		between := dst[beg:end]

		for k := len(between) - 1; k >= 0; k-- {
			between[k] = math.NaN()
		}

		i = next
		j--
	}

	return dst
}

func (res Resampler) downsample(agg AggregateFunc) Data {
	if agg == nil {
		panic("aggregation func must not be nil!")
	}

	if len(res.data.index) == 0 {
		return res.data.Clone()
	}

	var (
		data = res.data

		// bucket is samples count of resampling group.
		bucket       = int(math.Ceil(Dtype(res.freq) / Dtype(res.data.freq)))
		bucketsTotal = int(math.Ceil(Dtype(res.data.Len()) / Dtype(bucket)))

		srcIndex = data.Index()

		aggValue = make([]Dtype, 0, bucketsTotal)
		aggIndex = make([]int64, 0, bucketsTotal)

		value Dtype
		beg   = 0
		end   = bucket
		idx   = res.align(srcIndex[0])
	)

	for i := 0; ; i++ {
		if end >= len(srcIndex) {
			end = len(srcIndex)

			if beg >= end {
				break
			}

			value = agg(data.Slice(beg, len(srcIndex)))

			aggValue = append(aggValue, value)
			aggIndex = append(aggIndex, idx)

			break
		}

		view := data.Slice(beg, end)
		value = agg(view)

		aggValue = append(aggValue, value)
		aggIndex = append(aggIndex, idx)

		idx += res.freq
		beg += bucket
		end += bucket
	}

	return MakeData(res.freq, aggIndex, aggValue)
}

func (res Resampler) align(point int64) int64 {
	var newPoint int64

	freq := res.freq

	switch res.origin {
	case OriginStart:
		newPoint = point
	case OriginEpoch:
		newPoint = freq * (point / freq)
	case OriginStartDay:
		t := time.Unix(0, point)
		dayStart := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
		point = int64(dayStart.Sub(t))
		newPoint = freq * (point / freq)
	}

	return newPoint
}

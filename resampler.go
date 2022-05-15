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

// // Interpolate fills all NaNs between known values after applied upsamping.
// func (res Resampler) Interpolate(method InterpolationMethod) Data {
// 	return
// }

func (res Resampler) downsample(agg AggregateFunc) Data {
	if agg == nil {
		panic("aggregation func must not be null!")
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

		startPointTS = srcIndex[0]
		endPointTS   = res.adjust(startPointTS)
		endPoint     = int(math.Ceil(Dtype(endPointTS) / Dtype(res.data.freq)))
	)

	var (
		value Dtype
		beg   = 0
		end   = endPoint
		idx   = srcIndex[end-1]
	)

	for i := 0; i < bucketsTotal; i++ {
		if end > len(srcIndex) {
			value = agg(data.Slice(beg, len(srcIndex)))

			aggValue = append(aggValue, value)
			aggIndex = append(aggIndex, idx)

			break
		}

		value = agg(data.Slice(beg, end))

		aggValue = append(aggValue, value)
		aggIndex = append(aggIndex, idx)

		idx += res.freq
		beg += bucket
		end += bucket
	}

	return MakeData(res.freq, aggIndex, aggValue)
}

func (res Resampler) adjust(startPoint int64) int64 {
	var nextPoint int64

	switch res.origin {
	case OriginStart:
		nextPoint = res.freq
	case OriginEpoch:
		nextPoint = (startPoint/res.freq+1)*res.freq - startPoint
	case OriginStartDay:
		t := time.Unix(0, startPoint*int64(time.Millisecond))
		dayStart := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		nextPoint = dayStart.AddDate(0, 0, 1).UnixMilli()
	}

	return nextPoint
}

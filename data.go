package series

import (
	"github.com/WinPooh32/series/math"
)

// Data is the series data container.
type Data struct {
	freq  int64
	index []int64
	data  []Dtype
}

// MakeData makes series data instance.
// freq is the size of data sample.
func MakeData(freq int64, index []int64, data []Dtype) Data {
	if len(index) != len(data) {
		panic("length of index and data must be equal")
	}
	return Data{
		freq:  freq,
		index: index,
		data:  data,
	}
}

// ArgAt returns index value at i offset.
func (d Data) ArgAt(i int) int64 {
	return d.index[i]
}

// At returns data value at i offset.
func (d Data) At(i int) Dtype {
	return d.data[i]
}

// Index returns underlying index slice.
func (d Data) Index() (index []int64) {
	return d.index
}

func (d Data) Data() (data []Dtype) {
	return d.data
}

// Len returns size of series data.
func (d Data) Len() int {
	return len(d.index)
}

// Freq returns period length of one sample.
func (d Data) Freq() int64 {
	return d.freq
}

// Slice makes slice of data.
func (d Data) Slice(l, r int) Data {
	return Data{
		d.freq,
		d.index[l:r],
		d.data[l:r],
	}
}

// Clone makes full copy of data.
func (d Data) Clone() Data {
	clone := Data{
		freq:  d.freq,
		index: append([]int64(nil), d.index...),
		data:  append([]Dtype(nil), d.data...),
	}
	return clone
}

func (d Data) Add(r Data) Data {
	// Slices prevent implicit bounds checks.
	sl := d.data
	sr := r.data

	if len(sl) != len(sr) {
		panic("sizes of data series must be equal")
	}

	for i := range sl {
		sl[i] += sr[i]
	}

	return d
}

func (d Data) Sub(r Data) Data {
	// Slices prevent implicit bounds checks.
	sl := d.data
	sr := r.data

	if len(sl) != len(sr) {
		panic("sizes of data series must be equal")
	}

	for i := range sl {
		sl[i] -= sr[i]
	}

	return d
}

func (d Data) Mul(r Data) Data {
	// Slices prevent implicit bounds checks.
	sl := d.data
	sr := r.data

	if len(sl) != len(sr) {
		panic("sizes of data series must be equal")
	}

	for i := range sl {
		sl[i] *= sr[i]
	}

	return d
}

func (d Data) Div(r Data) Data {
	// Slices prevent implicit bounds checks.
	sl := d.data
	sr := r.data

	if len(sl) != len(sr) {
		panic("sizes of data series must be equal")
	}

	for i := range sl {
		sl[i] /= sr[i]
	}

	return d
}

func (d Data) AddScalar(s Dtype) Data {
	sl := d.data
	for i := range sl {
		sl[i] += s
	}
	return d
}

func (d Data) SubScalar(s Dtype) Data {
	sl := d.data
	for i := range sl {
		sl[i] -= s
	}
	return d
}

func (d Data) MulScalar(s Dtype) Data {
	sl := d.data
	for i := range sl {
		sl[i] *= s
	}
	return d
}

func (d Data) DivScalar(s Dtype) Data {
	sl := d.data
	for i := range sl {
		sl[i] /= s
	}
	return d
}

// Log applies natural logarithm function to values of data.
func (d Data) Log() Data {
	sl := d.data
	for i, v := range sl {
		sl[i] = math.Log(v)
	}
	return d
}

// Abs replace each elemnt by their absolute value.
func (d Data) Abs() Data {
	sl := d.data
	for i, v := range sl {
		sl[i] = math.Abs(v)
	}
	return d
}

// Floor returns the greatest integer value less than or equal to x.
func (d Data) Floor() Data {
	sl := d.data
	for i, v := range sl {
		sl[i] = math.Floor(v)
	}
	return d
}

// Trunc returns the integer value of x.
func (d Data) Trunc() Data {
	sl := d.data
	for i, v := range sl {
		sl[i] = math.Trunc(v)
	}
	return d
}

// Round returns the nearest integer, rounding half away from zero.
func (d Data) Round() Data {
	sl := d.data
	for i, v := range sl {
		sl[i] = Dtype(math.Round(v))
	}
	return d
}

// RoundToEven returns the nearest integer, rounding ties to even.
func (d Data) RoundToEven() Data {
	sl := d.data
	for i, v := range sl {
		sl[i] = Dtype(math.RoundToEven(v))
	}
	return d
}

func (d Data) Ceil() Data {
	sl := d.data
	for i, v := range sl {
		sl[i] = math.Ceil(v)
	}
	return d
}

// Apply applies user's function to every value of data.
func (d Data) Apply(fn func(Dtype) Dtype) Data {
	sl := d.data
	for i, v := range sl {
		sl[i] = fn(v)
	}
	return d
}

// Reverse reverses index and data values.
func (d Data) Reverse() Data {
	return d.ArgReverse().DataReverse()
}

// Reverse reverses only index values.
func (d Data) ArgReverse() Data {
	sl := d.index

	if l := len(sl); l <= 1 {
		return d
	} else if l == 2 {
		sl[0], sl[1] = sl[1], sl[0]
		return d
	}

	half := len(sl) / 2

	left := sl[:half]
	right := sl[half:]

	l := 0
	r := len(right) - 1

	for l < len(left) && r >= 0 {
		left[l], right[r] = right[r], left[l]
		l++
		r--
	}

	return d
}

// Reverse reverses only data values.
func (d Data) DataReverse() Data {
	sl := d.data

	if l := len(sl); l <= 1 {
		return d
	} else if l == 2 {
		sl[0], sl[1] = sl[1], sl[0]
		return d
	}

	half := len(sl) / 2

	left := sl[:half]
	right := sl[half:]

	l := 0
	r := len(right) - 1

	for l < len(left) && r >= 0 {
		left[l], right[r] = right[r], left[l]
		l++
		r--
	}

	return d
}

// Fill NaN values.
func (d Data) Fillna(value Dtype, inplace bool) Data {
	var data Data
	if inplace {
		data = d
	} else {
		data = d.Clone()
	}
	dd := data.Data()
	for i, v := range dd {
		if math.IsNaN(v) {
			dd[i] = value
		}
	}
	return data
}

// Rolling provides rolling window calculations.
func (d Data) Rolling(window int) Window {
	return Window{
		len:  window,
		data: d,
	}
}

// EWM provides exponential weighted calculations.
func (d Data) EWM(atype AlphaType, param Dtype, adjust bool, ignoreNA bool) ExpWindow {
	return ExpWindow{
		data:     d,
		atype:    atype,
		param:    param,
		adjust:   adjust,
		ignoreNA: ignoreNA,
	}
}

// RollData applies custom function to rolling window of data.
// Function accepts window bounds.
func (d Data) RollData(window int, cb func(l int, r int)) {
	if len(d.data) <= window {
		cb(0, len(d.data))
	}
	for i := window; i <= len(d.data); i++ {
		cb(i-window, i)
	}
}

func (d Data) Resample(freq int64, origin ResampleOrigin) Resampler {
	if freq <= 0 {
		panic("resampling frequency must be greater than zero")
	}
	switch origin {
	case OriginEpoch, OriginStart, OriginStartDay:
	default:
		panic("unknown resampling origin type")
	}
	return Resampler{
		data:   d,
		freq:   freq,
		origin: origin,
	}
}

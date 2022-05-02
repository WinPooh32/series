package series

import "github.com/WinPooh32/math"

type Data struct {
	samplesize int64
	index      []int64
	data       []float32
}

// Index returns underlying index slice.
func (d Data) Index() (data []int64) {
	return d.index
}

// IndexAsFloat32 returns copy of underlying index slice converted to float32 array.
func (d Data) IndexAsFloat32() (index []float32) {
	index = make([]float32, len(d.index))
	for i, v := range d.index {
		index[i] = float32(v)
	}
	return index
}

// IndexAsFloat64 returns copy of underlying index slice converted to float64 array.
func (d Data) IndexAsFloat64() (index []float64) {
	index = make([]float64, len(d.index))
	for i, v := range d.index {
		index[i] = float64(v)
	}
	return index
}

func (d Data) Data() (data []float32) {
	return d.data
}

// DataAsInt64 returns copy of underlying data slice converted to int64 array.
func (d Data) DataAsInt64() (data []int64) {
	data = make([]int64, len(d.index))
	for i, v := range d.index {
		data[i] = int64(v)
	}
	return data
}

// DataAsFloat64 returns copy of underlying data slice converted to float64 array.
func (d Data) DataAsFloat64() (data []float64) {
	data = make([]float64, len(d.index))
	for i, v := range d.index {
		data[i] = float64(v)
	}
	return data
}

// Len returns size of series data.
func (d Data) Len() int {
	return len(d.index)
}

// SampleSize returns length of one sample.
func (d Data) SampleSize() int64 {
	return d.samplesize
}

// Slices makes slice of data.
func (d Data) Slice(l, r int) Data {
	return Data{
		d.samplesize,
		d.index[l:r],
		d.data[l:r],
	}
}

// Clone makes full copy of data.
func (d Data) Clone() Data {
	clone := Data{
		samplesize: d.samplesize,
		index:      append([]int64(nil), d.index...),
		data:       append([]float32(nil), d.data...),
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

func (d Data) AddScalar(s float32) Data {
	sl := d.data
	for i := range sl {
		sl[i] += s
	}
	return d
}

func (d Data) SubScalar(s float32) Data {
	sl := d.data
	for i := range sl {
		sl[i] -= s
	}
	return d
}

func (d Data) MulScalar(s float32) Data {
	sl := d.data
	for i := range sl {
		sl[i] *= s
	}
	return d
}

func (d Data) DivScalar(s float32) Data {
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

// Apply applies user's function to every value of data.
func (d Data) Apply(fn func(float32) float32) Data {
	sl := d.data
	for i, v := range sl {
		sl[i] = fn(v)
	}
	return d
}

// Rolling provides rolling window calculations.
func (d Data) Rolling(window int) Window {
	return Window{
		len:  window,
		data: d,
	}
}

// EWM provides exponential weighted calculations.
func (d Data) EWM(atype AlphaType, param float32, adjust bool, ignoreNA bool) ExpWindow {
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

func (d Data) Resample(samplesize int64) Data {
	switch {
	case samplesize == d.samplesize:
	case samplesize < d.samplesize:
		d = d.resampleLess(d, samplesize)
	default:
		d = d.resampleMore(d, samplesize)
	}
	return d
}

func (d Data) Fillna(value float32, inplace bool) Data {
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

func (d Data) resampleLess(data Data, samplesize int64) Data {
	// TODO
	panic("not implemented!")
}

func (d Data) resampleMore(data Data, samplesize int64) Data {
	index := data.index
	resIndex := data.Index()[:0]
	resData := data.Data()[:0]

	for i := 0; i < len(index); {
		beg, end := i, d.nextSample(index, i, samplesize)

		resIndex = append(resIndex, index[end-1])
		resData = append(resData, Mean(data.Slice(beg, end)))

		i = end
	}

	return MakeData(samplesize, resIndex, resData)
}

func (d Data) nextSample(index []int64, i int, samplesize int64) (end int) {
	size := len(index)
	border := index[i] + samplesize

	for i < size && index[i] < border {
		i++
		end = i
	}

	return end
}

func MakeData(samplesize int64, index []int64, data []float32) Data {
	if len(index) != len(data) {
		panic("size of index and data must be equal")
	}
	return Data{samplesize, index, data}
}

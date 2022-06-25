package series

import (
	"github.com/WinPooh32/series/math"
)

// Data is the series values container.
type Data struct {
	freq   int64
	index  []int64
	values []DType
}

// MakeData makes series data instance.
// freq is the size of values sample.
func MakeData(freq int64, index []int64, values []DType) Data {
	if len(index) != len(values) {
		panic("length of index and values must be equal")
	}
	return Data{
		freq:   freq,
		index:  index,
		values: values,
	}
}

// MakeValues makes vector of values without indices.
// Any manipulations with index will cause panic or incorrect results!
func MakeValues(values []DType) Data {
	return Data{
		freq:   0,
		index:  nil,
		values: values,
	}
}

// IndexAt returns index value at i offset.
func (d Data) IndexAt(i int) int64 {
	return d.index[i]
}

// At returns values value at i offset.
func (d Data) At(i int) DType {
	return d.values[i]
}

// Index returns underlying index values.
func (d Data) Index() (index []int64) {
	return d.index
}

// Values returns data  data values.
func (d Data) Values() (values []DType) {
	return d.values
}

// Len returns size of series values.
func (d Data) Len() int {
	return len(d.values)
}

// Freq returns period length of one sample.
func (d Data) Freq() int64 {
	return d.freq
}

// Equals tests data searies are equal to each other.
// NaN values are considered to be equal.
func (d Data) Equals(r Data, eps DType) bool {
	return d.IndexEquals(r) && d.ValuesEquals(r, eps)
}

func (d Data) IndexEquals(r Data) bool {
	valuesLeft := d.index
	valuesRight := r.index

	if len(valuesLeft) != len(valuesRight) {
		return false
	}

	for i := range valuesLeft {
		if valuesLeft[i] != valuesRight[i] {
			return false
		}
	}

	return true
}

func (d Data) ValuesEquals(r Data, eps DType) bool {
	valuesLeft := d.values
	valuesRight := r.values

	if len(valuesLeft) != len(valuesRight) {
		return false
	}

	for i := range valuesLeft {
		left := valuesLeft[i]
		right := valuesRight[i]

		nanL := math.IsNaN(left)
		nanR := math.IsNaN(right)

		nanEq := nanL && nanR

		if nanEq {
			continue
		}

		if nanL || nanR {
			return false
		} else if !fpEq(left, right, eps) {
			return false
		}
	}

	return true
}

// Slice makes valuesice of values.
func (d Data) Slice(l, r int) Data {
	return Data{
		d.freq,
		d.index[l:r],
		d.values[l:r],
	}
}

// Clone makes full copy of values.
func (d Data) Clone() Data {
	clone := Data{
		freq:   d.freq,
		index:  append([]int64(nil), d.index...),
		values: append([]DType(nil), d.values...),
	}
	return clone
}

// Resize resizes underlying arrays.
//
// New index values are filled by MaxInt64.
// New values values are filled by NaN.
func (d Data) Resize(newLen int) Data {
	if newLen < 0 {
		panic("newLen must be positive value")
	}

	oldLen := d.Len()

	switch {
	case newLen < oldLen:
		d.index = d.index[:newLen]
		d.values = d.values[:newLen]
	case newLen > oldLen:
		dt := newLen - oldLen

		for i := 0; i < dt; i++ {
			d.index = append(d.index, math.MaxInt64)
		}

		for i := 0; i < dt; i++ {
			d.values = append(d.values, math.NaN())
		}
	}

	return d
}

// Append appends new values to series values.
func (d Data) Append(r Data) Data {
	d.index = append(d.index, r.index...)
	d.values = append(d.values, r.values...)
	return d
}

func (d Data) Add(r Data) Data {
	// Slices prevent implicit bounds checks.
	values := d.values
	sr := r.values

	if len(values) != len(sr) {
		panic("sizes of values series must be equal")
	}

	for i := range values {
		values[i] += sr[i]
	}

	return d
}

func (d Data) Sub(r Data) Data {
	// Slices prevent implicit bounds checks.
	values := d.values
	sr := r.values

	if len(values) != len(sr) {
		panic("sizes of values series must be equal")
	}

	for i := range values {
		values[i] -= sr[i]
	}

	return d
}

func (d Data) Mul(r Data) Data {
	// Slices prevent implicit bounds checks.
	values := d.values
	sr := r.values

	if len(values) != len(sr) {
		panic("sizes of values series must be equal")
	}

	for i := range values {
		values[i] *= sr[i]
	}

	return d
}

func (d Data) Div(r Data) Data {
	// Slices prevent implicit bounds checks.
	values := d.values
	sr := r.values

	if len(values) != len(sr) {
		panic("sizes of values series must be equal")
	}

	for i := range values {
		values[i] /= sr[i]
	}

	return d
}

func (d Data) Mod(r Data) Data {
	// Slices prevent implicit bounds checks.
	values := d.values
	sr := r.values

	if len(values) != len(sr) {
		panic("sizes of values series must be equal")
	}

	for i, v := range values {
		values[i] = math.Mod(v, sr[i])
	}

	return d
}

func (d Data) Max(r Data) Data {
	// Slices prevent implicit bounds checks.
	values := d.values
	sr := r.values

	if len(values) != len(sr) {
		panic("sizes of values series must be equal")
	}

	for i, v := range values {
		values[i] = math.Max(v, sr[i])
	}

	return d
}

func (d Data) Min(r Data) Data {
	// Slices prevent implicit bounds checks.
	values := d.values
	sr := r.values

	if len(values) != len(sr) {
		panic("sizes of values series must be equal")
	}

	for i, v := range values {
		values[i] = math.Min(v, sr[i])
	}

	return d
}

func (d Data) AddScalar(s DType) Data {
	values := d.values
	for i := range values {
		values[i] += s
	}
	return d
}

func (d Data) SubScalar(s DType) Data {
	values := d.values
	for i := range values {
		values[i] -= s
	}
	return d
}

func (d Data) MulScalar(s DType) Data {
	values := d.values
	for i := range values {
		values[i] *= s
	}
	return d
}

func (d Data) DivScalar(s DType) Data {
	values := d.values
	for i := range values {
		values[i] /= s
	}
	return d
}

func (d Data) Sign() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Copysign(1, v)
	}
	return d
}

func (d Data) Sin() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Sin(v)
	}
	return d
}

func (d Data) Asin() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Asin(v)
	}
	return d
}

func (d Data) Cos() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Cos(v)
	}
	return d
}

func (d Data) Acos() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Acos(v)
	}
	return d
}

func (d Data) Tan() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Tan(v)
	}
	return d
}

func (d Data) Atan() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Atan(v)
	}
	return d
}

// Pow applies x**y, the base-x exponential of y.
func (d Data) Pow(exp DType) Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Pow(v, exp)
	}
	return d
}

// Pow10 applies 10**e, the base-10 exponential of e.
func (d Data) Pow10() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Pow10(int(v))
	}
	return d
}

// Sqr applies x**2, the base-x exponential of 2.
func (d Data) Sqr() Data {
	values := d.values
	for i, v := range values {
		values[i] *= v
	}
	return d
}

// Exp applies e**x, the base-e exponential of x.
func (d Data) Exp() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Exp(v)
	}
	return d
}

// Exp2 applies 2**x, the base-2 exponential of x.
func (d Data) Exp2() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Exp2(v)
	}
	return d
}

// Log applies natural logarithm function to values of values.
func (d Data) Log() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Log(v)
	}
	return d
}

// Log2 applies Log2(x).
func (d Data) Log2() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Log2(v)
	}
	return d
}

// Abs replace each elemnt by their absolute value.
func (d Data) Abs() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Abs(v)
	}
	return d
}

// Floor returns the greatest integer value less than or equal to x.
func (d Data) Floor() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Floor(v)
	}
	return d
}

// Trunc returns the integer value of x.
func (d Data) Trunc() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Trunc(v)
	}
	return d
}

// Round returns the nearest integer, rounding half away from zero.
func (d Data) Round() Data {
	values := d.values
	for i, v := range values {
		values[i] = DType(math.Round(v))
	}
	return d
}

// RoundToEven returns the nearest integer, rounding ties to even.
func (d Data) RoundToEven() Data {
	values := d.values
	for i, v := range values {
		values[i] = DType(math.RoundToEven(v))
	}
	return d
}

func (d Data) Ceil() Data {
	values := d.values
	for i, v := range values {
		values[i] = math.Ceil(v)
	}
	return d
}

// Apply applies user's function to every value of values.
func (d Data) Apply(fn func(DType) DType) Data {
	values := d.values
	for i, v := range values {
		values[i] = fn(v)
	}
	return d
}

// Reverse reverses index and values values.
func (d Data) Reverse() Data {
	return d.IndexReverse().DataReverse()
}

// Reverse reverses only index values.
func (d Data) IndexReverse() Data {
	values := d.index

	if l := len(values); l <= 1 {
		return d
	} else if l == 2 {
		values[0], values[1] = values[1], values[0]
		return d
	}

	half := len(values) / 2

	left := values[:half]
	right := values[half:]

	l := 0
	r := len(right) - 1

	for l < len(left) && r >= 0 {
		left[l], right[r] = right[r], left[l]
		l++
		r--
	}

	return d
}

// Reverse reverses only values values.
func (d Data) DataReverse() Data {
	values := d.values

	if l := len(values); l <= 1 {
		return d
	} else if l == 2 {
		values[0], values[1] = values[1], values[0]
		return d
	}

	half := len(values) / 2

	left := values[:half]
	right := values[half:]

	l := 0
	r := len(right) - 1

	for l < len(left) && r >= 0 {
		left[l], right[r] = right[r], left[l]
		l++
		r--
	}

	return d
}

// Fillna fills NaN values.
func (d Data) Fillna(value DType) Data {
	values := d.Values()
	for i, v := range values {
		if math.IsNaN(v) {
			values[i] = value
		}
	}
	return d
}

// Pad fills NaNs by known previous values.
func (d Data) Pad() Data {
	values := d.Values()
	gg := math.NaN()
	for i, v := range values {
		if math.IsNaN(v) {
			if !math.IsNaN(gg) {
				values[i] = gg
			}
		} else {
			gg = v
		}
	}
	return d
}

// Lerp fills NaNs between known values by linear interpolation method.
func (d Data) Lerp() Data {
	values := d.values

	if len(values) == 0 {
		return d
	}

	fill := func(y []DType, k, b DType) {
		for x := range y {
			y[x] = k*DType(x+1) + b
		}
	}

	var beg, end int

	// Find first non-NaN value.
	for i := 0; ; i++ {
		if v := values[i]; !math.IsNaN(v) {
			beg = i
			break
		}
		if i >= len(values) {
			// All values are NaNs.
			// Exit.
			return d
		}
	}

	var left, right DType

	left = values[beg]

	for i := beg + 1; i < len(values); i++ {
		val := values[i]

		if math.IsNaN(val) {
			continue
		}

		end = i
		right = val

		if dst := end - beg; dst >= 2 {
			line := values[beg+1 : end]
			k := (right - left) / DType(dst)
			b := left
			fill(line, k, b)
		}

		beg = end
		left = right
	}

	return d
}

// Diff calculates the difference of a series values elements.
func (d Data) Diff(periods int) Data {
	values := d.Values()

	if periods < 0 {
		panic("period must be positive value")
	} else if periods == 0 {
		return d
	}

	var naVals []DType

	if len(values) > periods {
		lv := values[:len(values)-periods]
		rv := values[periods:]

		for i := len(rv) - 1; i >= 0; i-- {
			rv[i] -= lv[i]
		}

		naVals = values[:periods]
	} else {
		naVals = values
	}

	for i := range naVals {
		naVals[i] = math.NaN()
	}

	return d
}

// Shift shifts values by specified periods count.
func (d Data) Shift(periods int) Data {
	if periods == 0 {
		return d
	}

	values := d.Values()

	var (
		naVals []DType
		dst    []DType
		src    []DType
	)

	if shlen := int(math.Abs(DType(periods))); shlen < len(values) {
		if periods > 0 {
			naVals = values[:shlen]
			dst = values[shlen:]
			src = values
		} else {
			naVals = values[len(values)-shlen:]
			dst = values[:len(values)-shlen]
			src = values[shlen:]
		}

		copy(dst, src)
	} else {
		naVals = values
	}

	for i := range naVals {
		naVals[i] = math.NaN()
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
func (d Data) EWM(atype AlphaType, param DType, adjust bool, ignoreNA bool) ExpWindow {
	return ExpWindow{
		data:     d,
		atype:    atype,
		param:    param,
		adjust:   adjust,
		ignoreNA: ignoreNA,
	}
}

// RollData applies custom function to rolling window of values.
// Function accepts window bounds.
func (d Data) RollData(window int, cb func(l int, r int)) {
	if len(d.values) <= window {
		cb(0, len(d.values))
	}
	for i := window; i <= len(d.values); i++ {
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

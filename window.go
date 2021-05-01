package series

import "github.com/WinPooh32/math"

type Window struct {
	len  int
	data Data
}

func (w Window) Sum() Data {
	return w.apply(Sum)
}

func (w Window) Mean() Data {
	return w.apply(Mean)
}

func (w Window) Min() Data {
	return w.apply(Min)
}

func (w Window) Max() Data {
	return w.apply(Max)
}

func (w Window) Diff() Data {
	return w.applyDiff()
}

func (w Window) Shift() Data {
	return w.applyShift()
}

func (w Window) apply(f func(Data) float32) Data {
	var (
		clone  = w.data.Clone()
		data   = clone.Data()
		period = w.len
	)

	for i := 0; i < w.len-1; i++ {
		data[i] = math.NaN()
	}

	w.data.RollData(period, func(l int, r int) {
		var slice = w.data.Slice(l, r)
		data[r-1] = f(slice)
	})

	return clone
}

func (w Window) applyDiff() Data {
	var (
		clone  = w.data.Clone()
		data   = clone.Data()
		period = w.len

		orig = w.data.Data()
	)

	var n = period - 1

	for i := n; i < len(data); i++ {
		data[i] -= orig[i-n]
	}

	for i := 0; i < n; i++ {
		data[i] = math.NaN()
	}

	return clone
}

func (w Window) applyShift() Data {
	var (
		clone  = w.data.Clone()
		data   = clone.Data()
		period = w.len
	)

	dst := data[period:]
	src := data
	copy(dst, src)

	for i := 0; i < period; i++ {
		data[i] = math.NaN()
	}

	return clone
}

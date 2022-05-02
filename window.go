package series

import "github.com/WinPooh32/math"

type Window struct {
	len  int
	data Data
}

func (w Window) Sum() Data {
	return w.Apply(Sum)
}

func (w Window) Mean() Data {
	return w.Apply(Mean)
}

func (w Window) Min() Data {
	return w.Apply(Min)
}

func (w Window) Max() Data {
	return w.Apply(Max)
}

func (w Window) Std(ma Data) Data {
	return w.applyStd(ma)
}

func (w Window) Diff() Data {
	return w.applyDiff()
}

func (w Window) Shift() Data {
	switch {
	case w.len == 0:
		return w.data
	case w.len > 0:
		return w.applyShiftPisitive()
	default:
		return w.applyShiftNegative()
	}
}

func (w Window) Apply(f func(Data) float32) Data {
	var (
		clone  = w.data.Clone()
		data   = clone.Data()
		period = w.len
	)

	for i := 0; i < w.len-1; i++ {
		data[i] = math.NaN()
	}

	w.data.RollData(period, func(l int, r int) {
		slice := w.data.Slice(l, r)
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

	total := period - 1

	for i := total; i < len(data); i++ {
		data[i] -= orig[i-total]
	}

	for i := 0; i < total; i++ {
		data[i] = math.NaN()
	}

	return clone
}

func (w Window) applyShiftPisitive() Data {
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

func (w Window) applyShiftNegative() Data {
	var (
		clone  = w.data.Clone()
		data   = clone.Data()
		period = -(w.len)
	)

	dst := data[:len(data)-period]
	src := data[period:]
	copy(dst, src)

	for i := len(data) - period; i < len(data); i++ {
		data[i] = math.NaN()
	}

	return clone
}

func (w Window) applyStd(ma Data) Data {
	var (
		clone  = w.data.Clone()
		data   = clone.Data()
		period = w.len
	)

	total := period - 1

	for i := total; i < len(data); i++ {
		p := i + 1
		v := w.data.Slice(p-period, p)
		data[i] = Std(v, ma.data[p-1])
	}

	for i := 0; i < total; i++ {
		data[i] = math.NaN()
	}

	return clone
}

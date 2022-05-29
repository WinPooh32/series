package series

import "github.com/WinPooh32/series/math"

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

func (w Window) Apply(agg AggregateFunc) Data {
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
		data[r-1] = agg(slice)
	})

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
		data[i] = Std(v, ma.values[p-1])
	}

	for i := 0; i < total; i++ {
		data[i] = math.NaN()
	}

	return clone
}

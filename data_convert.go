package series

// IndexAsInt32 returns copy of underlying index slice converted to int32 array.
func (d Data) IndexAsInt32() (index []int32) {
	index = make([]int32, len(d.index))
	for i, v := range d.index {
		index[i] = int32(v)
	}
	return index
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

// DataAsInt32 returns copy of underlying data slice converted to int32 array.
func (d Data) DataAsInt32() (data []int32) {
	data = make([]int32, len(d.data))
	for i, v := range d.data {
		data[i] = int32(v)
	}
	return data
}

// DataAsInt64 returns copy of underlying data slice converted to float32 array.
func (d Data) DataAsInt64() (data []int64) {
	data = make([]int64, len(d.data))
	for i, v := range d.data {
		data[i] = int64(v)
	}
	return data
}

// DataAsFloat32 returns copy of underlying data slice converted to float32 array.
func (d Data) DataAsFloat32() (data []float32) {
	data = make([]float32, len(d.data))
	for i, v := range d.data {
		data[i] = float32(v)
	}
	return data
}

// DataAsFloat64 returns copy of underlying data slice converted to float64 array.
func (d Data) DataAsFloat64() (data []float64) {
	data = make([]float64, len(d.data))
	for i, v := range d.data {
		data[i] = float64(v)
	}
	return data
}

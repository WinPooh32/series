package series

import (
	"reflect"
	"testing"

	"github.com/WinPooh32/math"
)

var NaN = math.NaN()

func TestWindow_Sum(t *testing.T) {

	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"odd length",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []float32{1, 2, 3, 4, 5}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5}, []float32{NaN, NaN, 6, 9, 12}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Sum(); !reflect.DeepEqual(
				got.Slice(got.Len()-w.len, got.Len()),
				tt.want.Slice(tt.want.Len()-w.len, tt.want.Len()),
			) {
				t.Errorf("Window.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_Mean(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"odd length",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []float32{1, 2, 3, 4, 5}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5}, []float32{NaN, NaN, 2, 3, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Mean(); !reflect.DeepEqual(
				got.Slice(got.Len()-w.len, got.Len()),
				tt.want.Slice(tt.want.Len()-w.len, tt.want.Len()),
			) {
				t.Errorf("Window.Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_applyDiff(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"odd length",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 1, 2, 3, 5, 8}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{NaN, NaN, NaN, 2, 4, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Diff(); !reflect.DeepEqual(
				got.Slice(got.Len()-w.len, got.Len()),
				tt.want.Slice(tt.want.Len()-w.len, tt.want.Len()),
			) {
				t.Errorf("Window.Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_applyShift(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"odd length",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 1, 2, 3, 5, 8}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{NaN, NaN, NaN, 1, 1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Shift(); !reflect.DeepEqual(
				got.Slice(got.Len()-w.len, got.Len()),
				tt.want.Slice(tt.want.Len()-w.len, tt.want.Len()),
			) {
				t.Errorf("Window.Shift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_Min(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"odd length",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []float32{4, 3, 5, 2, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5}, []float32{NaN, NaN, 3, 2, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Min(); !reflect.DeepEqual(
				got.Slice(got.Len()-w.len, got.Len()),
				tt.want.Slice(tt.want.Len()-w.len, tt.want.Len()),
			) {
				t.Errorf("Window.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_Max(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"odd length",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []float32{4, 3, 5, 2, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5}, []float32{NaN, NaN, 5, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Max(); !reflect.DeepEqual(
				got.Slice(got.Len()-w.len, got.Len()),
				tt.want.Slice(tt.want.Len()-w.len, tt.want.Len()),
			) {
				t.Errorf("Window.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_Fillna(t *testing.T) {
	type fields struct {
		len  int
		data Data
	}
	type args struct {
		value   float32
		inplace bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"odd length",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []float32{NaN, NaN, 5, 2, NaN}),
			},
			args{
				inplace: false,
				value:   0,
			},
			MakeData(1, []int64{1, 2, 3, 4, 5}, []float32{0, 0, 5, 2, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Fillna(tt.args.value, tt.args.inplace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Window.Fillna() = %v, want %v", got, tt.want)
			}
		})
	}
}

package series

import (
	"testing"

	"github.com/WinPooh32/series/math"
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
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{1, 2, 3, 4, 5}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{NaN, NaN, 6, 9, 12}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Sum(); !got.Equals(tt.want, Eps) {
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
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{1, 2, 3, 4, 5}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{NaN, NaN, 2, 3, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Mean(); !got.Equals(tt.want, Eps) {
				t.Errorf("Window.Mean() = %v, want %v", got, tt.want)
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
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{4, 3, 5, 2, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{NaN, NaN, 3, 2, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Min(); !got.Equals(tt.want, Eps) {
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
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{4, 3, 5, 2, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{NaN, NaN, 5, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Max(); !got.Equals(tt.want, Eps) {
				t.Errorf("Window.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_Std(t *testing.T) {
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
			"period = 3",
			fields{
				len:  3,
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{5, 5, 6, 7, 5, 5, 5}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7}, []DType{NaN, NaN, 0.57735026, 1, 1, 1.1547005, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}

			ma := tt.fields.data.Rolling(3).Mean()

			if got := w.Std(ma); !got.Equals(tt.want, Eps) {
				t.Errorf("Window.Std() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWindow_Median(t *testing.T) {
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
			"period = 3",
			fields{
				len:  3,
				data: MakeData(1, []int64{0, 1, 2, 3, 4}, []DType{0, 1, 2, 3, 4}),
			},
			MakeData(1, []int64{0, 1, 2, 3, 4}, []DType{NaN, NaN, 1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Window{
				len:  tt.fields.len,
				data: tt.fields.data,
			}
			if got := w.Median(); !got.Equals(tt.want, Eps) {
				t.Errorf("Window.Std() = %v, want %v", got, tt.want)
			}
		})
	}
}

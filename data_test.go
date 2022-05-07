package series

import (
	"reflect"
	"testing"
	"time"
)

func TestData_Resample(t *testing.T) {
	const (
		second = int64((1 * time.Second) / time.Millisecond)
		minute = int64((1 * time.Minute) / time.Millisecond)
	)

	dayStart := time.Date(2022, 5, 7, 0, 0, 0, 0, time.UTC).UnixMilli()

	type fields struct {
		samplesize int64
		index      []int64
		data       []float32
	}
	type args struct {
		samplesize int64
		origin     ResampleOrigin
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 2, 3, 4, 5, 6}},
			args{
				samplesize: 2,
				origin:     OriginStart,
			},
			MakeData(2, []int64{2, 4, 6}, []float32{3, 7, 11}),
		},
		{
			"odd length",
			fields{1, []int64{1, 2, 3, 4, 5, 6, 7}, []float32{1, 2, 3, 4, 5, 6, 7}},
			args{
				samplesize: 2,
				origin:     OriginStart,
			},
			MakeData(2, []int64{2, 4, 6, 8}, []float32{3, 7, 11, 7}),
		},
		{
			"even length minutes freq",
			fields{
				1 * minute,
				[]int64{
					1 * minute,
					2 * minute,
					3 * minute,
					4 * minute,
					5 * minute,
					6 * minute,
				},
				[]float32{1, 2, 3, 4, 5, 6},
			},
			args{
				samplesize: 2 * minute,
				origin:     OriginStart,
			},
			MakeData(
				2*minute,
				[]int64{
					2 * minute,
					4 * minute,
					6 * minute,
				},
				[]float32{3, 7, 11},
			),
		},
		{
			"odd length minutes freq",
			fields{
				1 * minute,
				[]int64{
					1 * minute,
					2 * minute,
					3 * minute,
					4 * minute,
					5 * minute,
					6 * minute,
					7 * minute,
				},
				[]float32{1, 2, 3, 4, 5, 6, 7},
			},
			args{
				samplesize: 2 * minute,
				origin:     OriginStart,
			},
			MakeData(
				2*minute,
				[]int64{
					2 * minute,
					4 * minute,
					6 * minute,
					8 * minute,
				},
				[]float32{3, 7, 11, 7},
			),
		},
		{
			"even length resample 1 min to 1min 30sec",
			fields{
				1 * minute,
				[]int64{
					1 * minute,
					2 * minute,
					3 * minute,
					4 * minute,
					5 * minute,
					6 * minute,
				},
				[]float32{1, 2, 3, 4, 5, 6},
			},
			args{
				samplesize: 1*minute + 30*second,
				origin:     OriginStart,
			},
			MakeData(
				1*minute+30*second,
				[]int64{
					120 * second,
					210 * second,
					300 * second,
				},
				[]float32{3, 7, 11},
			),
		},
		{
			"odd length resample 1 min to 1min 30sec",
			fields{
				1 * minute,
				[]int64{
					1 * minute,
					2 * minute,
					3 * minute,
					4 * minute,
					5 * minute,
					6 * minute,
					7 * minute,
				},
				[]float32{1, 2, 3, 4, 5, 6, 7},
			},
			args{
				samplesize: 1*minute + 30*second,
				origin:     OriginStart,
			},
			MakeData(
				1*minute+30*second,
				[]int64{
					120 * second,
					210 * second,
					300 * second,
					390 * second,
				},
				[]float32{3, 7, 11, 7},
			),
		},
		{
			"even length minutes freq origin epoch",
			fields{
				1 * minute,
				[]int64{
					dayStart + 1*minute,
					dayStart + 2*minute,
					dayStart + 3*minute,
					dayStart + 4*minute,
					dayStart + 5*minute,
					dayStart + 6*minute,
				},
				[]float32{1, 2, 3, 4, 5, 6},
			},
			args{
				samplesize: 2 * minute,
				origin:     OriginEpoch,
			},
			MakeData(
				2*minute,
				[]int64{
					dayStart + 2*minute,
					dayStart + 4*minute,
					dayStart + 6*minute,
				},
				[]float32{3, 7, 11},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:  tt.fields.samplesize,
				index: tt.fields.index,
				data:  tt.fields.data,
			}
			if got := d.Resample(tt.args.samplesize, OriginStart).Sum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.Resample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Rolling(t *testing.T) {
	type fields struct {
		samplesize int64
		index      []int64
		data       []float32
	}
	type args struct {
		window int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Window
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:  tt.fields.samplesize,
				index: tt.fields.index,
				data:  tt.fields.data,
			}
			if got := d.Rolling(tt.args.window); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.Rolling() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Add(t *testing.T) {
	type fields struct {
		samplesize int64
		index      []int64
		data       []float32
	}
	type args struct {
		r Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 2, 3, 4, 5, 6}},
			args{
				MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 2, 3, 4, 5, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{2, 4, 6, 8, 10, 12}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:  tt.fields.samplesize,
				index: tt.fields.index,
				data:  tt.fields.data,
			}
			if got := d.Add(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Sub(t *testing.T) {
	type fields struct {
		samplesize int64
		index      []int64
		data       []float32
	}
	type args struct {
		r Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 2, 3, 4, 5, 6}},
			args{
				MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 2, 3, 4, 5, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{0, 0, 0, 0, 0, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:  tt.fields.samplesize,
				index: tt.fields.index,
				data:  tt.fields.data,
			}
			if got := d.Sub(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Mul(t *testing.T) {
	type fields struct {
		samplesize int64
		index      []int64
		data       []float32
	}
	type args struct {
		r Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 2, 3, 4, 5, 6}},
			args{
				MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 2, 3, 4, 5, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 4, 9, 16, 25, 36}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:  tt.fields.samplesize,
				index: tt.fields.index,
				data:  tt.fields.data,
			}
			if got := d.Mul(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Div(t *testing.T) {
	type fields struct {
		samplesize int64
		index      []int64
		data       []float32
	}
	type args struct {
		r Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 2, 3, 4, 5, 6}},
			args{
				MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 2, 3, 4, 5, 6}),
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 1, 1, 1, 1, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:  tt.fields.samplesize,
				index: tt.fields.index,
				data:  tt.fields.data,
			}
			if got := d.Div(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_AddScalar(t *testing.T) {
	type fields struct {
		samplesize int64
		index      []int64
		data       []float32
	}
	type args struct {
		s float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 2, 3, 4, 5, 6}},
			args{
				4,
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{5, 6, 7, 8, 9, 10}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:  tt.fields.samplesize,
				index: tt.fields.index,
				data:  tt.fields.data,
			}
			if got := d.AddScalar(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.AddScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_SubScalar(t *testing.T) {
	type fields struct {
		samplesize int64
		index      []int64
		data       []float32
	}
	type args struct {
		s float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 2, 3, 4, 5, 6}},
			args{
				4,
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{-3, -2, -1, 0, 1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:  tt.fields.samplesize,
				index: tt.fields.index,
				data:  tt.fields.data,
			}
			if got := d.SubScalar(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.SubScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_MulScalar(t *testing.T) {
	type fields struct {
		samplesize int64
		index      []int64
		data       []float32
	}
	type args struct {
		s float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 2, 3, 4, 5, 6}},
			args{
				4,
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{4, 8, 12, 16, 20, 24}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:  tt.fields.samplesize,
				index: tt.fields.index,
				data:  tt.fields.data,
			}
			if got := d.MulScalar(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.MulScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_DivScalar(t *testing.T) {
	type fields struct {
		samplesize int64
		index      []int64
		data       []float32
	}
	type args struct {
		s float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Data
	}{
		{
			"even length",
			fields{1, []int64{1, 2, 3, 4, 5, 6}, []float32{2, 4, 8, 12, 14, 16}},
			args{
				2,
			},
			MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []float32{1, 2, 4, 6, 7, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:  tt.fields.samplesize,
				index: tt.fields.index,
				data:  tt.fields.data,
			}
			if got := d.DivScalar(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.DivScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Fillna(t *testing.T) {
	type fields struct {
		samplesize int64
		index      []int64
		data       []float32
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
			name: "simple fillna",
			fields: fields{
				samplesize: 1,
				index:      []int64{1, 2, 3, 4, 5},
				data:       []float32{NaN, NaN, 5, 2, NaN},
			},
			args: args{
				value:   0,
				inplace: false,
			},
			want: MakeData(1, []int64{1, 2, 3, 4, 5}, []float32{0, 0, 5, 2, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Data{
				freq:  tt.fields.samplesize,
				index: tt.fields.index,
				data:  tt.fields.data,
			}
			if got := d.Fillna(tt.args.value, tt.args.inplace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.Fillna() = %v, want %v", got, tt.want)
			}
		})
	}
}

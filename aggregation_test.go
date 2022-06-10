package series

import (
	"math"
	"reflect"
	"testing"
)

func TestMean(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mean(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMed(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		{
			name: "len = 0",
			args: args{
				data: MakeData(1, []int64{}, []DType{}),
			},
			want: NaN,
		},
		{
			name: "len = 1",
			args: args{
				data: MakeData(1, []int64{1}, []DType{1}),
			},
			want: 1,
		},
		{
			name: "even len = 4",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{1, 2, 3, 4}),
			},
			want: 2.5,
		},
		{
			name: "odd len = 5",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{1, 2, 3, 4, 5}),
			},
			want: 3,
		},
		{
			name: "even len = 4, with negative values",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{-1, -3, 3, 4}),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args.data); !reflect.DeepEqual(got, tt.want) && !(math.IsNaN(got) || math.IsNaN(tt.want)) {
				t.Errorf("Med() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgmin(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Argmin(tt.args.data); got != tt.want {
				t.Errorf("Argmin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgmax(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Argmax(tt.args.data); got != tt.want {
				t.Errorf("Argmax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStd(t *testing.T) {
	type args struct {
		data Data
		mean DType
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Std(tt.args.data, tt.args.mean); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Std() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirst(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := First(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLast(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Last(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

package series

import (
	"fmt"
	"testing"

	"github.com/WinPooh32/series/math"
)

func BenchmarkPow3_2(b *testing.B) {
	var s DType = 0
	for i := 0; i < b.N; i++ {
		s += math.Pow(DType(i), 1.5)
	}
	fmt.Println(s)
}

func BenchmarkSqrtMulasPow3_2(b *testing.B) {
	var s DType = 0
	for i := 0; i < b.N; i++ {
		s += math.Sqrt(DType(i)) * DType(i)
	}
	fmt.Println(s)
}

func BenchmarkSumRef(b *testing.B) {
	var s DType = 0
	for i := 0; i < b.N; i++ {
		s += 1.0
	}
	fmt.Println(s)
}

func TestMean(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		{
			name: "empty",
			args: args{
				data: MakeData(1, []int64{}, []DType{}),
			},
			want: NaN,
		},
		{
			name: "all NaN",
			args: args{
				data: MakeData(1, []int64{1, 2, 3}, []DType{NaN, NaN, NaN}),
			},
			want: NaN,
		},
		{
			name: "one value",
			args: args{
				data: MakeData(1, []int64{1}, []DType{2}),
			},
			want: 2,
		},
		{
			name: "simple",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{2, 2, 6, 6}),
			},
			want: 4,
		},
		{
			name: "NaNs leading",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{NaN, NaN, 2, 6}),
			},
			want: 4,
		},
		{
			name: "nonNaNs leading",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{2, 6, NaN, NaN}),
			},
			want: 4,
		},
		{
			name: "NaN at mid 1",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{2, 6, NaN, 2, 6}),
			},
			want: 4,
		},
		{
			name: "NaN at mid 2",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{2, 6, NaN, NaN, 2, 6}),
			},
			want: 4,
		},
		{
			name: "NaN bounds",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{NaN, 2, 6, 2, 6, NaN}),
			},
			want: 4,
		},
		{
			name: "NaN bounds 2",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7, 8}, []DType{NaN, NaN, 2, 6, 2, 6, NaN, NaN}),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Mean(tt.args.data)
			if !(IsNA(got) && IsNA(tt.want)) && !fpEq(got, tt.want, EpsFp32) {
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
		{
			name: "empty",
			args: args{
				data: MakeData(1, []int64{}, []DType{}),
			},
			want: NaN,
		},
		{
			name: "all NaN",
			args: args{
				data: MakeData(1, []int64{1, 2, 3}, []DType{NaN, NaN, NaN}),
			},
			want: NaN,
		},
		{
			name: "one value",
			args: args{
				data: MakeData(1, []int64{1}, []DType{2}),
			},
			want: 2,
		},
		{
			name: "simple",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{2, 2, 6, 6}),
			},
			want: 16,
		},
		{
			name: "NaNs leading",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{NaN, NaN, 2, 6}),
			},
			want: 8,
		},
		{
			name: "nonNaNs leading",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{2, 6, NaN, NaN}),
			},
			want: 8,
		},
		{
			name: "NaN at mid 1",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{2, 6, NaN, 2, 6}),
			},
			want: 16,
		},
		{
			name: "NaN at mid 2",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{2, 6, NaN, NaN, 2, 6}),
			},
			want: 16,
		},
		{
			name: "NaN bounds",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{NaN, 2, 6, 2, 6, NaN}),
			},
			want: 16,
		},
		{
			name: "NaN bounds 2",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7, 8}, []DType{NaN, NaN, 2, 6, 2, 6, NaN, NaN}),
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.args.data)
			if !(IsNA(got) && IsNA(tt.want)) && !fpEq(got, tt.want, EpsFp32) {
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
		{
			name: "empty",
			args: args{
				data: MakeData(1, []int64{}, []DType{}),
			},
			want: NaN,
		},
		{
			name: "all NaN",
			args: args{
				data: MakeData(1, []int64{1, 2, 3}, []DType{NaN, NaN, NaN}),
			},
			want: NaN,
		},
		{
			name: "one value",
			args: args{
				data: MakeData(1, []int64{1}, []DType{2}),
			},
			want: 2,
		},
		{
			name: "simple",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{2, 2, 6, 6}),
			},
			want: 2,
		},
		{
			name: "NaNs leading",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{NaN, NaN, 2, 6}),
			},
			want: 2,
		},
		{
			name: "nonNaNs leading",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{2, 6, NaN, NaN}),
			},
			want: 2,
		},
		{
			name: "NaN at mid 1",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{2, 6, NaN, 2, 6}),
			},
			want: 2,
		},
		{
			name: "NaN at mid 2",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{2, 6, NaN, NaN, 2, 6}),
			},
			want: 2,
		},
		{
			name: "NaN bounds",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{NaN, 2, 6, 2, 6, NaN}),
			},
			want: 2,
		},
		{
			name: "NaN bounds 2",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7, 8}, []DType{NaN, NaN, 2, 6, 2, 6, NaN, NaN}),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Min(tt.args.data)
			if !(IsNA(got) && IsNA(tt.want)) && !fpEq(got, tt.want, EpsFp32) {
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
		{
			name: "empty",
			args: args{
				data: MakeData(1, []int64{}, []DType{}),
			},
			want: NaN,
		},
		{
			name: "all NaN",
			args: args{
				data: MakeData(1, []int64{1, 2, 3}, []DType{NaN, NaN, NaN}),
			},
			want: NaN,
		},
		{
			name: "one value",
			args: args{
				data: MakeData(1, []int64{1}, []DType{2}),
			},
			want: 2,
		},
		{
			name: "simple",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{2, 2, 6, 6}),
			},
			want: 6,
		},
		{
			name: "NaNs leading",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{NaN, NaN, 2, 6}),
			},
			want: 6,
		},
		{
			name: "nonNaNs leading",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{2, 6, NaN, NaN}),
			},
			want: 6,
		},
		{
			name: "NaN at mid 1",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{2, 6, NaN, 2, 6}),
			},
			want: 6,
		},
		{
			name: "NaN at mid 2",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{2, 6, NaN, NaN, 2, 6}),
			},
			want: 6,
		},
		{
			name: "NaN bounds",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{NaN, 2, 6, 2, 6, NaN}),
			},
			want: 6,
		},
		{
			name: "NaN bounds 2",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7, 8}, []DType{NaN, NaN, 2, 6, 2, 6, NaN, NaN}),
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Max(tt.args.data)
			if !(IsNA(got) && IsNA(tt.want)) && !fpEq(got, tt.want, EpsFp32) {
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
			got := Median(tt.args.data)
			if !(IsNA(got) && IsNA(tt.want)) && !fpEq(got, tt.want, EpsFp32) {
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
		{
			name: "empty",
			args: args{
				data: MakeData(1, []int64{}, []DType{}),
			},
			want: -1,
		},
		{
			name: "all NaN",
			args: args{
				data: MakeData(1, []int64{1, 2, 3}, []DType{NaN, NaN, NaN}),
			},
			want: -1,
		},
		{
			name: "one value",
			args: args{
				data: MakeData(1, []int64{1}, []DType{2}),
			},
			want: 0,
		},
		{
			name: "simple",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{2, 2, 6, 6}),
			},
			want: 0,
		},
		{
			name: "NaNs leading",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{NaN, NaN, 2, 6}),
			},
			want: 2,
		},
		{
			name: "nonNaNs leading",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4}, []DType{2, 6, NaN, NaN}),
			},
			want: 0,
		},
		{
			name: "NaN at mid 1",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5}, []DType{2, 6, NaN, 2, 6}),
			},
			want: 0,
		},
		{
			name: "NaN at mid 2",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{2, 6, NaN, NaN, 2, 6}),
			},
			want: 0,
		},
		{
			name: "NaN bounds",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6}, []DType{NaN, 2, 6, 2, 6, NaN}),
			},
			want: 1,
		},
		{
			name: "NaN bounds 2",
			args: args{
				data: MakeData(1, []int64{1, 2, 3, 4, 5, 6, 7, 8}, []DType{NaN, NaN, 2, 6, 2, 6, NaN, NaN}),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Argmin(tt.args.data)
			if got != tt.want {
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
			if got := Std(tt.args.data, tt.args.mean, 1); !fpEq(got, tt.want, EpsFp32) {
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
			if got := First(tt.args.data); !fpEq(got, tt.want, EpsFp32) {
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
			if got := Last(tt.args.data); !fpEq(got, tt.want, EpsFp32) {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkew(t *testing.T) {
	type args struct {
		data Data
	}
	tests := []struct {
		name string
		args args
		want DType
	}{
		{
			name: "simple_with_nan",
			args: args{
				data: MakeValues([]DType{1, NaN, 1, 2}),
			},
			want: 1.7320508075688787,
		},
		{
			name: "simple_with",
			args: args{
				data: MakeValues([]DType{1, 1, 2}),
			},
			want: 1.7320508075688787,
		},
		{
			name: "simple_with3",
			args: args{
				data: MakeValues([]DType{1, 1, 1, 1, 2}),
			},
			want: 2.2360679774997902,
		},
		{
			name: "simple_2233",
			args: args{
				data: MakeValues([]DType{2, 2, 3, 3}),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Skew(tt.args.data)
			if !fpEq(got, tt.want, 1e-4) {
				t.Errorf("Skew() = %v, want %v", got, tt.want)
			}
		})
	}
}

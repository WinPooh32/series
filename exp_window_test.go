package series

import (
	"testing"

	"github.com/WinPooh32/series/math"
)

func TestExpWindow_Mean(t *testing.T) {
	type fields struct {
		data     Data
		atype    AlphaType
		param    dtype
		adjust   bool
		ignoreNA bool
	}
	tests := []struct {
		name   string
		fields fields
		want   Data
	}{
		{
			"simple",
			fields{
				data: MakeData(
					1,
					[]int64{1, 2, 3, 4, 5},
					[]dtype{0, 1, 2, NaN, 4},
				),
				atype:    AlphaCom,
				param:    0.5,
				adjust:   true,
				ignoreNA: false,
			},
			MakeData(
				1,
				[]int64{1, 2, 3, 4, 5},
				[]dtype{0, 0.6923077, 1.575, 1.575, 3.198347},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := ExpWindow{
				data:     tt.fields.data,
				atype:    tt.fields.atype,
				param:    tt.fields.param,
				adjust:   tt.fields.adjust,
				ignoreNA: tt.fields.ignoreNA,
			}

			equal := true
			got := w.Mean()

			for i, v := range got.data {
				if math.Abs(v-tt.want.data[i]) > 0.001 {
					equal = false
					break
				}
			}

			if !equal {
				t.Errorf("ExpWindow.Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

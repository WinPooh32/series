//go:build !series_f32

package series

import (
	"github.com/WinPooh32/series/math"
)

const Eps = 10e-8

const maxReal = math.MaxFloat64

type Dtype = float64

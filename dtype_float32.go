//go:build series_f32

package series

import (
	"github.com/WinPooh32/series/math"
)

type DType = float32

const (
	EpsFp32 = 1e-7
	EpsFp64 = 1e-14
	Eps     = EpsFp32
)

const maxReal = math.MaxFloat32

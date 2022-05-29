package series

import (
	"sort"

	"github.com/WinPooh32/series/math"
)

type sortable Data

func (x sortable) Len() int { return len(x.values) }

func (x sortable) Less(i, j int) bool {
	return x.values[i] < x.values[j] || (math.IsNaN(x.values[i]) && !math.IsNaN(x.values[j]))
}

func (x sortable) Swap(i, j int) {
	x.values[i], x.values[j] = x.values[j], x.values[i]
	x.index[i], x.index[j] = x.index[j], x.index[i]
}

type argSortable Data

func (x argSortable) Len() int { return len(x.index) }

func (x argSortable) Less(i, j int) bool { return x.index[i] < x.index[j] }

func (x argSortable) Swap(i, j int) {
	x.values[i], x.values[j] = x.values[j], x.values[i]
	x.index[i], x.index[j] = x.index[j], x.index[i]
}

// ArgSort sorts data's index.
func (d Data) ArgSort() {
	sort.Sort(argSortable(d))
}

// Sort sorts data.
func (d Data) Sort() {
	sort.Sort(sortable(d))
}

// ArgSortStable sorts data's index using stable sort algorithm.
func (d Data) ArgSortStable() {
	sort.Stable(argSortable(d))
}

// SortStable sorts data's index using stable sort algorithm.
func (d Data) SortStable() {
	sort.Stable(sortable(d))
}

package series

import (
	"sort"

	"github.com/WinPooh32/series/math"
)

type sortable Data

func (x sortable) Len() int { return len(x.data) }

func (x sortable) Less(i, j int) bool {
	return x.data[i] < x.data[j] || (math.IsNaN(x.data[i]) && !math.IsNaN(x.data[j]))
}

func (x sortable) Swap(i, j int) {
	x.data[i], x.data[j] = x.data[j], x.data[i]
	x.index[i], x.index[j] = x.index[j], x.index[i]
}

type argSortable Data

func (x argSortable) Len() int { return len(x.index) }

func (x argSortable) Less(i, j int) bool { return x.index[i] < x.index[j] }

func (x argSortable) Swap(i, j int) {
	x.data[i], x.data[j] = x.data[j], x.data[i]
	x.index[i], x.index[j] = x.index[j], x.index[i]
}

// ArgSort sorts data's' index.
func (d Data) ArgSort() {
	sort.Sort(argSortable(d))
}

// Sort sorts data.
func (d Data) Sort() {
	sort.Sort(sortable(d))
}

// ArgSortStable sorts data's' index using stable sort algorithm.
func (d Data) ArgSortStable() {
	sort.Stable(argSortable(d))
}

// SortStable sorts data's' index using stable sort algorithm.
func (d Data) SortStable() {
	sort.Stable(sortable(d))
}

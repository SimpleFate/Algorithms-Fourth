package ch2

import (
	"fmt"
	"math"
)

type Comparable interface {
	CompareTo(other interface{}) int
}

type SortBase interface {
	Sort(l []Comparable) []Comparable
	IsSorted(l []Comparable) bool
	less(a Comparable, b Comparable) bool
	exchange(pa *Comparable, pb *Comparable)
	StartCount()
	EndCount()
	GetCount() (comp uint64, exch uint64)
}

type Cint int

func (v Cint) CompareTo(other interface{}) int {
	if _, ok := other.(Cint); !ok {
		panic(fmt.Errorf("invalid params in CompareTo"))
	}
	o := other.(Cint) //convert type
	return int(v - o)
}

type CommonSortBase struct {
	compCount uint64
	exchCount uint64
	counting  bool
}

func (csb *CommonSortBase) Sort(l []Comparable) []Comparable {
	//to be implemented by sub struct
	return l
}

func (csb *CommonSortBase) StartCount() {
	csb.compCount = 0
	csb.exchCount = 0
	csb.counting = true
}

func (csb *CommonSortBase) EndCount() {
	csb.counting = false
}

func (csb *CommonSortBase) GetCount() (comp uint64, exch uint64) {
	return csb.compCount, csb.exchCount
}

func (csb *CommonSortBase) addComp() {
	if csb.counting {
		if csb.compCount+1 > math.MaxUint64 {
			csb.counting = false
			panic(fmt.Errorf("compare count out of range"))
		} else {
			csb.compCount++
		}
	}
}

func (csb *CommonSortBase) addExch() {
	if csb.counting {
		if csb.exchCount+1 > math.MaxUint64 {
			csb.counting = false
			panic(fmt.Errorf("compare count out of range"))
		} else {
			csb.exchCount++
		}
	}
}

func (csb *CommonSortBase) IsSorted(l []Comparable) bool {
	// for i, j  if i < j
	// l[j] less than l[i] is always false
	for i := len(l) - 1; i > 0; i-- {
		if csb.less(l[i], l[i-1]) {
			return false
		}
	}
	return true
}
func (CommonSortBase) display(l []Comparable) {
	for _, v := range l {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
func (csb *CommonSortBase) less(a Comparable, b Comparable) bool {
	csb.addComp()
	return a.CompareTo(b) < 0
}
func (csb *CommonSortBase) exchange(pa *Comparable, pb *Comparable) {
	csb.addExch()
	*pa, *pb = *pb, *pa
}

package ch2

import (
	"fmt"
)

type Comparable interface {
	CompareTo(other interface{}) int
}

type SortBase interface {
	Sort(l []Comparable) []Comparable
	IsSorted(l []Comparable) bool
	Display(l []Comparable)
	less(a Comparable, b Comparable) bool
	exchange(pa *Comparable, pb *Comparable)
}

type Cint int

func (v Cint) CompareTo(other interface{}) int {
	if _, ok := other.(Cint); !ok {
		panic(fmt.Errorf("invalid params in CompareTo"))
	}
	o := other.(Cint) //convert type
	return int(v - o)
}

func Cint2Comparable(l []Cint) []Comparable {
	res := make([]Comparable, len(l))
	for i, v := range l {
		res[i] = v
	}
	return res
}

type CommonSortBase struct {
}

func (csb *CommonSortBase) Sort(l []Comparable) []Comparable {
	//to be implemented by sub struct
	return l
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
func (CommonSortBase) Display(l []Comparable) {
	for _, v := range l {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
func (CommonSortBase) less(a Comparable, b Comparable) bool {
	return a.CompareTo(b) < 0
}
func (CommonSortBase) exchange(pa *Comparable, pb *Comparable) {
	*pa, *pb = *pb, *pa
}

func (csb *CommonSortBase) Exchange(pa *Comparable, pb *Comparable) {
	csb.exchange(pa, pb)
}

package ch2

//normal
type TopMerge struct {
	CommonSortBase
}

//allocate support memory in recursion
type TopMerge2 struct {
	CommonSortBase
}

//3 improvements for normal
type TopMerge3 struct {
	CommonSortBase
}

//TODO
type BottomMerge struct {
	CommonSortBase
}

//TODO
type InplaceMerge struct {
	CommonSortBase
}

//[l,r)
func (tm *TopMerge) divide(l []Comparable, left, right int, sup []Comparable) {
	if right-left == 1 {
		return
	}
	mid := left + (right-left)/2
	tm.divide(l, left, mid, sup)
	tm.divide(l, mid, right, sup)
	tm.merge(l, left, mid, right, sup)
}
func (tm *TopMerge) merge(l []Comparable, left, mid, right int, sup []Comparable) {
	//copy
	for i := left; i < right; i++ {
		sup[i] = l[i]
	}
	i, j, k := left, mid, left
	for i < mid && j < right {
		if tm.less(sup[j], sup[i]) {
			l[k] = sup[j]
			k++
			j++
		} else {
			l[k] = sup[i]
			k++
			i++
		}
	}
	for i == mid && k < right {
		l[k] = sup[j]
		k++
		j++
	}
	for j == right && i < mid {
		l[k] = sup[i]
		k++
		i++
	}
}

func (tm *TopMerge) Sort(l []Comparable) []Comparable {
	sup := make([]Comparable, len(l))
	tm.divide(l, 0, len(l), sup)
	return l
}

//-------------------------------------------------------------------------

//[l,r)
func (tm *TopMerge2) divide2(l []Comparable, left, right int) {
	if right-left == 1 {
		return
	}
	mid := left + (right-left)/2
	tm.divide2(l, left, mid)
	tm.divide2(l, mid, right)
	tm.merge2(l, left, mid, right)
}
func (tm *TopMerge2) merge2(l []Comparable, left, mid, right int) {
	sup := make([]Comparable, right-left)
	//copy
	for i := left; i < right; i++ {
		sup[i-left] = l[i]
	}
	i, j, k := left, mid, left
	for i < mid && j < right {
		if tm.less(sup[j-left], sup[i-left]) {
			l[k] = sup[j-left]
			k++
			j++
		} else {
			l[k] = sup[i-left]
			k++
			i++
		}
	}
	for i == mid && k < right {
		l[k] = sup[j-left]
		k++
		j++
	}
	for j == right && i < mid {
		l[k] = sup[i-left]
		k++
		i++
	}
}

func (tm *TopMerge2) Sort(l []Comparable) []Comparable {
	tm.divide2(l, 0, len(l))
	return l
}

//-------------------------------------------------------------------------

//[l,r)
func (tm *TopMerge3) divide3(l []Comparable, left, right int, sup []Comparable) int {
	if right-left <= 5 {
		tm.insert(l, left, right)
		return 0
	}
	mid := left + (right-left)/2
	tm.divide3(l, left, mid, sup)
	tm.divide3(l, mid, right, sup)

	tm.merge3(l, left, mid, right, sup)

}
func (tm *TopMerge3) merge3(l []Comparable, left, mid, right int, sup []Comparable) {
	if mid-1 >= left && !tm.less(l[mid], l[mid-1]) {
		return
	}

	//copy
	for i := left; i < right; i++ {
		sup[i] = l[i]
	}

	//compare sup and put result to l
	i, j, k := left, mid, left
	for i < mid && j < right {
		if tm.less(sup[j], sup[i]) {
			l[k] = sup[j]
			k++
			j++
		} else {
			l[k] = sup[i]
			k++
			i++
		}
	}
	for i == mid && k < right {
		l[k] = sup[j]
		k++
		j++
	}
	for j == right && i < mid {
		l[k] = sup[i]
		k++
		i++
	}
}

func (tm *TopMerge3) insert(l []Comparable, left, right int) {
	for i := left + 1; i < right; i++ {
		for j := i - 1; j >= left && tm.less(l[j+1], l[j]); j-- {
			tm.exchange(&l[j], &l[j+1])
		}
	}
}

func (tm *TopMerge3) Sort(l []Comparable) []Comparable {
	sup := make([]Comparable, len(l))
	for i, v := range l {
		sup[i] = v
	}
	tm.divide3(l, 0, len(l), sup)
	return l
}

package ch2

type SelectionSort struct {
	CommonSortBase
}

//交换次数固定 n-1
//比较次数固定 1+2+...+n = n*(n+1)/2
//不受待排元素影响
func (ss *SelectionSort) Sort(l []Comparable) []Comparable {
	for i := 0; i < len(l)-1; i++ {
		min := l[i]
		index := i
		for j := i; j < len(l); j++ {
			if ss.less(l[j], min) {
				min = l[j]
				index = j
			}
		}
		ss.exchange(&l[i], &l[index])
	}
	return l
}

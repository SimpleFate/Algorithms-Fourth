package ch2

type SelectionSort struct {
	CommonSortBase
}

//不受待排元素影响
//最好情况 比较：1+2+...+n = n*(n+1)/2  交换：n-1
//最坏情况 比较：1+2+...+n = n*(n+1)/2  交换：n-1
//不稳定  eg: 5 8 5 2 9
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
		//ss.exchange(&l[i], &l[index])
		l[i], l[index] = l[index], l[i]
	}
	return l
}

package ch2

type InsertionSort struct {
	CommonSortBase
}

//受待排数组影响，越有序表现越好
//最好情况 比较：n-1  交换：0
//最坏情况 比较：1+2+...+n-1 = n*(n-1)/2 交换： 1+2+...+n-1 = n*(n-1)/2
//稳定
func (is *InsertionSort) Sort(l []Comparable) []Comparable {
	for i := 1; i < len(l); i++ {
		for j := i - 1; j >= 0 && is.less(l[j+1], l[j]); j-- {
			//is.exchange(&l[j], &l[j+1])
			l[j], l[j+1] = l[j+1], l[j]
		}
	}
	return l
}

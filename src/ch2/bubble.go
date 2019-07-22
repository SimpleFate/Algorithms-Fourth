package ch2

type Bubble struct {
	CommonSortBase
}

//受待排数组影响，越有序表现越好, 类似插入排序
//最好情况 比较：1+2+...+n-1 = n*(n-1)/2  交换：0
//最坏情况 比较：1+2+...+n-1 = n*(n-1)/2 交换： 1+2+...+n-1 = n*(n-1)/2
//稳定
func (b *Bubble) Sort(l []Comparable) []Comparable {
	for i := len(l) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if b.less(l[j+1], l[j]) {
				b.exchange(&l[j+1], &l[j])
			}
		}
	}
	return l
}

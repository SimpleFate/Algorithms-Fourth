package main

import (
	"ch2"
	"fmt"
	"time"
)

var (
	//sorter = &ch2.SelectionSort{}
	//sorter = &ch2.InsertionSort{}
	//sorter = &ch2.Bubble{}
	//sorter = &ch2.Shell{}
	//sorter = &ch2.TopMerge{}
	//sorter = &ch2.TopMerge2{}
	sorter = &ch2.TopMerge3{}

	//nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//nums = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	nums = getRange(100)

	shuffle = true
)

func getRange(n int) []int {
	l := make([]int, n)
	for i := range l {
		l[i] = i + 1
	}
	return l
}

func sortNums() {
	cnums := make([]ch2.Cint, len(nums))
	for i, v := range nums {
		cnums[i] = ch2.Cint(v)
	}

	var n []ch2.Comparable

	if shuffle {
		n = ch2.InitCint(cnums)
	} else {
		n = ch2.Cint2Comparable(cnums)
	}

	sorter.StartCount()

	fmt.Println("befor :", n)
	n = sorter.Sort(n)
	fmt.Println("after :", n)

	sorter.EndCount()

	comp, exch := sorter.GetCount()
	fmt.Printf("compare : %d\texchange : %d\n", comp, exch)
	fmt.Println("isSorted : ", sorter.IsSorted(n))
}

var (
	markN = 10000
	markT = 100

	//markSort = &ch2.SelectionSort{} //110s N=10000 T=100
	//markSort = &ch2.InsertionSort{}  //45s N=10000 T=100
	//markSort = &ch2.Bubble{} //139s N=10000 T=100
	//markSort = &ch2.Shell{} //0.519s N=10000 T=100
	//markSort = &ch2.TopMerge{} //0.320s N=10000 T=100 normal
	//markSort = &ch2.TopMerge2{} //0.490s N=10000 T=100  allocate support memory in recursion
	markSort = &ch2.TopMerge3{} // N=10000 T=100

)

func benchMark() {
	cnums := make([]ch2.Cint, markN)
	for i := range cnums {
		cnums[i] = ch2.Cint(i)
	}
	comNums := ch2.Cint2Comparable(cnums)

	var elapse time.Duration
	failed := 0
	for i := 0; i < markT; i++ {
		comNums = ch2.ShuffleComparable(comNums)
		tmpNums := make([]ch2.Comparable, len(comNums))
		copy(tmpNums, comNums)

		before := time.Now()
		markSort.Sort(tmpNums)
		subElapse := time.Since(before)
		elapse += subElapse
		if !markSort.IsSorted(tmpNums) {
			failed++
		}
	}

	fmt.Println(elapse)
	fmt.Println("failed : ", failed)
}

func main() {
	sortNums()
	//benchMark()
}

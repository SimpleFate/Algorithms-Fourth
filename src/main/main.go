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
	sorter = &ch2.Shell{}

	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//nums = []int{10,9,8,7,6,5,4,3,2,1}

	shuffle = true
)

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
	fmt.Println("after :", sorter.Sort(n))

	sorter.EndCount()

	comp, exch := sorter.GetCount()
	fmt.Printf("compare : %d\texchange : %d\n", comp, exch)
}

var (
	markN = 10000
	markT = 100

	//markSort = &ch2.SelectionSort{} //110s N=10000 T=100
	//markSort = &ch2.InsertionSort{}  //45s N=10000 T=100
	//markSort = &ch2.Bubble{} //139s N=10000 T=100
	markSort = &ch2.Shell{} //0.519s N=10000 T=100

)

func benchMark() {
	cnums := make([]ch2.Cint, markN)
	for i := range cnums {
		cnums[i] = ch2.Cint(i)
	}
	comNums := ch2.Cint2Comparable(cnums)

	var elapse time.Duration
	for i := 0; i < markT; i++ {
		comNums = ch2.ShuffleComparable(comNums)
		tmpNums := make([]ch2.Comparable, len(comNums))
		copy(tmpNums, comNums)

		before := time.Now()
		markSort.Sort(tmpNums)
		subElapse := time.Since(before)

		elapse += subElapse
	}

	fmt.Println(elapse)
}

func main() {
	//sortNums()
	benchMark()
}

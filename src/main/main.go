package main

import (
	"ch2"
	"fmt"
)

var (
	//sorter = &ch2.SelectionSort{}
	//sorter = &ch2.InsertionSort{}
	sorter = &ch2.Bubble{}

	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//nums = []int{10,9,8,7,6,5,4,3,2,1}

	shuffle = false
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
	fmt.Printf("compare : %d\texch : %d\n", comp, exch)
}

func main() {
	sortNums()
}

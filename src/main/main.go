package main

import (
	"ch2"
	"fmt"
)

func main() {
	nums := []ch2.Cint{1, 2, 2, 3, 4, 5, 5}
	n := ch2.Cint2Comparable(nums)
	s := ch2.CommonSortBase{}

	s.Display(n)
	fmt.Println(s.Sort(n))

	s.Display(n)
	fmt.Println(s.IsSorted(n))

	s.Exchange(&n[0], &n[1])
	s.Display(n)
}

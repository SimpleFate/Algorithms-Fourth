package ch2

import (
	"math/rand"
	"time"
)

func nextInt(n int) int {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return r.Intn(n)
}

func ShuffleComparable(l []Comparable) []Comparable {
	for i := len(l) - 1; i >= 0; i-- {
		j := nextInt(i + 1)
		l[i], l[j] = l[j], l[i]
	}
	return l
}

func Cint2Comparable(l []Cint) []Comparable {
	res := make([]Comparable, len(l))
	for i, v := range l {
		res[i] = v
	}
	return res
}

func InitCint(nums []Cint) []Comparable {
	cl := Cint2Comparable(nums)
	cl = ShuffleComparable(cl)
	return cl
}

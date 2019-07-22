package ch2

type Shell struct {
	CommonSortBase
}

//h(1) = 1, h(n) = 3 * h(n-1) + 1  => h(n) = 1/2(3^n - 1)
// h -> h/3
//

func (s *Shell) Sort(l []Comparable) []Comparable {
	h := 1
	for h < len(l)/3 {
		h = 3*h + 1
	}
	for ; h >= 1; h /= 3 {
		for i := h; i < len(l); i++ {
			for j := i - h; j >= 0 && s.less(l[j+h], l[j]); j -= h {
				s.exchange(&l[j+h], &l[j])
			}
		}
	}
	return l
}

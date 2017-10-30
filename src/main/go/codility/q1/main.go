package solution

import (
	"sort"
)

type IntPair struct {
	First  int
	Second int
}

type SecondSortIntPairs []IntPair

func (sip SecondSortIntPairs) Len() int {
	return len(sip)
}

func (sip SecondSortIntPairs) Less(i, j int) bool {
	return sip[i].Second < sip[j].Second
}

func (sip SecondSortIntPairs) Swap(i, j int) {
	sip[i], sip[j] = sip[j], sip[i]
}

func (sip SecondSortIntPairs) Sort() {
	sort.Sort(sip)
}

func Solution(A []int) int {

	pairA := make(SecondSortIntPairs, len(A))
	for k, v := range A {
		pairA[k] = IntPair{
			First:  k,
			Second: v,
		}
	}
	sort.Stable(pairA)

	max := -1
	previous := pairA[0]

	for i := 1; i < len(pairA); i++ {
		current := pairA[i]
		if current.Second > previous.Second {

			nextPrevious := current

			for j := i + 1; j < len(pairA); j++ {
				if pairA[j].Second == current.Second {
					i += 1
				} else {
					break
				}
			}

			distance := Abs(pairA[i].First - previous.First)
			if max < distance {
				max = distance
			}

			previous = nextPrevious
		}
	}

	return max
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}


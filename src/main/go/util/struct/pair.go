package _struct

import "sort"

type IntPair struct {
	First  int
	Second int
}

type FirstSortIntPairs []IntPair

func (fip FirstSortIntPairs) Len() int {
	return len(fip)
}

func (fip FirstSortIntPairs) Less(i, j int) bool {
	return fip[i].First < fip[j].First
}

func (fip FirstSortIntPairs) Swap(i, j int) {
	fip[i], fip[j] = fip[j], fip[i]
}

func (fip FirstSortIntPairs) Sorted() {
	sort.Sort(fip)
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

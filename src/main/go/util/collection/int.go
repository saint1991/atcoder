package collection

import "sort"

type IntSlice sort.IntSlice
type SortedIntSlice IntSlice

func (is IntSlice) FirstIndexOf(predicate func(int) bool) int {
	for index, element := range is {
		if predicate(element) {
			return index
		}
	}
	return -1
}

func (is IntSlice) Contains(i int) bool {
	index := is.FirstIndexOf(func(e int) bool {
		return i == e
	})
	return index != -1
}

func (is IntSlice) Filter(predicate func(int) bool) IntSlice {
	ret := make([]int, 0)
	for _, i := range is {
		if predicate(i) {
			ret = append(ret, i)
		}
	}
	return ret
}

func (is IntSlice) Map(mapper func(int) int) IntSlice {
	ret := make([]int, len(is))
	for index, i := range is {
		ret[index] = mapper(i)
	}
	return ret
}

func (is IntSlice) CountIf(predicate func(int) bool) int {
	count := 0
	for _, i := range is {
		if predicate(i) {
			count++
		}
	}
	return count
}

func (is IntSlice) Sorted() SortedIntSlice {
	sort.Ints(is)
	return SortedIntSlice(is)
}

func (sis SortedIntSlice) IndexOf(i int) int {
	return sort.SearchInts(sis, i)
}

func (sis SortedIntSlice) Contains(i int) bool {
	return sis.IndexOf(i) != -1
}

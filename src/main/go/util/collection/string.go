package collection

import "sort"

type StringSlice sort.StringSlice
type SortedStringSlice StringSlice

func (ss StringSlice) FirstIndexOf(predicate func(string) bool) int {
	for index, element := range ss {
		if predicate(element) {
			return index
		}
	}
	return -1
}

func (ss StringSlice) Contains(s string) bool {
	index := ss.FirstIndexOf(func(e string) bool {
		return e == s
	})
	return index != -1
}

func (ss StringSlice) Filter(predicate func(string) bool) StringSlice {
	ret := make([]string, 0)
	for _, s := range ss {
		if predicate(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func (ss StringSlice) Map(mapper func(string) string) StringSlice {
	ret := make([]string, len(ss))
	for index, s := range ss {
		ret[index] = mapper(s)
	}
	return ret
}

func (ss StringSlice) CountIf(predicate func(string) bool) int {
	count := 0
	for _, s := range ss {
		if predicate(s) {
			count++
		}
	}
	return count
}

func (ss StringSlice) Sorted() SortedStringSlice {
	sort.Strings(ss)
	return SortedStringSlice(ss)
}

func (sss SortedStringSlice) IndexOf(s string) int {
	return sort.SearchStrings(sss, s)
}

func (sss SortedStringSlice) Contains(s string) bool {
	return sss.IndexOf(s) != -1
}

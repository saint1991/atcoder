package collection

import "sort"

type FloatSlice sort.Float64Slice
type SortedFloatSlice FloatSlice

func (fs FloatSlice) FirstIndexOf(predicate func(f float64) bool) int {
	for index, element := range fs {
		if predicate(element) {
			return index
		}
	}
	return -1
}

func (fs FloatSlice) Contains(f float64) bool {
	index := fs.FirstIndexOf(func(e float64) bool {
		return f == e
	})
	return index != -1
}

func (fs FloatSlice) Filter(predicate func(float64) bool) FloatSlice {
	ret := make([]float64, 0)
	for _, f := range fs {
		if predicate(f) {
			ret = append(ret, f)
		}
	}
	return ret
}

func (fs FloatSlice) Map(mapper func(float64) float64) FloatSlice {
	ret := make([]float64, len(fs))
	for index, f := range fs {
		ret[index] = mapper(f)
	}
	return ret
}

func (fs FloatSlice) CountIf(predicate func(float64) bool) int {
	count := 0
	for _, f := range fs {
		if predicate(f) {
			count++
		}
	}
	return count
}

func (fs FloatSlice) Sorted() SortedFloatSlice {
	sort.Float64s(fs)
	return SortedFloatSlice(fs)
}

func (sfs SortedFloatSlice) IndexOf(f float64) int {
	return sort.SearchFloat64s(sfs, f)
}

func (sfs SortedFloatSlice) Contains(f float64) bool {
	return sfs.IndexOf(f) != -1
}

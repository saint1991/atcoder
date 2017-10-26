package math

import (
	"math"
	"fmt"
)

func Pow(base int, deg uint) int {
	value := 1
	for i := 0; uint(i) < deg; i++ {
		value *= base
	}
	return value
}

func Max(is ...int) int {
	max := math.MinInt32
	for _, i := range is {
		if max < i {
			max = i
		}
	}
	return max
}

func Min(is ...int) int {
	min := math.MaxInt32
	for _, i := range is {
		if min > i {
			min = i
		}
	}
	return min
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ToBinary(num int) string {
	return fmt.Sprintf("%b", num)
}

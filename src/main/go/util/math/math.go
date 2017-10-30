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

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func IsPrime(n int) bool {
	for i := 2; i * i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return n != 1
}

func Devisors(n int) []int {
	ret := make([]int, 0)
	for i := 1; i * i <= n; i++ {
		if n % i == 0 {
			ret = append(ret, i)
			if i != n / i {
				ret = append(ret, n / i)
			}
		}
	}
	return ret
}

func PrimeFactor(n int) map[int]int {
	ret := make(map[int]int, 0)
	for i := 2; i * i <= n; i++ {
		for n % i == 0 {
			ret[i] += 1
			n /= i
		}
	}
	if n != 1 {
		ret[n] = 1
	}
	return ret
}

func ToBinary(num int) string {
	return fmt.Sprintf("%b", num)
}

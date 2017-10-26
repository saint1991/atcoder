package main

import "fmt"

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Solution(A []int) int {

	accA := make([]int, len(A))
	acc := 0
	for i, v := range A {
		acc += v
		accA[i] = acc
	}

	min := 100000000
	length := accA[len(accA)-1]
	for _, v := range accA[:len(accA) - 1] {
		rest := length - v
		diff := Abs(rest - v)
		if diff < min {
			min = diff
		}
	}
	return min
}

func main() {
	fmt.Println(Solution([]int {1000, -1000}))
}
package main

import "fmt"

func Solution(A []int, K int) []int {
	N := len(A)
	if N == 0 {
		return []int{}
	}
	after := A[len(A) - (K % N):]
	before := A[0:len(A) - (K % N)]
	return append(after, before...)
}

func main() {
	fmt.Println(Solution([]int{1}, 2))
}
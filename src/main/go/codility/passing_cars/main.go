package main

import "fmt"

func Solution(A []int) int {
	acc := 0
	total := 0

	N := len(A)
	for i := range A {
		if A[N-1-i] == 1 {
			total++
		} else {
			acc += total
			if acc > 1000000000 {
				return -1
			}
		}
	}
	return acc
}

func main() {
	fmt.Println(Solution([]int{0, 1, 0, 1, 1}))
}

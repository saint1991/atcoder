package main

import "fmt"

func Solution(A []int) int {
	N := len(A)
	flags := make([]bool, N + 1)

	for i := range flags {
		flags[i] = false
	}
	for _, v := range A {
		flags[v-1] = true
	}

	for i, flag := range flags {
		if !flag {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(Solution([]int{1, 2, 3, 5, 4}))
}
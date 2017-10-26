package main

import "fmt"

func Solution(A []int) int {
	N := len(A)

	flags := make([]bool, N)
	for i := range flags {
		flags[i] = false
	}

	for _, v := range A {
		if v > N {
			return 0
		}
		if flags[v - 1] {
			return 0
		}
		flags[v - 1] = true
	}
	return 1
}

func main() {
	fmt.Println(Solution([]int{2, 2, 1}))
}
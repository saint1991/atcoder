package main

import "fmt"

func Solution(X int, A []int) int {
	flags := make([]bool, X)
	for i := range flags {
		flags[i] = false
	}

	hitCnt := 0
	for k, v := range A {
		if !flags[v-1] {
			flags[v-1] = true
			hitCnt++
		}
		if hitCnt >= X {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println(Solution(2, []int{1, 2}))
}

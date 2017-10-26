package main

import "fmt"

func Solution4(A []int) int {

	m := make(map[int]int, 0)

	for _, v := range A {
		cnt, ok := m[v]
		if !ok {
			cnt = 0
		}
		m[v] = cnt + 1
	}

	for k, v := range m {
		if v % 2 != 0 {
			return k
		}
	}

	return 0
}

func main() {
	fmt.Println(Solution4([]int{1, 2, 3, 1, 2, 4, 3}))
}
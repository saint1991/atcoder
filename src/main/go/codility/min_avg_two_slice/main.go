package main

import "fmt"

func average(acc []int, from, to int) float64 {
	fromVal := 0
	if from > 0 {
		fromVal = acc[from - 1]
	}
	toVal := acc[to]

	return float64(toVal - fromVal) / float64(to - from + 1)
}


// O(N^2) still not solved...
func Solution(A []int) int {

	total := 0
	acc := make([]int, len(A))
	for i, v := range A {
		total += v
		acc[i] = total
	}

	min := 10000.0
	minI := -1
	for i := 0; i < len(acc) - 1; i++ {
		if float64(A[i]) > min && float64(A[i+1]) > min {
			continue
		}
		for j := i+1; j < len(acc); j++ {
			avg := average(acc, i, j)
			if avg < min {
				min = avg
				minI = i
			}
		}
	}
	return minI
}

func main() {
	fmt.Println(Solution([]int{4, 2, 2, 5, 1, 5, 8}))
}

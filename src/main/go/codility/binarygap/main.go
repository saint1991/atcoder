package main

import (
	"fmt"
)

func Solution(N int) int {
	if N <= 4 {
		return 0
	}
	bin := fmt.Sprintf("%b", N)
	fmt.Println(bin)

	max := 0
	previous := 0
	for i, c := range bin[1:] {
		if string(c) == "1" {
			gap := i - previous
			if gap > max {
				max = gap
			}
			previous = i + 1
		}
	}
	return max
}

func main() {
	fmt.Println(Solution(4))
}
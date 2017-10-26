package main

import "fmt"

func Solution(X int, Y int, D int) int {
	rest := Y - X
	cnt := rest / D
	mod := rest % D
	if mod == 0 {
		return cnt
	}
	return cnt + 1
}

func main() {
	fmt.Println(Solution(10, 85, 30))
}
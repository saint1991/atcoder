package main

import (
	"fmt"
	"math"
)

type IntStack struct {
	vector []int
}

func NewIntStack() *IntStack {
	return &IntStack{
		vector: make([]int, 0),
	}
}

func (is *IntStack) Size() int {
	return len(is.vector)
}

func (is *IntStack) IsEmpty() bool {
	return len(is.vector) == 0
}

func (is *IntStack) Top() (int, bool) {
	if is.IsEmpty() {
		return math.MaxInt32, false
	}
	return is.vector[len(is.vector)-1], true
}

func (is *IntStack) Pop() (int, bool) {
	ret, ok := is.Top()
	is.vector = is.vector[0 : len(is.vector)-1]
	return ret, ok
}

func (is *IntStack) Push(i int) {
	is.vector = append(is.vector, i)
}

func Solution(N int, A []int) []int {

	max := 0
	stack := NewIntStack()

	for _, v := range A {
		if v == N+1 {
			// construct counter here
			localCounters := make(map[int]int, 0)
			for !stack.IsEmpty() {
				value, _ := stack.Pop()
				_, exists := localCounters[value]
				if exists {
					localCounters[value] += 1
				} else {
					localCounters[value] = 1
				}
			}
			localMax := 0
			for _, cnt := range localCounters {
				if localMax < cnt {
					localMax = cnt
				}
			}
			max += localMax
		} else {
			stack.Push(v)
		}
	}

	// initialize counters to calculated maximum value so far
	counters := make([]int, N)
	for i := range counters {
		counters[i] = max
	}

	// apply all remaining elements in the stack
	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		counters[value - 1] += 1
	}

	return counters
}

func main() {
	fmt.Println(Solution(5, []int{3, 4, 4, 6, 1, 4, 4}))
}

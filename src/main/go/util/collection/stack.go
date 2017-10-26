package collection

import "math"

type Stack struct {
	vector []interface{}
}

func NewStack() *Stack {
	return &Stack{
		vector: make([]interface{}, 0),
	}
}

func (s *Stack) Size() int {
	return len(s.vector)
}

func (s *Stack) IsEmpty() bool {
	return len(s.vector) == 0
}

func (s *Stack) Push(e interface{}) {
	s.vector = append(s.vector, e)
}

func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.vector[len(s.vector)-1]
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	ret := s.Top()
	s.vector = s.vector[0 : len(s.vector)-1]
	return ret
}

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

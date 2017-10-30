package general_question

type Fibonacci struct {
	memo []int
}

func NewFibonacci(N int) Fibonacci {
	return Fibonacci{
		memo: make([]int, N),
	}
}

func (f *Fibonacci) At(n int) int {
	if n <= 1 {
		return n
	}
	if len(f.memo) > n {
		if f.memo[n] != 0 {
			return f.memo[n]
		} else {
			f.memo[n] = f.At(n-1) + f.At(n-2)
			return f.memo[n]
		}
	} else {
		return f.At(n-1) + f.At(n-2)
	}
}

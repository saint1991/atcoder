package genomic_range_query

func minIF(acc [][]int, from, to int) int {
	for factor, a := range acc {
		if a[from] != a[to] {
			return factor + 1
		}
	}
	return -1
}

func ifOf(c string) int {
	switch c {
	case "A":
		return 1
	case "C":
		return 2
	case "G":
		return 3
	case "T":
		return 4
	default:
		return -1
	}
}

func Solution(S string, P []int, Q []int) []int {

	N := len(S)

	A := 0
	C := 1
	G := 2
	T := 3

	acc := make([][]int, 4)
	acc[A] = make([]int, N)
	acc[C] = make([]int, N)
	acc[G] = make([]int, N)
	acc[T] = make([]int, N)

	accA := 0
	accC := 0
	accG := 0
	accT := 0
	for k, v := range S {
		char := string(v)
		switch char {
		case "A":
			accA += 1
			acc[A][k] = accA
		case "C":
			accC += 1
			acc[C][k] = accC
		case "G":
			accG += 1
			acc[G][k] = accG
		case "T":
			accT += 1
			acc[T][k] = accT
		}
	}

	ans := make([]int, len(P))
	for i, from := range P {
		to := Q[i]
		if from == to {
			ans[i] = ifOf(string(S[from]))
		} else {
			ans[i] = minIF(acc, from, to)
		}
	}
	return ans
}

package general_question

import "math"

type LCS struct {
	s    string
	t    string
	n    int
	m    int
	memo [][]int
}

func NewLCS(s, t string) LCS {
	N := len(s)
	M := len(t)
	memo := make([][]int, N+1)
	for i := range memo {
		memo[i] = make([]int, M+1)
	}
	memo[0][0] = 0
	return LCS{
		s:    s,
		t:    t,
		n:    N,
		m:    M,
		memo: memo,
	}
}

func max(is ...int) int {
	max := math.MinInt32
	for _, i := range is {
		if max < i {
			max = i
		}
	}
	return max
}

func (lcs *LCS) Solve() int {
	for i := 0; i < lcs.n; i++ {
		for j := 0; j < lcs.m; j++ {
			schar := string(lcs.s[i])
			tchar := string(lcs.t[j])
			if schar == tchar {
				lcs.memo[i+1][j+1] = lcs.memo[i][j] + 1
			} else {
				lcs.memo[i+1][j+1] = max(lcs.memo[i+1][j], lcs.memo[i][j+1])
			}
		}
	}
	return lcs.memo[lcs.n][lcs.m]
}

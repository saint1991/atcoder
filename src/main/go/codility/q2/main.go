package solution

func Solution(A []int) int {

	N := len(A)
	if N == 0 || N == 1 {
		return N
	}

	previousIndex2 := 0
	previousIndex1 := 0
	max := 0

	for i := previousIndex1 + 1; i < len(A); i++ {

		if previousIndex2 > previousIndex1 {
			tmp := previousIndex2
			previousIndex2 = previousIndex1
			previousIndex1 = tmp
		}

		if A[previousIndex1] != A[i] && A[previousIndex2] != A[i] {
			size := i - previousIndex2
			if max < size {
				max = size
			}
			previousIndex2 = previousIndex1
			previousIndex1 = i
		}
	}

	if A[len(A)-1] == A[previousIndex1] {
		max = len(A) - previousIndex2
	}

	return max
}

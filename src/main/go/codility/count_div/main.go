package main


func Solution(A int, B int, K int) int {
	quotientA := A / K
	quotientB := B / K

	ans := quotientB - quotientA
	if A % K == 0 {
		ans++
	}
	return ans
}

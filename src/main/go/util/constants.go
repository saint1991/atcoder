package util

const (
	LowerAlphabets = "abcdefghijklmnopqrstuvwxyz"
	UpperAlphabets = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func NewLowerAlphabetCounter() map[string]int {
	counter := make(map[string]int, 0)
	for _, c := range LowerAlphabets {
		char := string(c)
		counter[char] = 0
	}
	return counter
}

func NewUpperAlphabetCounter() map[string]int {
	counter := make(map[string]int, 0)
	for _, c := range UpperAlphabets {
		char := string(c)
		counter[char] = 0
	}
	return counter
}

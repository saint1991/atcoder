package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	maxBufferSize = 1048576
)

type Scanner struct {
	lineScanner *bufio.Scanner
	wordScanner *bufio.Scanner
}

var Stdin *Scanner

func init() {
	Stdin = NewScanner()
	Stdin.lineScanner.Buffer(make([]byte, 0), maxBufferSize)
	Stdin.wordScanner.Buffer(make([]byte, 0), maxBufferSize)
}

func NewScanner() *Scanner {
	scanner := &Scanner{
		lineScanner: bufio.NewScanner(os.Stdin),
		wordScanner: bufio.NewScanner(os.Stdin),
	}
	scanner.lineScanner.Split(bufio.ScanLines)
	scanner.wordScanner.Split(bufio.ScanWords)
	return scanner
}

func (sc *Scanner) NextLine() string {
	Stdin.lineScanner.Scan()
	return Stdin.lineScanner.Text()
}

func (sc *Scanner) Next() string {
	Stdin.wordScanner.Scan()
	return Stdin.wordScanner.Text()
}

func (sc *Scanner) NextInt() int {
	Stdin.wordScanner.Scan()
	ret, err := strconv.Atoi(Stdin.wordScanner.Text())
	if err != nil {
		panic(err)
	}
	return ret
}

func (sc *Scanner) NextNInts(N int) []int {
	ints := make([]int, N)
	for i := 0; i < N; i++ {
		ints[i] = sc.NextInt()
	}
	return ints
}

func (sc *Scanner) NextInt64() int64 {
	Stdin.wordScanner.Scan()
	ret, err := strconv.ParseInt(Stdin.wordScanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return ret
}

func (sc *Scanner) NextFloat() float64 {
	Stdin.wordScanner.Scan()
	ret, err := strconv.ParseFloat(Stdin.wordScanner.Text(), 64)
	if err != nil {
		panic(err)
	}
	return ret
}

const (
	LowerAlphabets = "abcdefghijklmnopqrstuvwxyz"
	UpperAlphabets = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func NewFlags() map[string]bool {
	ret := make(map[string]bool, 0)
	for _, alpha := range LowerAlphabets {
		ret[string(alpha)] = false
	}
	return ret
}

func IsPalindrome(flags map[string]bool) bool {
	counts := 0
	for _, flag := range flags {
		if flag {
			counts++
		}
	}
	return counts <= 1
}

func FindFirstGroupBorder(s string) int {
	maxIndex := 0
	flags := NewFlags()
	for i, c := range s {
		flags[string(c)] = !flags[string(c)]
		if IsPalindrome(flags) {
			maxIndex = i
		}
	}
	return maxIndex + 1
}

// Still Wrong Answer...
func main() {
	s := Stdin.NextLine()
	N := len(s)

	index := 0
	count := 0
	for index < N {
		current := FindFirstGroupBorder(s[index:])
		index += current
		count++
	}

	fmt.Println(count)
}

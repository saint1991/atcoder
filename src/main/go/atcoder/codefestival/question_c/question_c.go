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

type StringSlice []string

func (ss StringSlice) Contains(s string) bool {
	for _, element := range ss {
		if element == s {
			return true
		}
	}
	return false
}

func (ss StringSlice) CountIf(predicate func(string) bool) int {
	count := 0
	for _, s := range ss {
		if predicate(s) {
			count++
		}
	}
	return count
}

func (ss StringSlice) Filter(predicate func(string) bool) StringSlice {
	ret := make([]string, 0)
	for _, s := range ss {
		if predicate(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func (ss StringSlice) Map(mapper func(string) string) StringSlice {
	ret := make([]string, 0)
	for _, s := range ss {
		ret = append(ret, mapper(s))
	}
	return ret
}

type IntSlice []int

func (is IntSlice) CountIf(predicate func(int) bool) int {
	count := 0
	for _, i := range is {
		if predicate(i) {
			count++
		}
	}
	return count
}

func Pow(base int, power int) int {
	p := 1
	for i := 0; i < power; i++ {
		p *= base
	}
	return p
}

func init() {
	Stdin = NewScanner()
	Stdin.lineScanner.Buffer(make([]byte, 0), maxBufferSize)
	Stdin.wordScanner.Buffer(make([]byte, 0), maxBufferSize)
}

func main() {
	s := Stdin.NextLine()
	N := len(s)

	if N == 1 {
		fmt.Println(0)
		return
	}

	count := 0
	leftIndex := 0
	rightIndex := N - 1

	for leftIndex < rightIndex {

		left := string(s[leftIndex])
		right := string(s[rightIndex])

		if left != "x" && right == "x" {
			count++
			rightIndex--
		} else if left == "x" && right != "x" {
			count++
			leftIndex++
		} else if left == right {
			leftIndex++
			rightIndex--
		} else {
			count = -1
			break
		}

	}
	fmt.Println(count)
}

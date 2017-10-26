package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func init() {
	Stdin = NewScanner()
}

func main() {
	answerA()
}

func answerA() {
	s := Stdin.NextLine()
	if strings.Contains(s, "AC") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

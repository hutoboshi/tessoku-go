package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Io struct {
	reader    *bufio.Reader
	writer    *bufio.Writer
	tokens    []string
	nextToken int
}

func NewIo() *Io {
	return &Io{
		reader: bufio.NewReader(os.Stdin),
		writer: bufio.NewWriter(os.Stdout),
	}
}

func (io *Io) Flush() {
	err := io.writer.Flush()
	if err != nil {
		panic(err)
	}
}

func (io *Io) NextLine() string {
	var buffer []byte
	for {
		line, isPrefix, err := io.reader.ReadLine()
		if err != nil {
			panic(err)
		}
		buffer = append(buffer, line...)
		if !isPrefix {
			break
		}
	}
	return string(buffer)
}

func (io *Io) Next() string {
	for io.nextToken >= len(io.tokens) {
		line := io.NextLine()
		io.tokens = strings.Fields(line)
		io.nextToken = 0
	}
	r := io.tokens[io.nextToken]
	io.nextToken++
	return r
}

func (io *Io) NextInt() int {
	i, err := strconv.Atoi(io.Next())
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) NextFloat() float64 {
	i, err := strconv.ParseFloat(io.Next(), 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) PrintLn(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *Io) Printf(format string, a ...interface{}) {
	fmt.Fprintf(io.writer, format, a...)
}

func (io *Io) PrintIntLn(a []int) {
	b := []interface{}{}
	for _, x := range a {
		b = append(b, x)
	}
	io.PrintLn(b...)
}

func (io *Io) PrintStringLn(a []string) {
	b := []interface{}{}
	for _, x := range a {
		b = append(b, x)
	}
	io.PrintLn(b...)
}

func Log(name string, value interface{}) {
	fmt.Fprintf(os.Stderr, "%s=%+v\n", name, value)
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	k := io.NextInt()

	arrA := make([]int, n+1)
	arrB := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arrA[i] = io.NextInt()
	}
	for i := 1; i <= n; i++ {
		arrB[i] = io.NextInt()
	}

	dp := make([]bool, n+1)
	ep := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = false
		ep[i] = false
	}

	dp[1] = true
	ep[1] = true

	for i := 2; i <= n; i++ {
		if dp[i-1] {
			if abs(arrA[i-1]-arrA[i]) <= k {
				dp[i] = true
			}
			if abs(arrA[i-1]-arrB[i]) <= k {
				ep[i] = true
			}
		}
		if ep[i-1] {
			if abs(arrB[i-1]-arrA[i]) <= k {
				dp[i] = true
			}
			if abs(arrB[i-1]-arrB[i]) <= k {
				ep[i] = true
			}
		}
	}

	//答え
	if dp[n] || ep[n] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

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
	m := io.NextInt()

	arrA := make([][]int, m)
	for i := 0; i < m; i++ {
		arrA[i] = make([]int, n)
		for j := 0; j < n; j++ {
			arrA[i][j] = io.NextInt()
		}
	}

	arrB := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arrB[i] += arrA[i][j] << j
		}
	}

	dp := make([]int, (1<<n)+5)
	for i := 0; i < (1<<n)+5; i++ {
		dp[i] = 1000
	}
	dp[0] = 0

	for i := 0; i < m; i++ {
		for j := 0; j < (1 << n); j++ {
			dp[j|arrB[i]] = intMin(dp[j|arrB[i]], dp[j]+1)
		}
	}
	if dp[(1<<n)-1] == 1000 {
		fmt.Println(-1)
	} else {
		fmt.Println(dp[(1<<n)-1])
	}
}

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
	arrS := make([]string, n)
	for i := 0; i < n; i++ {
		arrS[i] = io.NextLine()
	}
	ans := solve(n, arrS)
	fmt.Println(ans)
}

func solve(n int, s []string) string {
	const size = 6
	for k := 0; k < 2; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n-size+1; j++ {
				black := 0
				for k := 0; k < size; k++ {
					if s[i][j+k] == '#' {
						black++
					}
				}
				if black >= size-2 {
					return "Yes"
				}
			}
		}
		for i := 0; i < n-size+1; i++ {
			for j := 0; j < n-size+1; j++ {
				black := 0
				for k := 0; k < size; k++ {
					if s[i+k][j+k] == '#' {
						black++
					}
				}
				if black >= size-2 {
					return "Yes"
				}
			}
		}

		t := make([][]rune, n)
		for i := range t {
			t[i] = make([]rune, n)
		}
		for i := 0; i < n; i++ {
			for j := n - 1; j >= 0; j-- {
				t[i][n-1-j] = rune(s[j][i])
			}
		}
		for i := 0; i < n; i++ {
			s[i] = string(t[i])
		}
	}
	return "No"
}

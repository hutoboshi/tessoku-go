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

var ans int
var a, s [9]int

func ck(x, y, z int) bool {
	if s[x] < s[y] {
		x, y = y, x
	}
	if s[x] < s[z] {
		x, z = z, x
	}
	return !(a[y] == a[z] && a[x]^a[y] != 0)
}

func dfs(step int) {
	if step > 9 {
		if ck(0, 1, 2) && ck(3, 4, 5) && ck(6, 7, 8) && ck(0, 3, 6) && ck(1, 4, 7) && ck(2, 5, 8) && ck(0, 4, 8) && ck(2, 4, 6) {
			ans++
		}
		return
	}
	for i := 0; i < 9; i++ {
		if s[i] == 0 {
			s[i] = step
			dfs(step + 1)
			s[i] = 0
		}
	}
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	a[0] = io.NextInt()
	a[1] = io.NextInt()
	a[2] = io.NextInt()
	a[3] = io.NextInt()
	a[4] = io.NextInt()
	a[5] = io.NextInt()
	a[6] = io.NextInt()
	a[7] = io.NextInt()
	a[8] = io.NextInt()
	dfs(1)
	fmt.Printf("%.10f", float64(ans)/362880.0)
}

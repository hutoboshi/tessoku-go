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

type Query struct {
	a, b, c, d int
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	h := io.NextInt()
	w := io.NextInt()
	arrX := make([][]int, h)
	for i := 0; i < h; i++ {
		arrX[i] = make([]int, w)
		for j := 0; j < w; j++ {
			arrX[i][j] = io.NextInt()
		}
	}

	q := io.NextInt()
	query := make([]Query, q)
	for i := 0; i < q; i++ {
		query[i] = Query{a: io.NextInt(), b: io.NextInt(), c: io.NextInt(), d: io.NextInt()}
	}

	arrZ := make([][]int, h+1)
	for i := 0; i < h+1; i++ {
		arrZ[i] = make([]int, w+1)
	}

	for i := 1; i < h+1; i++ {
		for j := 1; j < w+1; j++ {
			arrZ[i][j] = arrZ[i][j-1] + arrX[i-1][j-1]
		}
	}

	for j := 1; j < w+1; j++ {
		for i := 1; i < h+1; i++ {
			arrZ[i][j] += arrZ[i-1][j]
		}
	}

	for _, v := range query {
		fmt.Println(arrZ[v.c][v.d] + arrZ[v.a-1][v.b-1] - arrZ[v.a-1][v.d] - arrZ[v.c][v.b-1])
	}
}

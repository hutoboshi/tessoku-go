package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type menu struct {
	x int
	y int
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	x := io.NextInt()
	y := io.NextInt()
	// arrM:=make([]menu,n)
	// for i:=0;i<n;i++{
	// 	arrM[i].x=io.NextInt()
	// }
	// for i:=0;i<n;i++{
	// 	arrM[i].y=io.NextInt()
	// }
	arrA := make([]int, n)
	arrB := make([]int, n)
	for i := 0; i < n; i++ {
		arrA[i] = io.NextInt()
	}
	for i := 0; i < n; i++ {
		arrB[i] = io.NextInt()
	}

	sort.Ints(arrA)
	sort.Ints(arrB)

	sumA := 0
	a := 0
	for i := len(arrA) - 1; i >= 0; i-- {
		sumA += arrA[i]
		a++
		if sumA > x {
			break
		}
	}

	sumB := 0
	b := 0
	for i := len(arrB) - 1; i >= 0; i-- {
		sumB += arrB[i]
		b++
		if sumB > y {
			break
		}
	}

	fmt.Println(intMin(a, b))

}

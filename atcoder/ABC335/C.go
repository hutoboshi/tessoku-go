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

type Pos struct {
	x, y int
}

type Query struct {
	q, vI int
	vS    string
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	q := io.NextInt()
	arrC := make([]Pos, n)
	for i := n - 1; i >= 0; i-- {
		arrC[i] = Pos{x: n - i, y: 0}
	}

	arrQ := make([]Query, q)
	for i := 0; i < q; i++ {
		tmpQ := io.NextInt()
		if tmpQ == 2 {
			tmpI := io.NextInt()
			arrQ[i] = Query{q: tmpQ, vI: tmpI, vS: ""}
		} else if tmpQ == 1 {
			tmpS := io.Next()
			arrQ[i] = Query{q: tmpQ, vI: 0, vS: tmpS}
		}
	}

	for _, v := range arrQ {
		if v.q == 2 {
			fmt.Println(arrC[len(arrC)-v.vI].x, arrC[len(arrC)-v.vI].y)
		} else if v.q == 1 {
			if v.vS == "R" {
				arrC = append(arrC, Pos{x: arrC[len(arrC)-1].x + 1, y: arrC[len(arrC)-1].y})
			} else if v.vS == "L" {
				arrC = append(arrC, Pos{x: arrC[len(arrC)-1].x - 1, y: arrC[len(arrC)-1].y})
			} else if v.vS == "U" {
				arrC = append(arrC, Pos{x: arrC[len(arrC)-1].x, y: arrC[len(arrC)-1].y + 1})
			} else if v.vS == "D" {
				arrC = append(arrC, Pos{x: arrC[len(arrC)-1].x, y: arrC[len(arrC)-1].y - 1})
			}
		}
	}
	// fmt.Println(arrC)
}

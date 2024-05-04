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

type tuple struct {
	id, x, y int
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()

	arrX := make([]int, n)
	arrY := make([]int, n)
	for i := 0; i < n; i++ {
		arrX[i] = io.NextInt()
		arrY[i] = io.NextInt()
	}

	s := io.NextLine()

	arrP := make([]tuple, n)
	for i := 0; i < n; i++ {
		arrP[i] = tuple{i, arrX[i], arrY[i]}
	}
	sort.Slice(arrP, func(i, j int) bool {
		if arrP[i].y < arrP[j].y {
			return true
		} else if arrP[i].y == arrP[j].y {
			if arrP[i].x < arrP[j].x {
				return true
			}
		}
		return false
	})

	for i := 0; i < n; i++ {
		pre := arrP[i]
		if s[pre.id] != 'R' {
			continue
		}

		for j := i + 1; j < n; j++ {
			cur := arrP[j]
			if pre.y != cur.y {
				i = j
				break
			}
			if s[cur.id] == 'L' {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")
}

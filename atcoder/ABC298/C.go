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

type query struct {
	qi, i, j int
}

type inbox struct {
	boxI, n int
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	q := io.NextInt()
	arrQ := make([]query, q)
	for i := 0; i < q; i++ {
		i1 := io.NextInt()
		if i1 == 1 {
			arrQ[i] = query{i1, io.NextInt(), io.NextInt()}
			continue
		}
		if i1 == 2 {
			arrQ[i] = query{i1, io.NextInt(), 0}
			continue
		}
		if i1 == 3 {
			arrQ[i] = query{i1, io.NextInt(), 0}
			continue
		}
	}

	boxes := make(map[int][]int, n+1)

	inboxes := make(map[inbox]bool)

	for _, v := range arrQ {
		if v.qi == 1 {
			if _, ok := boxes[v.j]; !ok {
				boxes[v.j] = []int{}
			}
			boxes[v.j] = append(boxes[v.j], v.i)
			inboxes[inbox{v.j, v.i}] = true
			continue
		}
		if v.qi == 2 {
			sort.Ints(boxes[v.i])
			for _, idx := range boxes[v.i] {
				fmt.Print(idx, " ")
			}
			fmt.Println()
			continue
		}
		if v.qi == 3 {
			ans := []int{}
			for k, _ := range boxes {
				if _, ok := inboxes[inbox{k, v.i}]; ok {
					ans = append(ans, k)
				}
			}
			sort.Ints(ans)
			for _, v2 := range ans {
				fmt.Print(v2, " ")
			}
			fmt.Println()
			continue
		}
	}
}

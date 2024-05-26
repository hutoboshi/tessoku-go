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

// 処理がTLEする　高速化すればTLEは免れそう
func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	box := make([][]int, 200010)
	card := make([][]int, 200010)

	io.NextInt()
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

	for i := 0; i < q; i++ {
		que := arrQ[i]
		if que.qi == 1 {
			card[que.i] = append(card[que.i], que.j)
			box[que.j] = append(box[que.j], que.i)
			continue
		}
		if que.qi == 2 {
			sort.Ints(box[que.i])
			for _, v := range box[que.i] {
				fmt.Print(v, " ")
			}
			fmt.Println()
			continue
		}
		if que.qi == 3 {
			sort.Ints(card[que.i])
			box2 := unique(card[que.i])
			for _, v := range box2 {
				fmt.Print(v, " ")
			}
			fmt.Println()
			continue
		}
	}
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

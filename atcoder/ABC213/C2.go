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

// ２回目
type Card struct {
	a, b int
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	io.NextInt()
	io.NextInt()
	n := io.NextInt()

	cards := make([]Card, n)
	for i := 0; i < n; i++ {
		cards[i] = Card{a: io.NextInt(), b: io.NextInt()}
	}

	var points [][]int
	var heightList, widthList []int

	heightMap := make(map[int]int)
	widthMap := make(map[int]int)

	for _, v := range cards {
		if _, ok := heightMap[v.a]; !ok {
			heightMap[v.a] = 1
			heightList = append(heightList, v.a)
		}
		if _, ok := widthMap[v.b]; !ok {
			widthMap[v.b] = 1
			widthList = append(widthList, v.b)
		}
		points = append(points, []int{v.a, v.b})
	}

	sort.Ints(heightList)
	sort.Ints(widthList)

	for i := 0; i < len(heightList); i++ {
		heightMap[heightList[i]] = i
	}
	for i := 0; i < len(widthList); i++ {
		widthMap[widthList[i]] = i
	}

	for _, v := range points {
		fmt.Println(heightMap[v[0]]+1, widthMap[v[1]]+1)
	}
}

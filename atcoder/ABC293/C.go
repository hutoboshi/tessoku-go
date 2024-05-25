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

var m [][]int
var h, w int
var ans int

func dfs(i int, j int, visited map[int]bool) {
	//hwに来た時にpathを判定
	if i == h-1 && j == w-1 {
		ans++
		return
	}
	//i+1で再帰
	if i+1 < h {
		if _, ok := visited[m[i+1][j]]; ok {
			return
		}
		copyVisitedH := make(map[int]bool)
		for k, v := range visited {
			copyVisitedH[k] = v
		}
		copyVisitedH[m[i+1][j]] = true
		dfs(i+1, j, copyVisitedH)
	}
	//j+1で再帰
	if j+1 < w {
		if _, ok := visited[m[i][j+1]]; ok {
			return
		}
		copyVisitedW := make(map[int]bool)
		for k, v := range visited {
			copyVisitedW[k] = v
		}
		copyVisitedW[m[i][j+1]] = true
		dfs(i, j+1, copyVisitedW)
	}
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	h = io.NextInt()
	w = io.NextInt()

	m = make([][]int, h)
	for i := 0; i < h; i++ {
		m[i] = make([]int, w)
		for j := 0; j < w; j++ {
			m[i][j] = io.NextInt()
		}
	}
	mapX := make(map[int]bool)
	mapX[m[0][0]] = true
	dfs(0, 0, mapX)

	fmt.Println(ans)

}

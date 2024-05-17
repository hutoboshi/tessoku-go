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

// グラフ
type Graph struct {
	n     int
	edges [][]Edge
}

type Edge struct {
	from, to, w int
}

func NewGraph(n int) *Graph {
	g := &Graph{}
	g.n = n
	g.edges = make([][]Edge, g.n)
	return g
}

func (g *Graph) AddEdge(a, b, c int) {
	g.edges[a] = append(g.edges[a], Edge{a, b, c})
}

// DFSで指定したのーどに辿り着けるかどうか、ノードまでの経路を返す
func dfs(g *Graph, from int, to int) (bool, []int) {
	visited := make([]bool, g.n)
	path := make([]int, 0)
	var f func(int, int) bool
	f = func(pos int, to int) bool {
		visited[pos] = true
		path = append(path, pos)
		if pos == to {
			return true
		}
		for _, v := range g.edges[pos] {
			if visited[v.to] == false {
				ret := f(v.to, to)
				if ret == true {
					return true
				}
				path = path[:len(path)-1]
			}
		}
		return false
	}
	ret := f(from, to)
	if !ret {
		path = []int{}
	}
	return ret, path
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	x := io.NextInt()
	y := io.NextInt()

	g := NewGraph(n + 1)
	for i := 0; i < n-1; i++ {
		a := io.NextInt()
		b := io.NextInt()
		g.AddEdge(a, b, 1)
		g.AddEdge(b, a, 1)
	}
	_, ret := dfs(g, x, y)
	for _, v := range ret {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println()
}

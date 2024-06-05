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

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	h := io.NextInt()
	w := io.NextInt()
	arrS := make([]string, h)
	for i := 0; i < h; i++ {
		arrS[i] = io.NextLine()
	}

	uf := NewUnionFind(h * w)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if arrS[i][j] != '#' {
				continue
			}
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					ni := i + di
					nj := j + dj
					if ni < 0 || ni >= h || nj < 0 || nj >= w {
						continue
					}
					if arrS[ni][nj] != '#' {
						continue
					}
					if i == ni && j == nj {
						continue
					}
					uf.merge(i*w+j, ni*w+nj)
				}
			}
		}
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if arrS[i][j] == '#' {
				if uf.leader(i*w+j) == i*w+j {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}

type UnionFind struct {
	parentSize []int //親ノードとサイズの配列
}

// コンストラクタみたいな役割
func NewUnionFind(n int) *UnionFind {
	parentSize := make([]int, n)
	for i := range parentSize {
		parentSize[i] = -1
	}
	return &UnionFind{parentSize}
}

// aとbを結合
func (uf *UnionFind) merge(a, b int) {
	x, y := uf.leader(a), uf.leader(b)
	if x == y {
		return
	}
	if abs(uf.parentSize[x]) < abs(uf.parentSize[y]) {
		x, y = y, x
	}
	uf.parentSize[x] += uf.parentSize[y]
	uf.parentSize[y] = x
}

// aとbが同じグループに属しているかを判定
func (uf *UnionFind) sama(a, b int) bool {
	return uf.leader(a) == uf.leader(b)
}

// aの親ノードを取得
func (uf *UnionFind) leader(a int) int {
	if uf.parentSize[a] < 0 {
		return a
	}
	uf.parentSize[a] = uf.leader(uf.parentSize[a])
	return uf.parentSize[a]
}

// aが属するグループのサイズを取得
func (uf *UnionFind) size(a int) int {
	return abs(uf.parentSize[uf.leader(a)])
}

// unionfindのグループを取得
func (uf *UnionFind) groups() [][]int {
	result := make([][]int, len(uf.parentSize))
	for i := range uf.parentSize {
		result[uf.leader(i)] = append(result[uf.leader(i)], i)
	}
	groups := make([][]int, 0)
	for _, r := range result {
		if len(r) > 0 {
			groups = append(groups, r)
		}
	}
	return groups
}

// 整数の絶対値
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

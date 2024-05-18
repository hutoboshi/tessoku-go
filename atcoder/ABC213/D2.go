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

var arrG [][]int
var ans []int

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	arrA := make([]int, n-1)
	arrB := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		arrA[i] = io.NextInt()
		arrB[i] = io.NextInt()
	}

	arrG = make([][]int, n+1)
	for i := range arrG {
		arrG[i] = make([]int, 0)
	}

	//辺の情報arrA,arrBから隣接リストを作成
	for i := 0; i < n-1; i++ {
		arrG[arrA[i]] = append(arrG[arrA[i]], arrB[i])
		arrG[arrB[i]] = append(arrG[arrB[i]], arrA[i])
	}

	//各頂点の隣接をソートする（小さい順に回るため）
	for i := range arrG {
		sort.Ints(arrG[i])
	}

	//深さ優先探索を実行して結果を出力する
	dfs(1, -1)

	for _, val := range ans {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

func dfs(crr, pre int) {
	//現在の頂点を結果に追加
	ans = append(ans, crr)

	//隣接する各頂点について再帰的に探索する
	for _, nxt := range arrG[crr] {
		if nxt != pre {
			dfs(nxt, crr)
			ans = append(ans, crr)
		}
	}
}

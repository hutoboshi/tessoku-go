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

var (
	N, K int
	R    []int
)

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	N = io.NextInt()
	K = io.NextInt()
	R = make([]int, N)
	for i := 0; i < N; i++ {
		R[i] = io.NextInt()
	}

	var results [][]int
	generateSequence([]int{}, 0, 0, &results)

	// 辞書順に並べるためにソート
	sort.Slice(results, func(i, j int) bool {
		for k := 0; k < N; k++ {
			if results[i][k] != results[j][k] {
				return results[i][k] < results[j][k]
			}
		}
		return false
	})

	// 結果を出力
	for _, seq := range results {
		for i, num := range seq {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(num)
		}
		fmt.Println()
	}
}

func generateSequence(current []int, index, currentSum int, results *[][]int) {
	if len(current) == N {
		if currentSum%K == 0 {
			*results = append(*results, append([]int(nil), current...))
		}
		return
	}

	for i := 1; i <= R[index]; i++ {
		// 枝刈り：すでに総和がKの倍数でない場合、その先を探索しない
		if (currentSum+i)%K != 0 && index == N-1 {
			continue
		}
		generateSequence(append(current, i), index+1, currentSum+i, results)
	}
}

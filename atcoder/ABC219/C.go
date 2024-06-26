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

/*
小英字
小文字の列が決まっていない時に独自の列挙列を作って判定する
*/
func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	x := io.Next()
	n := io.NextInt()
	arrS := make([]string, n)
	for i := 0; i < n; i++ {
		arrS[i] = io.Next()
	}

	//sort.Sliceを使って文字列のスライス（arrS）をソートする比較関数
	sort.Slice(arrS, func(i, j int) bool {
		//文字列を先頭から比較
		for u := 0; u < intMin(len(arrS[i]), len(arrS[j])); u++ {
			//文字列が異なる場合、文字列xでのそれぞれの文字のインデックスを比較して、順序を決定する
			if arrS[i][u] != arrS[j][u] {
				return strings.Index(x, string(arrS[i][u])) < strings.Index(x, string(arrS[j][u]))
			}
		}
		//先頭部分が一致する場合、短い方の文字列が先に来るように長さを比較する
		return len(arrS[i]) < len(arrS[j])
	})

	for _, v := range arrS {
		fmt.Println(v)
	}
}

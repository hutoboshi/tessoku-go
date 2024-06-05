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
	// 入出力用のIo構造体を作成
	io := NewIo()
	defer io.Flush() // main関数の終了時にFlushを呼び出す

	n := io.NextInt() // 文字列の数を読み込む
	t := io.Next()    // 比較対象の文字列を読み込む
	arrS := make([]string, n)
	for i := 0; i < n; i++ {
		arrS[i] = io.NextLine() // 各文字列を読み込む
	}

	ans := make([]int, 0)
	for i, v := range arrS {
		if v == t {
			// 完全一致の場合
			ans = append(ans, i+1)
			continue
		}
		if len(v) == len(t) {
			// 長さが同じ場合の1文字違いのチェック
			difference := 0
			for j := 0; j < len(v); j++ {
				if v[j] != t[j] {
					difference++
				}
			}
			if difference == 1 {
				ans = append(ans, i+1)
			}
		} else if len(v)-len(t) == 1 {
			// vのほうが1文字長い場合のチェック
			for j := 0; j < len(v); j++ {
				// 修正点：v[:j]+v[j+1:] == t のチェック
				if v[:j]+v[j+1:] == t {
					ans = append(ans, i+1)
					break
				}
			}
		} else if len(t)-len(v) == 1 {
			// tのほうが1文字長い場合のチェック
			for j := 0; j < len(t); j++ {
				// 修正点：t[:j]+t[j+1:] == v のチェック
				if t[:j]+t[j+1:] == v {
					ans = append(ans, i+1)
					break
				}
			}
		}
	}

	// 修正点：結果の出力
	io.PrintIntLn(ans)
}

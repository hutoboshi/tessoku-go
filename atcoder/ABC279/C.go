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

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	h := io.NextInt()
	w := io.NextInt()
	RS := make([][]byte, w)
	RT := make([][]byte, w)
	for i := 0; i < h; i++ {
		s := io.NextLine()
		for j := 0; j < w; j++ {
			RS[j] = append(RS[j], s[j])
		}
	}
	for i := 0; i < h; i++ {
		s := io.NextLine()
		for j := 0; j < w; j++ {
			RT[j] = append(RT[j], s[j])
		}
	}

	S := make([]string, w)
	T := make([]string, w)

	for i := 0; i < w; i++ {
		S[i] = string(RS[i])
	}
	for i := 0; i < w; i++ {
		T[i] = string(RT[i])
	}

	//SとTをそれぞれ辞書順にソート
	sort.Slice(S, func(i, j int) bool {
		return S[i] < S[j]
	})
	sort.Slice(T, func(i, j int) bool {
		return T[i] < T[j]
	})

	//SとTが一致するか確認
	for i := 0; i < w; i++ {
		if S[i] != T[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

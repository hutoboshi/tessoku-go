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

// func main() {
// 	//決まり文句
// 	io := NewIo()
// 	defer io.Flush()

// 	N := io.NextInt()
// 	S := make([]string, N)
// 	M := -1
// 	for i := 0; i < N; i++ {
// 		S[i] = io.NextLine()
// 		M = intMax(M, len(S[i]))
// 	}

// 	// T を作成
// 	T := make([][]rune, M)
// 	for i := 0; i < M; i++ {
// 		T[i] = make([]rune, N)
// 		for j := 0; j < N; j++ {
// 			if i < len(S[j]) {
// 				T[i][j] = rune(S[j][N-j-1])
// 			} else {
// 				T[i][j] = '*'
// 			}
// 		}
// 	}

// 	// 末尾の '*' を除去して出力
// 	for i := 0; i < M; i++ {
// 		output := strings.TrimRight(string(T[i]), "*")
// 		fmt.Println(output)
// 	}

// }

func main() {
	// 決まり文句
	io := NewIo()
	defer io.Flush()

	N := io.NextInt()
	S := make([]string, N)
	M := -1
	for i := 0; i < N; i++ {
		S[i] = io.NextLine()
		M = intMax(M, len(S[i]))
	}

	// T を作成
	T := make([][]rune, M)
	for i := 0; i < M; i++ {
		T[i] = make([]rune, N)
		for j := 0; j < N; j++ {
			if i < len(S[j]) {
				T[i][j] = rune(S[j][i])
			} else {
				T[i][j] = '*'
			}
		}
	}

	// 末尾の '*' を除去して出力
	for i := 0; i < M; i++ {
		output := strings.TrimRight(string(T[i]), "*")
		fmt.Println(output)
	}
}

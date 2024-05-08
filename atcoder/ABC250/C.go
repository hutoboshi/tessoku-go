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

	n := io.NextInt()
	q := io.NextInt()

	arrX := make([]int, q)
	for i := 0; i < q; i++ {
		arrX[i] = io.NextInt()
	}

	//並び替えする配列
	arrN := make([]int, n)
	for i := 0; i < n; i++ {
		arrN[i] = i + 1
	}

	//並び替え順を記憶するmap
	mapN := make(map[int]int)
	for i := 0; i < n; i++ {
		mapN[i+1] = i
	}

	//並び替え
	for _, x := range arrX {
		pos := mapN[x]
		if pos == n-1 {
			tmp := arrN[n-2]
			//配列を入れ替える
			arrN[n-1], arrN[n-2] = arrN[n-2], arrN[n-1]
			//mapを入れ替える
			mapN[x] = n - 2
			mapN[tmp] = n - 1
		} else {
			tmp := arrN[pos+1]
			//配列を入れ替える
			arrN[pos], arrN[pos+1] = arrN[pos+1], arrN[pos]
			//mapを入れ替える
			mapN[x] = pos + 1
			mapN[tmp] = pos
		}
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Print(arrN[i])
		} else {
			fmt.Print(" ")
			fmt.Print(arrN[i])
		}
	}
	fmt.Println()
}

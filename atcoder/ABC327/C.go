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

	arrS := make([][]int, 9)
	for i := 0; i < 9; i++ {
		arrS[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			arrS[i][j] = io.NextInt()
		}
	}

	for i := 0; i < 9; i++ {
		han := make([]bool, 10)
		for j := 0; j < 9; j++ {
			han[arrS[i][j]] = true
		}
		for j := 1; j <= 9; j++ {
			if !han[j] {
				fmt.Println("No")
				return
			}
		}
	}

	for i := 0; i < 9; i++ {
		han := make([]bool, 10)
		for j := 0; j < 9; j++ {
			han[arrS[j][i]] = true
		}
		for j := 1; j <= 9; j++ {
			if !han[j] {
				fmt.Println("No")
				return
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			han := make([]bool, 10)
			for h := 0; h < 3; h++ {
				for w := 0; w < 3; w++ {
					han[arrS[i*3+h][j*3+w]] = true
				}
			}
			for a := 1; a <= 9; a++ {
				if !han[a] {
					fmt.Println("No")
					return
				}
			}
		}
	}
	fmt.Println("Yes")
}

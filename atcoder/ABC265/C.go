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

	arrG := make([]string, h)
	for i := 0; i < h; i++ {
		arrG[i] = io.NextLine()
	}

	//訪問済みか確認する
	wentHW := make([][]bool, h)
	for i := 0; i < h; i++ {
		wentHW[i] = make([]bool, w)
		for j := 0; j < w; j++ {
			wentHW[i][j] = false
		}
	}
	wentHW[0][0] = true

	nowH, nowW := 0, 0
	for {
		pos := arrG[nowH][nowW]
		if pos == 'U' {
			if nowH == 0 {
				break
			} else {
				nowH--
			}
		}
		if pos == 'D' {
			if nowH == h-1 {
				break
			} else {
				nowH++
			}
		}
		if pos == 'L' {
			if nowW == 0 {
				break
			} else {
				nowW--
			}
		}
		if pos == 'R' {
			if nowW == w-1 {
				break
			} else {
				nowW++
			}
		}
		if wentHW[nowH][nowW] {
			fmt.Println(-1)
			return
		} else {
			wentHW[nowH][nowW] = true
		}
	}

	fmt.Println(nowH+1, nowW+1)
}

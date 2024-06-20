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

	n := io.NextInt()
	m := io.NextInt()
	arrA := make([]int, n)
	for i := 0; i < n; i++ {
		arrA[i] = io.NextInt()
	}
	arrB := make([]int, m)
	for i := 0; i < m; i++ {
		arrB[i] = io.NextInt()
	}

	sort.Ints(arrA)
	sort.Ints(arrB)

	index := sort.Search(len(arrA), func(i int) bool {
		return arrA[i] >= arrB[0]
	})

	// index2 := 0
	// for i := 0; i < n; i++ {
	// 	if arrA[i] > arrB[0] {
	// 		index2 = i
	// 		break
	// 	}
	// }

	// fmt.Println(index2 + m + 1)
	// fmt.Println(index)
	// fmt.Println(index2)

	// if index+m > len(arrA)-1 {
	// 	fmt.Println("-1")
	// 	return
	// }

	ans := 0
	arrAindex := index
	arrBindex := 0

	for arrBindex < len(arrB) && arrAindex < len(arrA) {
		if arrB[arrBindex] <= arrA[arrAindex] {
			// fmt.Println(arrA[arrAindex])
			ans += arrA[arrAindex]
			arrAindex++
			arrBindex++
		} else {
			arrAindex++
		}
	}

	// fmt.Println(arrAindex)
	// fmt.Println(arrBindex)
	if arrAindex == len(arrA) && arrBindex != len(arrB) {
		fmt.Println(-1)
		return
	}

	fmt.Println(ans)

	// fmt.Println(arrA)
	// fmt.Println(arrB)

}
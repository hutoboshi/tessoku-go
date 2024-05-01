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
	m := io.NextInt()

	arrA := make([]int, m)
	arrB := make([]int, m)
	arrC := make([]int, m)
	arrD := make([]int, m)
	for i := 0; i < m; i++ {
		arrA[i] = io.NextInt()
		arrB[i] = io.NextInt()
	}
	for i := 0; i < m; i++ {
		arrC[i] = io.NextInt()
		arrD[i] = io.NextInt()
	}

	arrP := make([][]bool, n)
	arrQ := make([][]bool, n)
	for i := 0; i < n; i++ {
		arrP[i], arrQ[i] = make([]bool, n), make([]bool, n)
	}
	for i := 0; i < m; i++ {
		arrP[arrA[i]-1][arrB[i]-1] = true
		arrP[arrB[i]-1][arrA[i]-1] = true
	}
	for i := 0; i < m; i++ {
		arrQ[arrC[i]-1][arrD[i]-1] = true
		arrQ[arrD[i]-1][arrC[i]-1] = true
	}

	rng := make([]int, n)
	for i := 0; i < n; i++ {
		rng[i] = i
	}
	for _, v := range permutation(rng) {
		flag := true
		for i := 0; flag && i < n; i++ {
			for j := 0; flag && j < n; j++ {
				if arrP[i][j] != arrQ[v[i]][v[j]] {
					flag = false
				}
			}
		}
		if flag {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")

}

func permutation(nums []int) [][]int {
	n := 1
	for i := 2; i <= len(nums); i++ {
		n *= i
	}
	result := make([][]int, 0, n)
	result = append(result, append([]int{}, nums...))
	lenn := len(nums)
	p := make([]int, lenn+1)
	for i := 0; i < lenn+1; i++ {
		p[i] = i
	}
	for i := 1; i < lenn; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		nums[i], nums[j] = nums[j], nums[i]
		result = append(result, append([]int{}, nums...))
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
	return result
}

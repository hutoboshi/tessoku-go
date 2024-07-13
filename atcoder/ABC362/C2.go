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

func findCombinations(arrMin []int, arrMax []int, index int, currentSum int, combination []int, result *[][]int) {
	if index == len(arrMin) {
		if currentSum == 0 {
			comboCopy := make([]int, len(combination))
			copy(comboCopy, combination)
			*result = append(*result, comboCopy)
		}
		return
	}

	for val := arrMin[index]; val <= arrMax[index]; val++ {
		if currentSum+val <= 0 {
			combination = append(combination, val)
			findCombinations(arrMin, arrMax, index+1, currentSum+val, combination, result)
			combination = combination[:len(combination)-1]
		}
	}
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	var n int
	fmt.Scan(&n)
	arrMin := make([]int, n)
	arrMax := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arrMin[i])
		fmt.Scan(&arrMax[i])
	}

	var result [][]int
	var combination []int

	findCombinations(arrMin, arrMax, 0, 0, combination, &result)

	if len(result) > 0 {
		fmt.Println("Possible combinations that sum to 0:")
		for _, combo := range result {
			fmt.Println(combo)
		}
	} else {
		fmt.Println("No combination found that sums to 0.")
	}
}

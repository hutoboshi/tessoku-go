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

func generateCombinations(arrMin []int, arrMax []int, index int, combination []int, result *[][]int) {
	if index == len(arrMin) {
		// Make a copy of combination and append to result
		comboCopy := make([]int, len(combination))
		copy(comboCopy, combination)
		*result = append(*result, comboCopy)
		return
	}

	for val := arrMin[index]; val <= arrMax[index]; val++ {
		combination = append(combination, val)
		generateCombinations(arrMin, arrMax, index+1, combination, result)
		combination = combination[:len(combination)-1]
	}
}

// Helper function to find combinations that sum to target
func findZeroSumCombinations(combinations [][]int) [][]int {
	var zeroSumCombinations [][]int
	for _, combo := range combinations {
		sum := 0
		for _, num := range combo {
			sum += num
		}
		if sum == 0 {
			zeroSumCombinations = append(zeroSumCombinations, combo)
		}
	}
	return zeroSumCombinations
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	arrMin := make([]int, n)
	arrMax := make([]int, n)
	for i := 0; i < n; i++ {
		arrMin[i] = io.NextInt()
		arrMax[i] = io.NextInt()
	}

}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

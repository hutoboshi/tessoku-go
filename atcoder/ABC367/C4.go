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
	io := NewIo()
	defer io.Flush()

	N := io.NextInt()
	K := io.NextInt()
	R := make([]int, N)
	for i := 0; i < N; i++ {
		R[i] = io.NextInt()
	}

	var results [][]int

	// DPテーブル
	dp := make([][][]bool, N+1)
	for i := range dp {
		dp[i] = make([][]bool, K)
		for j := range dp[i] {
			dp[i][j] = make([]bool, K)
		}
	}
	dp[0][0][0] = true

	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			for m := 0; m < K; m++ {
				if dp[i][j][m] {
					for x := 1; x <= R[i]; x++ {
						nj := (j + x) % K
						dp[i+1][nj][(m+x)%K] = true
					}
				}
			}
		}
	}

	// 可能な組み合わせを生成
	var sequence []int
	backtrack(N, K, 0, R, dp, &results, &sequence)

	// 結果を辞書順にソート
	sort.Slice(results, func(i, j int) bool {
		for k := 0; k < len(results[i]); k++ {
			if results[i][k] != results[j][k] {
				return results[i][k] < results[j][k]
			}
		}
		return false
	})

	// 結果を出力
	for _, seq := range results {
		for i, num := range seq {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(num)
		}
		fmt.Println()
	}
}

func backtrack(N, K, mod int, R []int, dp [][][]bool, results *[][]int, sequence *[]int) {
	if len(*sequence) == N {
		if mod == 0 {
			*results = append(*results, append([]int(nil), *sequence...))
		}
		return
	}

	for x := 1; x <= R[len(*sequence)]; x++ {
		nextMod := (mod + x) % K
		if dp[len(*sequence)+1][nextMod][nextMod] {
			*sequence = append(*sequence, x)
			backtrack(N, K, nextMod, R, dp, results, sequence)
			*sequence = (*sequence)[:len(*sequence)-1]
		}
	}
}

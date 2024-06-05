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

type Respondent struct {
	s     string
	score int
	noans []int
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	m := io.NextInt()
	arrA := make([]int, m)
	for i := 0; i < m; i++ {
		arrA[i] = io.NextInt()
	}

	respondents := make([]Respondent, n)
	for i := 0; i < n; i++ {
		respondents[i].s = io.NextLine()
		respondents[i].score += i + 1
		respondents[i].noans = make([]int, 0)
		for j := 0; j < m; j++ {
			if respondents[i].s[j] == 'o' {
				respondents[i].score += arrA[j]
			} else {
				respondents[i].noans = append(respondents[i].noans, arrA[j])
			}
		}
	}

	//回答者の最高得点を求める
	maxScore := 0
	for _, v := range respondents {
		maxScore = intMax(maxScore, v.score)
	}

	//プレイヤー全員の得点を上回れるか判定
	for _, v := range respondents {
		sort.Slice(v.noans, func(i, j int) bool {
			return v.noans[i] > v.noans[j]
		})
		addAns := 0
		if maxScore == v.score {
			fmt.Println(addAns)
			continue
		}
		for _, v2 := range v.noans {
			addAns++
			v.score += v2
			if v.score > maxScore {
				break
			}
		}
		fmt.Println(addAns)
	}
}

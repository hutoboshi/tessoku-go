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
	arrA := make([]int, n)
	arrB := make([]int, n)
	for i := 0; i < n; i++ {
		arrA[i] = io.NextInt()
		arrB[i] = io.NextInt()
	}

	arrS := make([]float64, n)                    //長さnのスライスを作成
	arrS[0] = float64(arrA[0]) / float64(arrB[0]) //arrS[0]をarrA[0]/arrB[0]で初期化
	for i := 1; i < n; i++ {
		arrS[i] = arrS[i-1] + float64(arrA[i])/float64(arrB[i]) //arrS[i]をarrS[i-1]+arrA[i]/arrB[i]で計算
	}

	// ansを計算
	var ans float64
	for i := 0; ; i++ {
		if arrS[i] < (arrS[n-1] / 2.0) { // S[i] が S[N-1]/2.0 より小さい場合
			ans += float64(arrA[i]) // ans に A[i] を加算
			continue
		} else { // S[i] が S[N-1]/2.0 以上の場合
			ans += float64(arrA[i])                               // ans に A[i] を加算
			ans -= (arrS[i] - (arrS[n-1])/2.0) * float64(arrB[i]) // ans から超過した部分を減算
			break
		}
	}
	fmt.Println(ans)
}

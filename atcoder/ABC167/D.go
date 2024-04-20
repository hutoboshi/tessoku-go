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

	n := io.NextInt() //町の数
	k := io.NextInt() //移動回数

	arrA := make([]int, n+1)
	arrA[0] = 0
	for i := 1; i < n+1; i++ {
		arrA[i] = io.NextInt() //各町への移動先を記録
	}

	visited := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		visited[i] = -1
	}

	nowTown := 1 //現在の街の初期値
	moveCnt := 0 //移動回数の初期値
	cycle := 0   //周期の初期値

	//k解移動するループ
	for i := 1; i < 1000001; i++ {
		moveCnt += 1
		nowTown = arrA[nowTown] //移動先の街を更新

		//移動回数がkに達したら、現在の街を出力して終了
		if moveCnt == k {
			fmt.Println(nowTown)
			return
		}

		//街の訪問履歴を更新
		if visited[nowTown] == -1 {
			visited[nowTown] = moveCnt
		} else {
			//周期を検出した場合、周期の長さを求める
			cycle = moveCnt - visited[nowTown]
			break
		}
	}

	//k回移動する際の補正を行う
	k -= moveCnt
	k %= cycle

	//k回移動した後の街を求めて出力
	for i := 1; i <= k; i++ {
		nowTown = arrA[nowTown]
	}

	fmt.Println(nowTown)
}

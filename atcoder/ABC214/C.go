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

	//arrSとarrTの配列を初期化して入力を読み込む
	arrS := make([]int, n)
	arrT := make([]int, n)
	for i := 0; i < n; i++ {
		arrS[i] = io.NextInt()
	}

	//最小値を記録する変数とその時のインデックスを初期化する
	var si int
	var mint int = 1e9 + 1

	for i := 0; i < n; i++ {
		arrT[i] = io.NextInt()
		if mint > arrT[i] {
			mint = arrT[i]
			si = i
		}
	}

	//arrSとarrTの値を読み取りながら最小値を更新
	ans := make([]int, n)            //長さnの整数型スライスansを作成し、0で初期化する
	ans[si] = arrT[si]               //最初の最小値のインデックスsiにarrT[si]の値を格納する
	for i := si + 1; i < si+n; i++ { //si+1からsi＋n-1までの範囲でループする
		j := i % n                         //iをnで割った余りをjに格納する
		br := ans[(i-1)%n] + arrS[(i-1)%n] //現在の要素を更新するための値brを計算する。直前の要素のansとarrSの値を加算する
		if br < arrT[j] {                  //もしbrがarrT[j]より小さい場合
			ans[j] = br //ans[j]をbrに更新する
		} else { //そうではない場合
			ans[j] = arrT[j] //ans[j]をarrT[j]に更新
		}
	}

	//答えを出力
	for _, v := range ans {
		fmt.Println(v)
	}

}

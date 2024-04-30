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

// 文字列をソートする関数
func sortStr(w string) string {
	s := strings.Split(w, "")  //文字列を1文字ずつに分解する
	sort.Strings(s)            //文字列をソートする
	return strings.Join(s, "") //ソートされた文字列を結合してreturn
}

// 文字列を反転させる関数
func reverseStr(s string) string {
	runes := []rune(s) //文字列をルーンのスライスに変換
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] //文字を反転させる
	}
	return string(runes) //反転された文字列をreturn
}

/*
１つの整数を2つの整数に分けてa*bの積の最大値を求める
*/
func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextLine() //文字列で受け取る

	var tmp, ans int

	for i := 0; i < (1 << len(n)); i++ {
		var a, b string
		for j := 0; j < len(n); j++ {
			if i>>j&1 == 1 {
				a += n[j : j+1] //2進数でのビット演算を使って、文字列を2つの部分文字列に分割
			} else {
				b += n[j : j+1]
			}
		}

		if a == "" || b == "" {
			continue //空の文字列は無視
		}

		a = reverseStr(sortStr(a)) //部分文字列をソートして反転
		b = reverseStr(sortStr(b))

		if a[0] == '0' || b[0] == '0' {
			continue //先頭が0の場合は無視
		}

		tmp_a, _ := strconv.Atoi(a)
		tmp_b, _ := strconv.Atoi(b)
		tmp = tmp_a * tmp_b //部分文字列を変数に変換して積を計算

		if tmp >= ans {
			ans = tmp //最大の席を更新する
		}
	}

	fmt.Println(ans) //最大の積を出力する

}

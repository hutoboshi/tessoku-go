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

var l int
var s, t string
var res []string

// find 文字列sから部分文字列を生成する再帰関数
// flはビットフラ部で使用されている文字を示す
func find(fl int) {
	//flが0の場合、全ての文字が使用されたことを意味し、tを結果に追加する
	if fl == 0 {
		res = append(res, t)
		return
	}
	//flのビットごとに処理を行い、使用されている文字をtに追加して再帰呼び出しをする
	for i := 0; i < l; i++ {
		if fl&(1<<i) != 0 {
			t += string(s[i])
			find(fl ^ (1 << i))
			t = t[:len(t)-1] //tから最後の文字を削除して元に戻す
		}
	}
}

// unique 文字列スライスから重複を削除する関数
func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	s = io.Next()
	k := io.NextInt()

	l = len(s)            //文字列sの長さを取得する
	find((1 << l) - 1)    //全ての文字を使った状態のフラグを設定して、find関数を呼び出す
	sort.Strings(res)     //結果を辞書順にソートする
	res = unique(res)     //結果から重複を削除する
	fmt.Println(res[k-1]) //k番目の結果を出力する
}

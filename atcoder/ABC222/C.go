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

// プレイヤーの情報
type Player struct {
	win int //勝利数
	num int //プレイヤー数
}

// ２つの手を比較し、照射を返す関数
func judge(a, b byte) int {
	if a == b {
		return 0 //引き分け
	} else if a == 'G' && b == 'C' || a == 'C' && b == 'P' || a == 'P' && b == 'G' {
		return 1 //aが勝つ
	} else {
		return -1 //bが勝つ
	}
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	m := io.NextInt()

	//プレイヤーの情報を保持する配列を初期化
	result := make([]Player, 2*n)
	for i := 0; i < 2*n; i++ {
		result[i] = Player{win: 0, num: i}
	}

	//プレイヤーの手を保持する配列を初期化する
	hands := make([]string, 2*n)
	for i := 0; i < 2*n; i++ {
		hands[i] = io.NextLine()
	}

	//各試合ごとに勝敗を計算し、プレイヤーの勝利数を更新する
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p1 := result[2*j].num
			p2 := result[2*j+1].num
			ret := judge(hands[p1][i], hands[p2][i])
			if ret == 1 {
				result[2*j].win++
			} else if ret == -1 {
				result[2*j+1].win++
			}
		}

		//勝利数でプレイヤーをソートする。同点の場合はプレイヤーの番号で比較
		sort.Slice(result, func(i, j int) bool {
			if result[i].win == result[j].win {
				return result[i].num < result[j].num
			} else {
				return result[i].win > result[j].win
			}
		})
	}

	//結果を出力
	for i := 0; i < 2*n; i++ {
		fmt.Println(result[i].num + 1)
	}
}

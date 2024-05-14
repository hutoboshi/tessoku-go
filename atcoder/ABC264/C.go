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

	h1 := io.NextInt()
	w1 := io.NextInt()
	arrA := make([][]int, h1)
	for i := 0; i < h1; i++ {
		arrA[i] = make([]int, w1)
		for j := 0; j < w1; j++ {
			arrA[i][j] = io.NextInt()
		}
	}
	h2 := io.NextInt()
	w2 := io.NextInt()
	arrB := make([][]int, h1)
	for i := 0; i < h2; i++ {
		arrB[i] = make([]int, w1)
		for j := 0; j < w2; j++ {
			arrB[i][j] = io.NextInt()
		}
	}

	r := false
	//AからBを作れるか
	for i := 0; i < (1 << uint64(h1)); i++ { //行から組み合わせを全探索
		for j := 0; j < (1 << uint64(w1)); j++ { //列の組み合わせを全探索
			tmp_v := []int{} //行の組み合わせを保持する一時的なスライス
			tmp_y := []int{} //列の組み合わせを保持する一時的なスライス
			for m := 0; m <= h1; m++ {
				if (uint64(i) >> m & 1) == 1 { //1ビットが立っているかどうかチェック
					tmp_v = append(tmp_v, m) //立っていれば行の組み合わせに追加
				}
			}
			for m := 0; m <= w1; m++ {
				if (uint64(j) >> m & 1) == 1 { //1ビットが立っているかどうかチェック
					tmp_y = append(tmp_y, m) //立っていれば列の組み合わせに追加
				}
			}

			//行の組み合わせと列の組み合わせが合致するかどうかを判定
			if len(tmp_v) != h2 || len(tmp_y) != w2 {
				continue
			}

			match := true //マッチフラグ
			//AとBの要素が対応するかどうかチェック
			for k := 0; k < h2; k++ {
				for l := 0; l < w2; l++ {
					if arrA[tmp_v[k]][tmp_y[l]] != arrB[k][l] { //対応する要素が異なる場合
						match = false //マッチしないと判定
						break         //内側のループを抜ける
					}
				}
			}

			//全ての要素が対応していればマッチフラグを垂れて
			if match {
				r = true
				break
			}
		}
		if r {
			break
		}
	}

	//結果を出力
	if r {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

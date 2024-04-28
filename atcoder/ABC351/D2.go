/*
深さ優先探索
*/

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

var e [][]int  //隣接リスト
var used []int //ノードの訪問状態を管理する配列
var cnt int    //連結成分の大きさをカウントする変数

func dfs(s, v int) {
	if used[v] == s {
		return
	}
	used[v] = s //ノードを訪問済みとしてマーク
	cnt++       //連結成分の大きさをインクリメント

	//隣接するノードを再起的に探索
	for _, u := range e[v] {
		dfs(s, u)
	}
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	dx := []int{0, 0, 1, -1} //x方向の変位
	dy := []int{1, -1, 0, 0} //x方向の変位
	h := io.NextInt()
	w := io.NextInt()

	arrS := make([]string, h) //行列を表す文字列のスライスを初期化
	for i := 0; i < h; i++ {
		arrS[i] = io.NextLine()
	}

	e = make([][]int, h*w)  //隣接リストを初期化
	used = make([]int, h*w) //訪問状態を管理する配列を初期化

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if arrS[i][j] == '#' {
				continue //壁の場合はスキップ
			}
			can := true //連結成分の探索が可能かどうかを示すフラグ
			for k := 0; k < 4; k++ {
				nx, ny := i+dx[k], j+dy[k]                  //隣接するノードを座標を計算
				if nx >= 0 && nx < h && ny >= 0 && ny < w { //行列内かどうか確認
					if arrS[nx][ny] == '#' { //隣接するノードが壁の場合
						can = false //連結成分の探索が不可能
					} else {
						e[i*w+j] = append(e[i*w+j], nx*w+ny) //隣接リストに追加
					}
				}
			}
			if !can {
				e[i*w+j] = nil //連結成分の探索が不可能な場合は隣接リストをクリア
			}
		}
	}

	ans := 0
	for i := 0; i < h*w; i++ {
		used[i] = -1 //訪問状態を初期化
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if arrS[i][j] == '.' && used[i*w+j] < 0 {
				cnt = 0           //連結成分の大きさを初期化
				dfs(i*w+j, i*w+j) //DFSで連結成分の大きさを求める
				if cnt > ans {    //現在の連結成分の大きさが最大かどうか判定
					ans = cnt //最大の場合は答えを更新
				}
			}
		}
	}
	fmt.Println(ans) //答えを出力
}

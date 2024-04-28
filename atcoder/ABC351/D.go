package main

import (
	"bufio"
	"container/list"
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

// deque
type Deque struct {
	list *list.List
}

func NewDeque() *Deque {
	return &Deque{
		list: list.New(),
	}
}

func (d *Deque) PushFront(item interface{}) {
	d.list.PushFront(item)
}

func (d *Deque) PushBack(item interface{}) {
	d.list.PushBack(item)
}

func (d *Deque) PopFront() interface{} {
	if d.list.Len() == 0 {
		return nil
	}
	front := d.list.Front()
	d.list.Remove(front)
	return front.Value
}

func (d *Deque) PopBack() interface{} {
	if d.list.Len() == 0 {
		return nil
	}
	back := d.list.Back()
	d.list.Remove(back)
	return back.Value
}

func (d *Deque) Len() int {
	return d.list.Len()
}

// 迷路探索
func explore(start_gyou, start_retu, h, w int, maze []string) int {
	//各セルへの最短距離を保持する配列を初期化
	maze_count := make([][]int, h)
	for i := 0; i < h; i++ {
		maze_count[i] = make([]int, w)
		for j := 0; j < w; j++ {
			maze_count[i][j] = -1
		}
	}
	maze_count[start_gyou][start_retu] = 0

	//幅優先探索用のキューを初期化し、スタート地点を追加
	gyou_deque := NewDeque()
	retu_deque := NewDeque()
	gyou_deque.PushBack(start_gyou)
	retu_deque.PushBack(start_retu)

	//上下左右への移動方向を表す配列
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	//幅優先探索で各セルへの最短距離を計算
	for gyou_deque.Len() > 0 {
		now_gyou, now_retu := gyou_deque.PopFront().(int), retu_deque.PopFront().(int)
		now_count := maze_count[now_gyou][now_retu]

		//上下左右に移動して次のセルを探索
		for _, dir := range dirs {
			next_gyou, next_retu := now_gyou+dir[0], now_retu+dir[1]
			if next_gyou >= 0 && next_gyou < h && next_retu >= 0 && next_retu < w && maze[next_gyou][next_retu] == '.' && maze_count[next_gyou][next_retu] == -1 {
				//移動可能な迷路内で見探索かる通行可能な場合、距離を更新してキューに追加
				if now_gyou > 0 && maze[now_gyou-1][now_retu] != '#' {
					continue
				}
				if now_gyou > 0 && maze[now_gyou+1][now_retu] != '#' {
					continue
				}
				if now_retu > 0 && maze[now_gyou][now_retu-1] != '#' {
					continue
				}
				if now_retu > 0 && maze[now_gyou][now_retu+1] != '#' {
					continue
				}
				maze_count[next_gyou][next_retu] = now_count + 1
				gyou_deque.PushBack(next_gyou)
				retu_deque.PushBack(next_retu)
			}
		}
	}

	//全てのセルの最短距離を調べて最大距離を求める
	// max_count := 0
	// for _, row := range maze_count {
	// 	for _, val := range row {
	// 		if val != -1 {
	// 			max_count++
	// 		}
	// 	}
	// }

	// max_count:=0
	// for _,val :=maze_count{
	// 	for _,val2:=val{
	// 		if val2!=-1{
	// 			max_count++
	// 		}
	// 	}
	// }

	max_count := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if maze_count[i][j] != -1 {
				max_count++
			}
		}
	}

	return max_count
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	h := io.NextInt()
	w := io.NextInt()
	arrS := make([]string, h)
	for i := 0; i < h; i++ {
		arrS[i] = io.NextLine()
	}

	ans := 0
	for gyou := 0; gyou < h; gyou++ {
		for retu := 0; retu < w; retu++ {
			if arrS[gyou][retu] == '.' {
				ans = intMax(ans, explore(gyou, retu, h, w, arrS))
			}
		}
	}
	fmt.Println(ans)
}

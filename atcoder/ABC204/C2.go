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

func bfs(start int, n int, connect [][]int) int {
	//訪問済みリスト作成
	visited := make([]bool, n+1)
	//訪問済みにする
	visited[start] = true
	cnt := 1
	deque := NewDeque()
	deque.PushBack(start)

	for deque.Len() > 0 {
		now_city := deque.PopFront()
		for _, to_city := range connect[now_city.(int)] {
			if !visited[to_city] {
				visited[to_city] = true
				cnt++
				deque.PushBack(to_city)
			}
		}
	}
	return cnt
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	m := io.NextInt()

	arrA := make([]int, m)
	arrB := make([]int, m)
	for i := 0; i < m; i++ {
		arrA[i] = io.NextInt()
		arrB[i] = io.NextInt()
	}

	connect := make([][]int, n+1)

	for i := 0; i < m; i++ {
		connect[arrA[i]] = append(connect[arrA[i]], arrB[i])
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += bfs(i, n, connect)
	}

	fmt.Println(ans)
}

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
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

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()

	n := io.NextInt()
	arrA := make([]int, n)
	for i := 0; i < n; i++ {
		arrA[i] = io.NextInt()
	}

	deque := NewDeque()
	for i := 0; i < n; i++ {
		deque.PushBack(int(math.Pow(2, float64(arrA[i]))))
		flg := true
		for flg {
			if deque.Len() <= 1 {
				flg = false
				continue
			}
			i1 := deque.PopFront().(int)
			i2 := deque.PopFront().(int)
			if i1 != i2 {
				deque.PushFront(i2)
				deque.PushFront(i1)
				flg = false
				continue
			} else {
				i3 := i1 + i2
				deque.PushBack(i3)
			}
		}
	}

	fmt.Println(deque.Len())

}

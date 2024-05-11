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

}

type MapMultiSet struct {
	m map[int]*MultiSet
}

func NewMapMultiSet() *MapMultiSet {
	ms := &MapMultiSet{}
	ms.m = make(map[int]*MultiSet)
	return ms
}
func (ms *MapMultiSet) Get(key int) *MultiSet {
	if ms.m[key] == nil {
		ms.m[key] = NewMultiSet()
	}
	return ms.m[key]
}

type MultiSet struct {
	tree *AvlTree
	n    int
}

func NewMultiSet() *MultiSet {
	s := &MultiSet{}
	s.tree = NewAvlTree()
	return s
}
func (s *MultiSet) Put(key, num int) {
	s.n++
	s.tree.Put(key, num)
}
func (s *MultiSet) Add(key, num int) {
	val, ok := s.tree.Get(key)
	if ok == false {
		s.n++
	}
	s.tree.Put(key, val+num)
}

func (s *MultiSet) Remove(key, num int) {
	val, _ := s.tree.Get(key)
	if val-num <= 0 {
		s.tree.Remove(key)
		s.n--
		return
	}
	s.tree.Put(key, val-num)
}
func (s *MultiSet) RemoveKey(key int) {
	_, ok := s.tree.Get(key)
	if ok == true {
		s.n--
		s.tree.Remove(key)
	}
}
func (s *MultiSet) Clear() {
	s.tree.Clear()
	s.n = 0
}
func (s *MultiSet) Len() int {
	return s.n
}
func (s *MultiSet) Get(key int) (int, bool) {
	val, ok := s.tree.Get(key)
	if ok == false {
		return 0, false
	}
	return val, ok
}
func (s *MultiSet) Ceil(key int) (int, int, bool) {
	ret1, ret2 := s.tree.Ceiling(key)
	if ret2 == false {
		return 0, 0, false
	}
	return ret1.Key, ret1.Value, ret2
}
func (s *MultiSet) Floor(key int) (int, int, bool) {
	ret1, ret2 := s.tree.Floor(key)
	if ret2 == false {
		return 0, 0, false
	}
	return ret1.Key, ret1.Value, ret2
}
func (s *MultiSet) GetNthSmall(x, nth int) (int, bool) {
	cnt := nth
	for i := 0; i < nth; i++ {
		r1, r2, r3 := s.Floor(x)
		if r3 == false {
			return 0, false
		}
		cnt -= r2
		if cnt <= 0 {
			return r1, true
		}
		x = r1 - 1
	}
	return 0, false
}
func (s *MultiSet) GetNthLarge(x, nth int) (int, bool) {
	cnt := nth
	for i := 0; i < nth; i++ {
		r1, r2, r3 := s.Ceil(x)
		if r3 == false {
			return 0, false
		}
		cnt -= r2
		if cnt <= 0 {
			return r1, true
		}
		x = r1 + 1
	}
	return 0, false
}
func (s *MultiSet) MaxKey() (int, bool) {
	if s.tree.Right() == nil {
		return 0, false
	}
	return s.tree.Right().Key, true
}
func (s *MultiSet) MinKey() (int, bool) {
	if s.tree.Left() == nil {
		return 0, false
	}
	return s.tree.Left().Key, true
}
func (s *MultiSet) GetKeys() []int {
	ret := make([]int, 0, s.Len())
	it := s.tree.Iterator()
	for it.Next() {
		key := it.Key()
		val := it.Value()
		for i := 0; i < val; i++ {
			ret = append(ret, key)
		}
	}
	return ret
}
func (s *MultiSet) GetUniqKeys() []int {
	return s.tree.Keys()
}
func (s *MultiSet) GetRangeKeys(from, to int) []int {
	ret := make([]int, 0, s.Len())
	it := s.tree.Iterator()
	for it.Next() {
		key := it.Key()
		if from <= key && key <= to {
			val := it.Value()
			for i := 0; i < val; i++ {
				ret = append(ret, key)
			}
		}
	}
	return ret
}
func (s *MultiSet) GetRangeUniqKeys(from, to int) []int {
	ret := make([]int, 0, s.Len())
	it := s.tree.Iterator()
	for it.Next() {
		key := it.Key()
		if from <= key && key <= to {
			ret = append(ret, key)
		}
	}
	return ret
}

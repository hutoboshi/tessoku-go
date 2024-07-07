package main

import (
	"bufio"
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

type RectangularPrism struct {
	x1, y1, z1, x2, y2, z2 float64
}

func NewRectangularPrism(a, b, c, d, e, f float64) RectangularPrism {
	return RectangularPrism{
		x1: math.Min(a, d),
		y1: math.Min(b, e),
		z1: math.Min(c, f),
		x2: math.Max(a, d),
		y2: math.Max(b, e),
		z2: math.Max(c, f),
	}
}

func IntersectionVolume(r1, r2 RectangularPrism) float64 {
	x_overlap := math.Max(0, math.Min(r1.x2, r2.x2)-math.Max(r1.x1, r2.x1))
	y_overlap := math.Max(0, math.Min(r1.y2, r2.y2)-math.Max(r1.y1, r2.y1))
	z_overlap := math.Max(0, math.Min(r1.z2, r2.z2)-math.Max(r1.z1, r2.z1))

	return x_overlap * y_overlap * z_overlap
}

func HasPositiveIntersection(r1, r2 RectangularPrism) bool {
	return IntersectionVolume(r1, r2) > 0
}

func main() {
	//決まり文句
	io := NewIo()
	defer io.Flush()
	a := io.NextInt()
	b := io.NextInt()
	c := io.NextInt()
	d := io.NextInt()
	e := io.NextInt()
	f := io.NextInt()
	g := io.NextInt()
	h := io.NextInt()
	i := io.NextInt()
	j := io.NextInt()
	k := io.NextInt()
	l := io.NextInt()

	r1 := NewRectangularPrism(float64(a), float64(b), float64(c), float64(d), float64(e), float64(f))
	r2 := NewRectangularPrism(float64(g), float64(h), float64(i), float64(j), float64(k), float64(l))

	if HasPositiveIntersection(r1, r2) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

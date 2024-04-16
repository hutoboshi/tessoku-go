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

	h := io.NextInt()
	w := io.NextInt()
	k := io.NextInt()

	arrC := make([][]byte, h)
	for i := 0; i < h; i++ {
		tmp := io.Next()
		arrC[i] = []byte(tmp)
	}

	ans := 0

	//行と列の選択パターンを全探索
	for gyouRed := 0; gyouRed < (1 << h); gyouRed++ {
		for retuRed := 0; retuRed < (1 << w); retuRed++ {
			block := 0

			//各ますをチェック
			for gyou := 0; gyou < h; gyou++ {
				for retu := 0; retu < w; retu++ {

					//赤い行または列の中に含まれておらず、かつ黒いマスの場合
					if gyouRed>>gyou&1 == 0 && retuRed>>retu&1 == 0 {
						if arrC[gyou][retu] == '#' {
							block++
						}
					}
				}
			}

			//黒いマスの数がKと一致する場合、解の候補としてカウント
			if block == k {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

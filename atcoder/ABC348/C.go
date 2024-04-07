package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Split(bufio.ScanWords)

	var n int
	n = nextInt()

	color := make(map[int]int)
	for i := 0; i < n; i++ {
		var a, c int
		a = nextInt()
		c = nextInt()
		if val, ok := color[c]; ok {
			color[c] = min(val, a)
		} else {
			color[c] = a
		}
	}

	var minTaste int
	for _, taste := range color {
		minTaste = max(minTaste, taste)
	}

	fmt.Println(minTaste)
}

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

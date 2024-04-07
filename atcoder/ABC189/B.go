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
	n := inputInt()
	x := inputInt()

	arrV := make([]int, n)
	arrP := make([]int, n)
	for i := 0; i < n; i++ {
		arrV[i] = inputInt()
		arrP[i] = inputInt()
	}

	alcohol := 0
	for i := 0; i < n; i++ {
		alcohol += arrV[i] * arrP[i]
		if 100*x < alcohol {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(-1)
}

func inputInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func inputString() string {
	scanner.Scan()
	text := scanner.Text()
	return text
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

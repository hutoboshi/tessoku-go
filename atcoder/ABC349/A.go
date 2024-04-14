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
	arrA := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		arrA[i] = inputInt()
	}

	var point int
	for i := 0; i < n-1; i++ {
		point += arrA[i]
	}

	if point == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(point * -1)
	}
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

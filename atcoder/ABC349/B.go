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
	s := inputString()

	arrS := make(map[rune]int)
	for _, v := range s {
		arrS[v]++
	}

	arrAns := make(map[int]int)
	for _, v := range arrS {
		arrAns[v]++
	}

	for _, v := range arrAns {
		if v != 0 && v != 2 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
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

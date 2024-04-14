package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Split(bufio.ScanWords)

	var s, t string
	fmt.Scan(&s, &t)

	j := 0
	for _, c := range s {
		if strings.ToUpper(string(c)) == string(t[j]) {
			j++
			if j == 3 {
				break
			}
		}
	}

	if j == 2 && t[j] == 'X' {
		j++
	}

	if j == 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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

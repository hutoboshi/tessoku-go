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

	lowerT := strings.ToLower(t)
	t1 := rune(lowerT[0])
	t2 := rune(lowerT[1])
	t3 := rune(lowerT[2])

	S1 := false
	S2 := false
	S3 := false

	for _, v := range s {
		if v == t1 && !S2 && !S3 && !S1 {
			S1 = true
		} else if v == t2 && S1 && !S3 && !S2 {
			S2 = true
		} else if v == t3 && S1 && S2 && !S3 {
			S3 = true
		}
	}

	if S1 && S2 && t3 == 'x' {
		fmt.Println("Yes")
	} else if S1 && S2 && S3 {
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

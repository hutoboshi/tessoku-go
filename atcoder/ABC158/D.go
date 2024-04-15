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
	var a string
	a = inputString()
	var q int
	q = inputInt()

	inv := false

	var s_deque string
	s_deque = a

	var input1, input2, input3 string
	for i := 0; i < q; i++ {
		input1 = inputString()
		if input1 == "1" {
			inv = !inv
		} else {
			input2 = inputString()
			input3 = inputString()
			if inv {
				if input2 == "1" {
					s_deque = s_deque + input3
				} else {
					s_deque = input3 + s_deque
				}
			} else {
				if input2 == "1" {
					s_deque = input3 + s_deque
				} else {
					s_deque = s_deque + input3
				}
			}
		}
	}

	ans := s_deque

	if inv {
		for i := len(ans) - 1; i >= 0; i-- {
			fmt.Print(string(ans[i]))
		}
		fmt.Println()
	} else {
		fmt.Println(ans)
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

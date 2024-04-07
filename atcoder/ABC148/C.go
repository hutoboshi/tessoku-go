/*
最小公倍数
*/
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
	a := inputInt()
	b := inputInt()
	fmt.Println(lcm(a, b))
}

func lcm(x, y int) int {
	return (x * y) / gcd(x, y)
}

// 最大公約数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
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

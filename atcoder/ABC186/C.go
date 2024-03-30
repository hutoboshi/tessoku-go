package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 10進数に7が入っているか調べる
func judgeTen(x int) bool {
	xStr := strconv.Itoa(x)
	return strings.Contains(xStr, "7")
}

// ８進数に7が入っているか調べる
func judgeEight(x int) bool {
	s := ""
	for x > 0 {
		s = strconv.Itoa(x%8) + s
		x /= 8
	}
	return strings.Contains(s, "7")
}

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0

	for i := 1; i <= n; i++ {
		if !judgeTen(i) && !judgeEight(i) {
			ans++
		}
	}
	fmt.Println(ans)
}

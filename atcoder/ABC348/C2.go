package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	colors := make(map[int]int)
	for i := 0; i < N; i++ {
		var A, C int
		fmt.Scan(&A, &C)
		if val, ok := colors[C]; ok {
			colors[C] = min(val, A)
		} else {
			colors[C] = A
		}
	}

	var minTaste int
	for _, taste := range colors {
		minTaste = max(minTaste, taste)
	}

	fmt.Println(minTaste)
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

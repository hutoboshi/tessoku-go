package main

import (
	"fmt"
	"math"
)

func main() {
	var H, N int
	fmt.Scan(&H, &N)

	dp := make([]int, H+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 0; i < N; i++ {
		var A, B int
		fmt.Scan(&A, &B)
		for h := 0; h <= H; h++ {
			if h+A <= H {
				dp[h+A] = min(dp[h+A], dp[h]+B)
			} else {
				dp[H] = min(dp[H], dp[h]+B)
			}
		}
	}

	fmt.Println(dp[H])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

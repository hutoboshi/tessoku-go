package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var N, W int
	fmt.Scan(&N, &W)

	items := make([][]int, N+1)
	items[0] = []int{0, 0}
	for i := 1; i <= N; i++ {
		var w, v int
		fmt.Scan(&w, &v)
		items[i] = []int{w, v}
	}

	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for gyou_i := 1; gyou_i <= N; gyou_i++ {
		for retu_w := 0; retu_w <= W; retu_w++ {
			w, v := items[gyou_i][0], items[gyou_i][1]
			if retu_w-w < 0 {
				dp[gyou_i][retu_w] = dp[gyou_i-1][retu_w]
			} else {
				red := dp[gyou_i-1][retu_w]
				blue := dp[gyou_i-1][retu_w-w] + v
				dp[gyou_i][retu_w] = max(red, blue)
			}
		}
	}

	fmt.Println(dp[N][W])
}

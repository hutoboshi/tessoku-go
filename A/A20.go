package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	n, m := len(s), len(t)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	dp[0][0] = 0
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i >= 1 && j >= 1 && s[i-1] == t[j-1] {
				dp[i][j] = max(dp[i-1][j], max(dp[i][j-1], dp[i-1][j-1]+1))
			} else if i >= 1 && j >= 1 {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			} else if i >= 1 {
				dp[i][j] = dp[i-1][j]
			} else if j >= 1 {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	fmt.Println(dp[n][m])
}

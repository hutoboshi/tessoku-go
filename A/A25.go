package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	arrC := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&arrC[i])
	}

	dp := make([][]int64, h)
	for i := range dp {
		dp[i] = make([]int64, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
				if i >= 1 && arrC[i-1][j] == '.' {
					dp[i][j] += dp[i-1][j]
				}
				if j >= 1 && arrC[i][j-1] == '.' {
					dp[i][j] += dp[i][j-1]
				}
			}
		}
	}

	fmt.Println(dp[h-1][w-1])
}

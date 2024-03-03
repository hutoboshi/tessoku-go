package main

import "fmt"

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	//入力
	var n, w int
	fmt.Scan(&n, &w)

	//入力
	arrW := make([]int, n+1)
	arrV := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arrW[i], &arrV[i])
	}

	//動的計画法2次元配列作成
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, w+1)
	}

	//動的計画法でありえない最低のマイナスを2次元配列入れる
	for i := 0; i <= n; i++ {
		for j := 0; j <= w; j++ {
			dp[i][j] = -1000000000000
		}
	}

	//最初は0
	dp[0][0] = 0
	//動的計画法
	for i := 1; i <= n; i++ {
		for j := 0; j <= w; j++ {
			if j < arrW[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-arrW[i]]+int64(arrV[i]))
			}
		}
	}

	//一番大きい値を調べる
	var ans int64
	for i := 0; i <= w; i++ {
		ans = max(ans, dp[n][i])
	}

	//答えを出力
	fmt.Println(ans)
}

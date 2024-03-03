package main

import "fmt"

func main() {
	//入力
	var n, s int
	fmt.Scan(&n, &s)
	//入力
	arrA := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arrA[i])
	}

	//動的計画法2次元配列作成
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, s+1)
	}

	//0には0しか入らない
	dp[0][0] = true
	for i := 1; i <= s; i++ {
		dp[0][i] = false
	}

	//動的計画法
	for i := 1; i <= n; i++ {
		for j := 0; j <= s; j++ {
			if j < arrA[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-arrA[i]]
			}
		}
	}

	if dp[n][s] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

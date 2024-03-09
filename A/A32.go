package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	//勝者を計算する
	dp := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		if i >= a && !dp[i-a] {
			dp[i] = true //勝ちの状態
		} else if i >= b && !dp[i-b] {
			dp[i] = true //勝ちの状態
		} else {
			dp[i] = false //負けの状態
		}
	}

	//出力
	if dp[n] {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}

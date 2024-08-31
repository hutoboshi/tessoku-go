package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 入力の取得
	var N int
	fmt.Scan(&N)

	// 演奏情報を格納するスライス
	keys := make([]int, N)
	hands := make([]rune, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&keys[i], &hands[i])
	}

	// 初期のDPテーブルの値を無限大に設定
	const inf = math.MaxInt32
	dp := make([][]int, 101)
	for i := range dp {
		dp[i] = make([]int, 101)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	// 両手を好きな場所に置けるので、初期疲労度は0
	dp[0][0] = 0

	// DPの更新
	for i := 0; i < N; i++ {
		nextKey := keys[i]
		if hands[i] == 'L' {
			// 左手で鍵盤を押す場合
			for l := 0; l <= 100; l++ {
				for r := 0; r <= 100; r++ {
					if dp[l][r] == inf {
						continue
					}
					// 左手を移動
					dp[nextKey][r] = min(dp[nextKey][r], dp[l][r]+int(math.Abs(float64(l-nextKey))))
				}
			}
		} else {
			// 右手で鍵盤を押す場合
			for l := 0; l <= 100; l++ {
				for r := 0; r <= 100; r++ {
					if dp[l][r] == inf {
						continue
					}
					// 右手を移動
					dp[l][nextKey] = min(dp[l][nextKey], dp[l][r]+int(math.Abs(float64(r-nextKey))))
				}
			}
		}
	}

	// 最小の疲労度を探す
	result := inf
	for l := 0; l <= 100; l++ {
		for r := 0; r <= 100; r++ {
			result = min(result, dp[l][r])
		}
	}

	fmt.Println(result)
}

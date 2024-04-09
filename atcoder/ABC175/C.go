package main

import (
	"fmt"
	"math"
)

func main() {
	var X, K, D int64
	fmt.Scan(&X, &K, &D)
	X = int64(math.Abs(float64(X))) // Xを絶対値に変換

	// 直線的に移動できる回数を計算
	straight := min(K, X/D)
	K -= straight
	X -= straight * D

	if K%2 == 0 {
		fmt.Println(X) // 偶数回の場合、現在地が答え
	} else {
		fmt.Println(D - X) // 奇数回の場合、逆方向に移動して現在地が答え
	}
}

// min関数を定義
func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

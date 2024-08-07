package main

import (
	"fmt"
)

// 交通費補助額の総和を計算する関数
func totalCost(A []int, x int) int64 {
	var total int64
	for _, cost := range A {
		if cost > x {
			total += int64(x)
		} else {
			total += int64(cost)
		}
	}
	return total
}

func main() {
	var N int
	var M int64
	fmt.Scan(&N, &M)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	// 交通費補助額の上限額の範囲を設定
	left, right := 0, int(1e9+1)

	for left < right {
		mid := (left + right + 1) / 2
		if totalCost(A, mid) <= M {
			left = mid
		} else {
			right = mid - 1
		}
	}

	if totalCost(A, int(1e9)) <= M {
		fmt.Println("infinite")
	} else {
		fmt.Println(left)
	}
}

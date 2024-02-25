package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	arrA := make([]int, n+1)
	arrB := make([]int, n+1)
	for i := 2; i <= n; i++ {
		fmt.Scan(&arrA[i])
	}
	for i := 3; i <= n; i++ {
		fmt.Scan(&arrB[i])
	}

	//動的計画法
	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = arrA[2]
	for i := 3; i <= n; i++ {
		dp[i] = min(dp[i-1]+arrA[i], dp[i-2]+arrB[i])
	}

	fmt.Println(dp[n])

}

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

	var ans []int
	place := n
	for {
		ans = append(ans, place)
		if place == 1 {
			break
		}
		if dp[place-1]+arrA[place] == dp[place] {
			place = place - 1
		} else {
			place = place - 2
		}
	}

	fmt.Println(len(ans))
	for i := len(ans) - 1; i >= 0; i-- {
		if i < len(ans)-1 {
			fmt.Print(" ")
		}
		fmt.Print(ans[i])
	}
	fmt.Println()

}

package main

import "fmt"

func fibonacciMemo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, exists := memo[n]; exists {
		return val
	}
	memo[n] = fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
	return memo[n]
}

func main() {
	var n int
	fmt.Scan(&n)
	memo := make(map[int]int)
	fmt.Printf("anser:%d\n", fibonacciMemo(n, memo))
}

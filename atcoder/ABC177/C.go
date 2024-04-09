package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	var mod int = 1000000007
	var sum int64 = 0
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
		sum += int64(A[i])
		sum %= int64(mod)
	}

	var ans int64 = 0

	// i について全探索する
	for i := 0; i < N; i++ {
		// A[i+1] ... A[N] の値を更新する
		sum -= int64(A[i])
		if sum < 0 {
			sum += int64(mod)
		}

		ans += int64(A[i]) * sum
		ans %= int64(mod)
	}

	fmt.Println(ans)
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arrA := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arrA[i])
	}

	cnt := make([]int64, 101)
	for i := 1; i <= 100; i++ {
		cnt[i] = 0
	}
	for i := 1; i <= n; i++ {
		cnt[arrA[i]]++
	}

	var ans int64 = 0
	for i := 1; i <= 100; i++ {
		ans += cnt[i] * (cnt[i] - 1) * (cnt[i] - 2) / 6
	}
	fmt.Println(ans)
}

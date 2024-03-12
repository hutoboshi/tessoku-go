package main

import "fmt"

func main() {
	var n, m int
	var b int64

	fmt.Scan(&n, &m, &b)

	arrA := make([]int64, n+1)
	arrC := make([]int64, m+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arrA[i])
	}
	for i := 1; i <= m; i++ {
		fmt.Scan(&arrC[i])
	}

	//答えの計算
	var ans int64 = 0
	for i := 1; i <= n; i++ {
		ans += arrA[i] * int64(m)
	}
	ans += b * int64(n) * int64(m)
	for j := 1; j <= m; j++ {
		ans += arrC[j] * int64(n)
	}

	//出力
	fmt.Println(ans)
}

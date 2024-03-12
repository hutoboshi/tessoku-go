package main

import "fmt"

func main() {
	var d, n int
	fmt.Scan(&d, &n)

	arrL := make([]int, n+1)
	arrR := make([]int, n+1)
	arrH := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arrL[i], &arrR[i], &arrH[i])
	}

	//配列の初期化（1日24時間）
	lim := make([]int, d+1)
	for i := 1; i <= d; i++ {
		lim[i] = 24
	}

	//上限値を求める
	for i := 1; i <= n; i++ {
		for j := arrL[i]; j <= arrR[i]; j++ {
			if lim[j] > arrH[i] {
				lim[j] = arrH[i]
			}
		}
	}

	//出力
	ans := 0
	for i := 1; i <= d; i++ {
		ans += lim[i]
	}
	fmt.Println(ans)
}

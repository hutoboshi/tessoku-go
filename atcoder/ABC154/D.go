package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i])
	}

	//pを変換したpExを計算する
	pEx := make([]float64, n)
	for i := 0; i < n; i++ {
		pEx[i] = float64(p[i]+1) / 2.0
	}

	//pExの累積話を計算する
	pExCum := make([]float64, n)
	pExCum[0] = pEx[0]
	for i := 1; i < n; i++ {
		pExCum[i] = pExCum[i-1] + pEx[i]
	}

	//初期値を設定する
	ans := -1e15

	//区間の最大値を求める
	for i := 0; i < n-k+1; i++ {
		var ansTmp float64
		if i == 0 {
			ansTmp = pExCum[i+k-1]
		} else {
			ansTmp = pExCum[i+k-1] - pExCum[i-1]
		}
		ans = max(ans, ansTmp)
	}

	fmt.Println(ans)
}

// 最大値を求めるユーティリティ関数
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

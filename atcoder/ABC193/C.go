/*
全探索2
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// 可能な数を格納するセット
	ableNum := make(map[int]bool)

	// aとbの範囲を探索してn以下の値をセットに追加
	for a := 2; a <= 100000+10; a++ {
		for b := 2; b < 100; b++ {
			// aのb乗を計算
			power := 1
			for i := 0; i < b; i++ {
				power *= a
				if power > n {
					break
				}
			}
			if power <= n {
				ableNum[power] = true
			} else {
				break
			}
		}
	}

	// 結果を出力
	fmt.Println(n - len(ableNum))
}

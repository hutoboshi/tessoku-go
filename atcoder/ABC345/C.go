// package main

// import "fmt"

// func swapChars(s string, i, j int) string {
// 	r := []rune(s)
// 	r[i], r[j] = r[j], r[i]
// 	return string(r)
// }

// func main() {
// 	var s string
// 	fmt.Scan(&s)

// 	cnt := 0
// 	mapTmp := make(map[string]bool)
// 	mapTmp[s] = true
// 	for i := 0; i <= len(s)-2; i++ {
// 		for j := 1; j <= len(s)-1; j++ {
// 			if s[i] == s[j] {
// 				continue
// 			} else {
// 				tmp := swapChars(s, i, j)
// 				if _, ok := mapTmp[tmp]; ok {
// 					continue
// 				} else {
// 					mapTmp[tmp] = true
// 					cnt++
// 				}
// 			}
// 		}
// 	}

//		if cnt == 0 {
//			fmt.Println(1)
//		} else {
//			fmt.Println(cnt)
//		}
//	}
package main

import (
	"fmt"
)

func main() {
	var n, ans int64  // n: 文字列の長さ, ans: 答えの変数
	var cnt [26]int64 // 各アルファベットの出現回数を保持する配列
	var same bool     // 同じ文字が複数回出現するかどうかを示すフラグ
	var s string      // 入力される文字列を保持する変数

	// 文字列を入力
	fmt.Scan(&s)
	n = int64(len(s))

	// 各アルファベットの出現回数をカウント
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}

	// 答えの初期値を設定
	ans = n * n

	// 同じ文字の組み合わせによって減るパターン数を計算
	for i := 0; i < 26; i++ {
		ans -= cnt[i] * cnt[i] // 同じアルファベット同士を入れ替えた場合の減少数を計算
		if cnt[i] > 1 {        // 同じ文字が2回以上出現する場合、sameフラグをtrueにする
			same = true
		}
	}

	ans /= 2 // 入れ替えた後の文字列のパターン数を計算

	if same { // 同じ文字が複数回出現する場合、結果に1を加える
		ans++
	}

	fmt.Println(ans) // 答えを出力
}

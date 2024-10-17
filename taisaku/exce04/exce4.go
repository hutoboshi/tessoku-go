// bit全探索
// ABC182 C
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Scan(&n)

	d := len(n)

	//答えの初期化
	ans := 100

	for num := 0; num < (1 << d); num++ {
		//桁を消して作れる数字を一時的に格納する変数
		n_tmp := ""
		//消した秒数をカウント
		ans_tmp := 0

		//各桁をチェックするループ
		for shift := 0; shift < d; shift++ {
			//1 &(numをshift個右シフト)が1ならばその桁を使う
			if (1 & (num >> shift)) == 1 {
				n_tmp += string(n[shift])
			} else {
				//1 & (numをshift個右シフト)が0ならば桁を消した個数をカウント
				ans_tmp++
			}
		}

		if n_tmp == "" {
			continue
		}

		numTmp, _ := strconv.Atoi(n_tmp)
		if numTmp%3 == 0 {
			if ans_tmp < ans {
				ans = ans_tmp
			}
		}
	}

	if ans == 100 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

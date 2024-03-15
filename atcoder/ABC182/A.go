/*
ビット全検索
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//入力を受け取る
	var n string
	fmt.Scan(&n)

	//nの桁数を取得
	d := len(n)

	//答えを初期化
	ans := 100

	// num=(000...0)~(111...1)の範囲でループ
	for num := 0; num < (1 << uint64(d)); num++ {
		//桁を消して作れる数字を一時的に格納する変数
		n_tmp := ""
		//消した行数をカウントする変数
		ans_tmp := 0

		//各桁をチェックするループ
		for shift := 0; shift < d; shift++ {
			//1 & (numをshift個右シフト)が1ならばその桁を使う
			if (1 & (num >> uint(shift))) == 1 {
				n_tmp += string(n[shift])
			} else {
				//1 & (numをshift個右シフト)が0ならば桁を消した個数をカウント
				ans_tmp++
			}
		}

		//n_tmpが空なら(numが000...0)の場合
		if n_tmp == "" {
			//次のnumへ
			continue
		}

		//n_tmpを整数に変換して3で割り切れるかチェック
		numTmp, _ := strconv.Atoi(n_tmp)
		if numTmp%3 == 0 {
			//ansよりans_tmpが小さければ更新
			if ans_tmp < ans {
				ans = ans_tmp
			}
		}
	}

	//答えが更新されていなければ-1を出力、更新されていれば答えを出力
	if ans == 100 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

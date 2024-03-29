package main

import (
	"fmt"
)

func main() {
	var w, b int
	fmt.Scan(&w, &b) // 白い石と黒い石の数を入力から受け取る

	// パターンtを定義
	t := "wbwbwwbwbwbw"

	// パターンtを利用して、すべての可能な開始位置から正方形を作り、条件を満たすかどうかを判定
	for i := 0; i < len(t); i++ {
		nw, nb := 0, 0 // 白い石と黒い石のカウントを初期化
		for j := 0; j < w+b; j++ {
			// パターンtをループしながら白い石と黒い石を数える
			if t[(i+j)%len(t)] == 'w' {
				nw++
			} else {
				nb++
			}
		}
		// 条件を満たす正方形が見つかれば"Yes"を出力して終了
		if w == nw && b == nb {
			fmt.Println("Yes")
			return
		}
	}
	// 条件を満たす正方形が見つからなければ"No"を出力して終了
	fmt.Println("No")
}

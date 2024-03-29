package main

import (
	"fmt"
)

func main() {
	var w, b int
	fmt.Scan(&w, &b)

	pt := "wbwbwwbwbwbwwbw"

	for {
		for _, v := range pt {
			if v == 'w' {
				w--
				if w == 0 && b == 0 {
					fmt.Println("Yes")
					return
				}
				if w == -1 {
					fmt.Println("No")
					return
				}
			}
			if v == 'b' {
				b--
				if w == 0 && b == 0 {
					fmt.Println("Yes")
					return
				}
				if b == -1 {
					fmt.Println("No")
					return
				}
			}
		}
	}
}

// package main

// import "fmt"

// func checkPatternExists(W, B int) bool {
// 	// 無限に繰り返す文字列S
// 	S := "wbwbwwbwbwbw"
// 	// W + B が S の周期で割り切れるかチェック
// 	return (W+B)%len(S) == 0
// }

// func main() {
// 	// WとBの数を指定
// 	var W, B int
// 	fmt.Scan(&W, &B)
// 	// 条件を満たす部分文字列が存在するかチェック
// 	exists := checkPatternExists(W, B)
// 	if exists {
// 		fmt.Println("条件を満たす部分文字列が存在します。")
// 	} else {
// 		fmt.Println("条件を満たす部分文字列は存在しません。")
// 	}
// }

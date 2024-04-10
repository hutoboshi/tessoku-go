package main

import (
	"fmt"
)

func main() {
	var N, Q int
	fmt.Scan(&N)
	var S string
	fmt.Scan(&S)
	fmt.Scan(&Q)

	// 左半分と右半分に分割
	S_left := S[:N]
	S_right := S[N:]

	flip := false // フリップ操作の状態を保持

	for i := 0; i < Q; i++ {
		var T, A, B int
		fmt.Scan(&T, &A, &B)

		if T == 1 {
			// フリップしている場合はインデックスを調整
			if flip {
				if A <= N {
					A += N
				} else {
					A -= N
				}
				if B <= N {
					B += N
				} else {
					B -= N
				}
			}
			// 文字列をスワップ
			S_left, S_right = swapChars(S_left, S_right, A-1, B-1)
		} else {
			// フリップ操作
			flip = !flip
		}
	}

	// フリップの状態に応じて出力を調整
	if flip {
		fmt.Println(S_right + S_left)
	} else {
		fmt.Println(S_left + S_right)
	}
}

// 文字列内の文字をスワップするユーティリティ関数
func swapChars(s1, s2 string, i, j int) (string, string) {
	r1 := []rune(s1)
	r2 := []rune(s2)
	if i >= 0 && i < len(r1) && j >= 0 && j < len(r2) {
		r1[i], r2[j] = r2[j], r1[i]
	}
	return string(r1), string(r2)
}

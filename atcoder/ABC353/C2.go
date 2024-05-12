package main

import "fmt"

func main() {
	// 入力値を受け取る
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	// 初期化
	result := 0
	prev := A[0]

	// f(A) を計算
	for i := 1; i < N; i++ {
		sum := (prev + A[i]) % 100000000
		result = (result + sum) % 100000000
		prev = sum
	}

	// 結果を出力
	fmt.Println(result)
}

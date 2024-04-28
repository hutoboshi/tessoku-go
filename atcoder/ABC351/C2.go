package main

import "fmt"

func simulate(N int, A []int) int {
	balls := make([]int, 0) // 空の列

	for i := 0; i < N; i++ {
		balls = append(balls, 2*A[i]) // ボールの大きさを列に追加

		for len(balls) > 1 {
			if balls[len(balls)-1] != balls[len(balls)-2] { // 右から1番目と2番目のボールの大きさが異なる場合
				break
			} else { // 右から1番目と2番目のボールの大きさが等しい場合
				removedSum := balls[len(balls)-1] + balls[len(balls)-2] // 取り除かれた2つのボールの大きさの和
				balls = balls[:len(balls)-2]                            // 2つのボールを取り除く
				balls = append(balls, removedSum)                       // 和の大きさのボールを列の一番右に追加
			}
		}
	}

	return len(balls)
}

func main() {
	// N :=
	// A := []int{1, 2, 3, 2, 1}
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	result := simulate(N, A)
	fmt.Println(result + 1) // 列に残るボールの数を表示
}

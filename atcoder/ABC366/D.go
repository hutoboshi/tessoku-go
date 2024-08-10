package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, Q int
	fmt.Scan(&N)

	// 3次元累積和を作成するためのスライスを宣言
	A := make([][][]int, N+1)
	for i := 0; i <= N; i++ {
		A[i] = make([][]int, N+1)
		for j := 0; j <= N; j++ {
			A[i][j] = make([]int, N+1)
		}
	}

	// 入力から3次元配列の値を読み込む
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			for k := 1; k <= N; k++ {
				fmt.Scan(&A[i][j][k])
			}
		}
	}

	// 3次元累積和を計算
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			for k := 1; k <= N; k++ {
				A[i][j][k] += A[i-1][j][k] + A[i][j-1][k] + A[i][j][k-1] - A[i-1][j-1][k] - A[i-1][j][k-1] - A[i][j-1][k-1] + A[i-1][j-1][k-1]
			}
		}
	}

	// クエリを処理する
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Scan(&Q)
	for i := 0; i < Q; i++ {
		var Lx, Rx, Ly, Ry, Lz, Rz int
		fmt.Fscanf(reader, "%d %d %d %d %d %d\n", &Lx, &Rx, &Ly, &Ry, &Lz, &Rz)

		result := A[Rx][Ry][Rz] -
			A[Lx-1][Ry][Rz] - A[Rx][Ly-1][Rz] - A[Rx][Ry][Lz-1] +
			A[Lx-1][Ly-1][Rz] + A[Lx-1][Ry][Lz-1] + A[Rx][Ly-1][Lz-1] -
			A[Lx-1][Ly-1][Lz-1]

		fmt.Fprintln(writer, result)
	}
}

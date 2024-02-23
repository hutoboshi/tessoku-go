package main

import "fmt"

func main() {
	var H, W, N int
	fmt.Scan(&H, &W, &N)

	A := make([]int, N)
	B := make([]int, N)
	C := make([]int, N)
	D := make([]int, N)
	X := make([][]int, H+2)
	for i := range X {
		X[i] = make([]int, W+2)
	}

	for t := 0; t < N; t++ {
		fmt.Scan(&A[t], &B[t], &C[t], &D[t])
	}

	for t := 0; t < N; t++ {
		X[A[t]][B[t]]++
		X[A[t]][D[t]+1]--
		X[C[t]+1][B[t]]--
		X[C[t]+1][D[t]+1]++
	}

	Z := make([][]int, H+2)
	for i := range Z {
		Z[i] = make([]int, W+2)
	}

	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			Z[i][j] = Z[i][j-1] + X[i][j]
		}
	}

	for j := 1; j <= W; j++ {
		for i := 1; i <= H; i++ {
			Z[i][j] += Z[i-1][j]
		}
	}

	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			if j >= 2 {
				fmt.Print(" ")
			}
			fmt.Print(Z[i][j])
		}
		fmt.Println()
	}
}

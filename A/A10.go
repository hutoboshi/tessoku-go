package main

import "fmt"

func main() {
	var N, D int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	fmt.Scan(&D)

	L := make([]int, D)
	R := make([]int, D)
	for i := 0; i < D; i++ {
		fmt.Scan(&L[i], &R[i])
	}

	P := make([]int, N)
	P[0] = A[0]
	for i := 1; i < N; i++ {
		if P[i-1] > A[i] {
			P[i] = P[i-1]
		} else {
			P[i] = A[i]
		}
	}

	Q := make([]int, N)
	Q[N-1] = A[N-1]
	for i := N - 2; i >= 0; i-- {
		if Q[i+1] > A[i] {
			Q[i] = Q[i+1]
		} else {
			Q[i] = A[i]
		}
	}

	for i := 0; i < D; i++ {
		fmt.Println(max(P[L[i]-2], Q[R[i]]))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

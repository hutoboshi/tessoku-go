package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	arrX := make([][]int, h+1)
	for i := 0; i < h+1; i++ {
		arrX[i] = make([]int, w+1)
	}

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			fmt.Scan(&arrX[i][j])
		}
	}

	var q int
	fmt.Scan(&q)
	arrA := make([]int, q)
	arrB := make([]int, q)
	arrC := make([]int, q)
	arrD := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&arrA[i], &arrB[i], &arrC[i], &arrD[i])
	}

}

package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	arrA := make([]int, n+1)
	arrA[0] = 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&arrA[i])
	}

	arrSum := make([]int, n+1)
	arrSum[0] = 0
	for i := 1; i <= n; i++ {
		arrSum[i] = arrSum[i-1] + arrA[i]
	}

	arrL := make([]int, q)
	arrR := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&arrL[i], &arrR[i])
	}

	for i := 0; i < q; i++ {
		fmt.Println(arrSum[arrR[i]] - arrSum[arrL[i]-1])
	}
}

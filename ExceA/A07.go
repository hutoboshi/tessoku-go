package main

import "fmt"

func main() {
	var d, n int
	fmt.Scan(&d, &n)

	arrL := make([]int, n)
	arrR := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arrL[i], &arrR[i])
	}

	arrAns := make([]int, d+2)
	for i := 0; i < d+2; i++ {
		arrAns[i] = 0
	}

	for i := 0; i < n; i++ {
		arrAns[arrL[i]]++
		arrAns[arrR[i]+1]--
	}

	ans := 0
	for i := 1; i <= d; i++ {
		ans += arrAns[i]
		fmt.Println(ans)
	}
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	arrA := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arrA[i][0], &arrA[i][1])
	}

	sort.Slice(arrA, func(i, j int) bool {
		return arrA[i][1] < arrA[j][1]
	})

	CurrentTime := 0
	ans := 0
	for i := 0; i < n; i++ {
		if CurrentTime <= arrA[i][0] {
			CurrentTime = arrA[i][1]
			ans++
		}
	}

	fmt.Println(ans)
}

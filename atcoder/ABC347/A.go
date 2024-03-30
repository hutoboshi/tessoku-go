package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	arrA := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arrA[i])
	}

	ans := []int{}
	for i := 0; i < n; i++ {
		if arrA[i]%k == 0 {
			ans = append(ans, arrA[i]/k)
		}
	}

	sort.Ints(ans)

	for i, v := range ans {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}

	fmt.Println()
}

package main

import (
	"fmt"
)

func solve(n, k int, r []int, seq []int, lv int) {
	if lv == n {
		sum := 0
		for _, v := range seq {
			sum += v
		}
		if sum%k == 0 {
			for _, v := range seq {
				fmt.Print(v, " ")
			}
			fmt.Println()
		}
		return
	}

	for i := 1; i <= r[lv]; i++ {
		seq[lv] = i
		solve(n, k, r, seq, lv+1)
	}
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	r := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&r[i])
	}

	seq := make([]int, n)
	solve(n, k, r, seq, 0)
}

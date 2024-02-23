package main

import "fmt"

func main() {
	var d, n int
	fmt.Scan(&d, &n)

	days := make([]int, d+1)

	var l, r int
	for i := 0; i < n; i++ {
		fmt.Scan(&l)
		fmt.Scan(&r)

		days[l-1]++
		days[r]--
	}

	var ans int
	for i := 0; i < d; i++ {
		ans += days[i]
		fmt.Println(ans)
	}
}

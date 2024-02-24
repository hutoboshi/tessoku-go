package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	r := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			r[i] = 0
		} else {
			r[i] = r[i-1]
		}

		for r[i] < n-1 && a[r[i]+1]-a[i] <= k {
			r[i]++
		}
	}

	var ans int
	for i := 0; i < n-1; i++ {
		ans += r[i] - i
	}
	fmt.Println(ans)

}

package main

import "fmt"

func search(x int, a []int) int {
	l := 0
	r := len(a) - 1
	for l <= r {
		m := (l + r) / 2
		if x < a[m] {
			r = m - 1
		} else if x == a[m] {
			return m
		} else {
			l = m + 1
		}
	}
	return -1
}

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	ans := search(x, a)
	fmt.Println(ans + 1)

}

package main

import "fmt"

func check(x, n, k int, a []int) bool {
	sum := 0
	for i := 0; i < n; i++ {
		sum += x / a[i]
	}
	if sum >= k {
		return true
	} else {
		return false
	}
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	left := 1
	right := int(1e9)
	for left < right {
		mid := (left + right) / 2
		ans := check(mid, n, k, a)

		if !ans {
			left = mid + 1
		} else {
			right = mid
		}
	}

	fmt.Println(left)
}

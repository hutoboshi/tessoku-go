package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	r := make([]int, q)
	l := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&l[i], &r[i])
	}

	s := make([]int, n+1)
	s[0] = 0
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + a[i]
	}

	for i := 0; i < q; i++ {
		fmt.Println(s[r[i]] - s[l[i]-1])
	}
}

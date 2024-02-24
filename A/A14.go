package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
	}

	var p []int
	for _, x := range a {
		for _, y := range b {
			p = append(p, x+y)
		}
	}
	var q []int
	for _, x := range c {
		for _, y := range d {
			q = append(q, x+y)
		}
	}

	sort.Ints(q)

	for _, p1 := range p {
		pos1 := sort.Search(len(q), func(i int) bool {
			return q[i] >= k-p1
		})
		if pos1 < len(q) && q[pos1] == k-p1 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

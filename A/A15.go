package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	t := make([]int, 0)
	seen := make(map[int]bool)
	for _, val := range a {
		if !seen[val] {
			t = append(t, val)
			seen[val] = true
		}
	}
	sort.Ints(t)

	b := make([]int, n)
	for i := 0; i < n; i++ {
		pos := sort.Search(len(t), func(j int) bool { return t[j] >= a[i] })
		b[i] = pos + 1
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Print(b[i])
		} else {
			fmt.Print(" ")
			fmt.Print(b[i])
		}
	}
	fmt.Println()
}

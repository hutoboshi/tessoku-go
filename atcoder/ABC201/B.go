package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	mountain := make(map[int]string)
	keys := make([]int, n)
	var tmp1 int
	var tmp2 string
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp1, &tmp2)
		mountain[tmp1] = tmp2
		keys[i] = tmp1
	}

	sort.Ints(keys)

	fmt.Println(mountain[keys[len(keys)-2]])
}

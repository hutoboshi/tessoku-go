package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	answer := false

	for i := 0; i < n; i++ {
		if a[i] == x {
			answer = true
			break
		}
	}

	if answer {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

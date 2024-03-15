package main

import "fmt"

func main() {
	var n, x, a int
	fmt.Scan(&n, &x)

	ans := false
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a == x {
			ans = true
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

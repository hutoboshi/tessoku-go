package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	if k >= 2*n-2 && k%2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arrA := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arrA[i])
	}

	for i := 0; i < n-1; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(arrA[i] * arrA[i+1])
	}

	fmt.Println()
}

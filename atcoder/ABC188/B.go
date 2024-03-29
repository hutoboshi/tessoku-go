package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arrA := make([]int, n)
	arrB := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arrA[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&arrB[i])
	}

	var ans int
	for i := 0; i < n; i++ {
		ans += arrA[i] * arrB[i]
	}

	if ans == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

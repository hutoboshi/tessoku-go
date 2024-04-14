package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arrA := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arrA[i])
	}

	var point int
	for i := 0; i < n; i++ {
		point += arrA[i]
	}

	if point == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(point * -1)
	}
}

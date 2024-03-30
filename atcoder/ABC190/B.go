package main

import "fmt"

func main() {
	var n, s, d int
	fmt.Scan(&n, &s, &d)

	arrX := make([]int, n)
	arrY := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arrX[i], &arrY[i])
	}

	for i := 0; i < n; i++ {
		if arrX[i] < s && arrY[i] > d {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

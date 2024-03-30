package main

import "fmt"

func main() {
	var n, x, tmp int
	fmt.Scan(&n, &x)
	arrA := []int{}

	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		if tmp != x {
			arrA = append(arrA, tmp)
		}
	}

	for i, v := range arrA {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()

}

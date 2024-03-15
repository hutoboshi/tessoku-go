package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 9; i >= 0; i-- {
		wari := 1 << i
		fmt.Print((n / wari) % 2)
	}

	fmt.Println()
}

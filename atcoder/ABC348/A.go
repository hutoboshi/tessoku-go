package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Print("x")
		} else {
			fmt.Print("o")
		}
	}
	fmt.Println()
}

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(GCD(a, b))
}

func GCD(a, b int) int {
	for a >= 1 && b >= 1 {
		if a >= b {
			a %= b
		} else {
			b %= a
		}
	}
	if a >= 1 {
		return a
	}
	return b
}

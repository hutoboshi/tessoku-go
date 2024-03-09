package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	A1 := n / 3
	A2 := n / 5
	A3 := n / 15

	fmt.Println(A1 + A2 - A3)
}

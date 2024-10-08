package main

import "fmt"

func main() {
	sum := 0
	for A := 1; A <= 512; A++ {
		for B := 1; B <= 512; B++ {
			sum += A % B
		}
	}

	fmt.Println(sum)
}

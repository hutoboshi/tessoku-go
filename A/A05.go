package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var count int
	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			z := k - x - y
			if 1 <= z && z <= n {
				count++
			}
		}
	}

	fmt.Println(count)
}

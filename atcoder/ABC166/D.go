package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	for a := -int(1e3); a <= int(1e3); a++ {
		for b := -int(1e3); b <= int(1e3); b++ {
			if a*a*a*a*a-b*b*b*b*b == x {
				fmt.Println(a, b)
				return
			}
		}
	}
}

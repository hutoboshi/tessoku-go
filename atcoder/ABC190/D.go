package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	i := 1

	for i*i <= 2*n {
		if 2*n%i == 0 {
			x := i
			y := 2 * n / i
			if x%2 != y%2 {
				ans += 2
			}
		}
		i++
	}
	fmt.Println(ans)
}

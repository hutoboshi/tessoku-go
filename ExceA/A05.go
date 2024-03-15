package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var ans int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if 1 <= k-i-j && k-i-j <= n {
				ans++
			}
		}
	}

	fmt.Println(ans)
}

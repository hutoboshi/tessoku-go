package main

import "fmt"

func main() {
	var a, b int64
	fmt.Scan(&a, &b)

	fmt.Println(power(a, b, 1000000007))
}

func power(a, b, m int64) int64 {
	p, ans := a, int64(1)
	for i := int64(0); i < 30; i++ {
		wari := int64(1 << i)
		if (b/wari)%2 == 1 {
			ans = (ans * p) % m
		}
		p = (p * p) % m
	}
	return ans
}

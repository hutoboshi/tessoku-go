package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	st := make(map[int]bool)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a) //数列を入力から受け取る
		if a <= k {
			st[a] = true
		}
	}

	ans := int64(k * (k + 1) / 2)
	for i := range st {
		ans -= int64(i)
	}

	fmt.Println(ans)
}

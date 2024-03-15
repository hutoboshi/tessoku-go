package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	arrP := make([]int, n)
	arrQ := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arrP[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&arrQ[i])
	}

	ans := "No"
	for _, v := range arrP {
		for _, v2 := range arrQ {
			if v+v2 == k {
				ans = "Yes"
			}
		}
	}

	fmt.Println(ans)
}
